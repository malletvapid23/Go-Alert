import React from 'react'
import { gql } from '@apollo/client'
import QueryList from '../../lists/QueryList'
import ScheduleOnCallNotificationAction, {
  Rule,
} from './ScheduleOnCallNotificationAction'
import ScheduleOnCallNotificationCreateFab from './ScheduleOnCallNotificationCreateFab'

interface ScheduleOnCallNotificationsProps {
  scheduleID: string
}

export const query = gql`
  query scheduleCalendarShifts($id: ID!) {
    schedule(id: $id) {
      id
      onCallNotificationRules {
        id
        target {
          id
          type
          name
        }
        time
        weekdayFilter
      }
    }
  }
`

export const setMutation = gql`
  mutation ($input: SetScheduleOnCallNotificationRulesInput!) {
    setScheduleOnCallNotificationRules(input: $input)
  }
`

export default function ScheduleOnCallNotificationsList(
  p: ScheduleOnCallNotificationsProps,
): JSX.Element {
  return (
    <React.Fragment>
      <QueryList
        query={query}
        variables={{ id: p.scheduleID }}
        headerNote='Configure notifications for on-call updates'
        noSearch
        mapDataNode={(nr) => ({
          id: nr.id,
          title: nr.target.name,
          action: (
            <ScheduleOnCallNotificationAction
              rule={nr as Rule}
              scheduleID={p.scheduleID}
            />
          ),
        })}
      />
      <ScheduleOnCallNotificationCreateFab scheduleID={p.scheduleID} />
    </React.Fragment>
  )
}
