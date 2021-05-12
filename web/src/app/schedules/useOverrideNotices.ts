import { Notice } from '../../schema'
import { parseInterval, SpanISO } from '../util/shifts'
import { useQuery, gql } from '@apollo/client'
import _ from 'lodash'
import { Interval } from 'luxon'

const scheduleQuery = gql`
  query ($id: ID!) {
    schedule(id: $id) {
      id
      name
      temporarySchedules {
        start
        end
      }
    }
  }
`
export default function useOverrideNotices(
  scheduleID: string,
  value: SpanISO,
): Notice[] {
  const { data, loading } = useQuery(scheduleQuery, {
    variables: {
      id: scheduleID,
    },
    pollInterval: 0,
  })
  if (loading) {
    return []
  }
  const tempSchedules = _.get(data, 'schedule.temporarySchedules')
  const valueInterval = parseInterval(value)
  const doesOverlap = tempSchedules
    .map(parseInterval)
    .some((invl: Interval) => invl.overlaps(valueInterval))

  if (!doesOverlap) {
    return []
  }
  return [
    {
      type: 'WARNING',
      message: 'This override overlaps with one or more temporary schedules',
      details: 'Overrides do not take affect during temporary schedules',
    },
  ]
}
