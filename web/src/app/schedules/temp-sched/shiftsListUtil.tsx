import _ from 'lodash'
import { DateTime, Interval } from 'luxon'
import { fmtLocal, Shift } from './sharedUtils'
import Tooltip from '@material-ui/core/Tooltip/Tooltip'

import {
  FlatListListItem,
  FlatListNotice,
  FlatListSub,
} from '../../lists/FlatList'
import * as luxonHelpers from '../../util/luxon-helpers'
import { parseInterval } from '../../util/shifts'
import React from 'react'

export const fmtTime = (dt: DateTime): string =>
  dt.toLocaleString(DateTime.TIME_SIMPLE)

export type Sortable<T> = T & {
  // at is the earliest point in time for a list item
  at: DateTime
  // itemType categorizes a list item
  itemType: 'subheader' | 'gap' | 'shift' | 'start' | 'end' | 'outOfBounds'
  // ends is the latest point in time for a list item
  ends?: DateTime
}

export function getSubheaderItems(
  schedInterval: Interval,
  shifts: Shift[],
  zone: luxonHelpers.ExplicitZone,
): Sortable<FlatListSub>[] {
  if (!schedInterval.isValid) {
    return []
  }

  // get earliest and farthest out start/end times
  const lowerBound = DateTime.min(
    schedInterval.start,
    ...shifts.map((s) => DateTime.fromISO(s.start, { zone })),
  )

  const upperBound = DateTime.max(
    schedInterval.end,
    ...shifts.map((s) => DateTime.fromISO(s.end, { zone })),
  )

  const dayInvs = luxonHelpers.splitAtMidnight(
    Interval.fromDateTimes(lowerBound, upperBound),
  )

  return dayInvs.map((day) => {
    const at = day.start.startOf('day')
    return {
      id: 'header_' + at.toISO(),
      subHeader: day.start.toFormat('cccc, LLLL d'),
      at,
      itemType: 'subheader',
    }
  })
}

export function getOutOfBoundsItems(
  schedInterval: Interval,
  shifts: Shift[],
  zone: luxonHelpers.ExplicitZone,
): Sortable<FlatListNotice>[] {
  if (!schedInterval.isValid) {
    return []
  }

  // get earliest and farthest out start/end times
  const lowerBound = DateTime.min(
    schedInterval.start,
    ...shifts.map((s) => DateTime.fromISO(s.start, { zone })),
  )

  const upperBound = DateTime.max(
    schedInterval.end,
    ...shifts.map((s) => DateTime.fromISO(s.end, { zone })),
  )

  const beforeStart = Interval.fromDateTimes(
    lowerBound,
    schedInterval.start,
  ).mapEndpoints((e) => e.startOf('day')) // ensure sched start date is not included

  const afterEnd = Interval.fromDateTimes(
    schedInterval.end,
    upperBound,
  ).mapEndpoints((e) => e.plus({ day: 1 }).startOf('day')) // ensure sched end date is not included

  const daysBeforeStart = luxonHelpers.splitAtMidnight(beforeStart)
  const daysAfterEnd = luxonHelpers.splitAtMidnight(afterEnd)
  const intervals = daysBeforeStart.concat(daysAfterEnd)

  let details = ''
  return intervals.map((interval) => {
    if (interval.end <= schedInterval.start) {
      details = 'This day is before the set start date.'
    } else if (interval.start >= schedInterval.end) {
      details = 'This day is after the set end date.'
    }

    return {
      id: 'day-out-of-bounds_' + interval.start.toISO(),
      type: 'INFO',
      message: '',
      details,
      at: interval.start.startOf('day'),
      itemType: 'outOfBounds',
    }
  })
}

export function getCoverageGapItems(
  schedInterval: Interval,
  shifts: Shift[],
  zone: luxonHelpers.ExplicitZone,
  handleCoverageClick: (coverageGap: Interval) => void,
): Sortable<FlatListNotice>[] {
  if (!schedInterval.isValid) {
    return []
  }
  const shiftIntervals = shifts.map((s) => parseInterval(s, zone))
  const gapIntervals = _.flatMap(
    schedInterval.difference(...shiftIntervals),
    (inv) => luxonHelpers.splitAtMidnight(inv),
  )
  const isLocalZone = zone === DateTime.local().zoneName
  return gapIntervals.map((gap) => {
    let details = 'No coverage'
    let title = 'No coverage'
    if (gap.length('hours') === 24) {
      // nothing to do
      title = ''
    } else if (gap.start.equals(gap.start.startOf('day'))) {
      details += ` until ${fmtTime(gap.end)}`
      if (!isLocalZone) {
        title += ` until ${fmtLocal(gap.end.toISO())}`
      }
    } else if (gap.end.equals(gap.start.plus({ day: 1 }).startOf('day'))) {
      details += ` after ${fmtTime(gap.start)}`
      if (!isLocalZone) {
        title += ` after ${fmtLocal(gap.start.toISO())}`
      }
    } else {
      details += ` from ${fmtTime(gap.start)} to ${fmtTime(gap.end)}`
      if (!isLocalZone) {
        title += ` from ${fmtLocal(gap.start.toISO())} to ${fmtLocal(
          gap.end.toISO(),
        )}`
      }
    }

    return {
      'data-cy': 'day-no-coverage',
      id: 'day-no-coverage_' + gap.start.toISO(),
      type: 'WARNING',
      message: '',
      details: (
        <Tooltip title={title} placement='right'>
          <span>{details}</span>
        </Tooltip>
      ),
      at: gap.start,
      ends: gap.end,
      itemType: 'gap',
      handleOnClick: () => {
        handleCoverageClick(gap)
      },
    }
  })
}

export function sortItems(
  items: Sortable<FlatListListItem>[],
): Sortable<FlatListListItem>[] {
  return items.sort((a, b) => {
    if (a.at < b.at) return -1
    if (a.at > b.at) return 1

    // a and b are at same time; use item type priority instead
    // subheaders first
    if (a.itemType === 'subheader') return -1
    if (b.itemType === 'subheader') return 1
    // out of bounds info next
    if (a.itemType === 'outOfBounds') return -1
    if (b.itemType === 'outOfBounds') return 1
    // then start notice
    if (a.itemType === 'start') return -1
    if (b.itemType === 'start') return 1
    // then gaps
    if (a.itemType === 'gap') return -1
    if (b.itemType === 'gap') return 1
    // then shifts
    if (
      // both shifts
      a.itemType === 'shift' &&
      b.itemType === 'shift' &&
      // typescript hints
      'title' in a &&
      'title' in b &&
      a.title &&
      b.title
    ) {
      return a.title < b.title ? -1 : 1
    }
    if (a.itemType === 'shift') return -1
    if (b.itemType === 'shift') return 1
    // then end notice
    if (a.itemType === 'end') return -1
    if (b.itemType === 'end') return 1

    // identical items; should never get to this point
    return 0
  })
}
