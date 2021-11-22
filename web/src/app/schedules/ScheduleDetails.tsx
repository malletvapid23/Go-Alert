import React, { useState, useCallback } from 'react'
import { gql, useQuery } from '@apollo/client'
import { Redirect } from 'react-router-dom'
import _ from 'lodash'
import { Edit, Delete } from '@material-ui/icons'

import DetailsPage from '../details/DetailsPage'
import ScheduleEditDialog from './ScheduleEditDialog'
import ScheduleDeleteDialog from './ScheduleDeleteDialog'
import ScheduleCalendarQuery from './calendar/ScheduleCalendarQuery'
import { QuerySetFavoriteButton } from '../util/QuerySetFavoriteButton'
import CalendarSubscribeButton from './calendar-subscribe/CalendarSubscribeButton'
import Spinner from '../loading/components/Spinner'
import { ObjectNotFound, GenericError } from '../error-pages'
import TempSchedDialog from './temp-sched/TempSchedDialog'
import TempSchedDeleteConfirmation from './temp-sched/TempSchedDeleteConfirmation'
import { ScheduleAvatar } from '../util/avatars'
import { useConfigValue } from '../util/RequireConfig'
import ScheduleCalendarOverrideDialog from './calendar/ScheduleCalendarOverrideDialog'
import { useIsWidthDown } from '../util/useWidth'
import { TempSchedValue } from './temp-sched/sharedUtils'

const query = gql`
  fragment ScheduleTitleQuery on Schedule {
    id
    name
    description
  }
  query scheduleDetailsQuery($id: ID!) {
    schedule(id: $id) {
      ...ScheduleTitleQuery
      timeZone
    }
  }
`

interface ScheduleCalendarContextProps {
  onNewTempSched: () => void
  onEditTempSched: (v: TempSchedValue) => void
  onDeleteTempSched: React.Dispatch<React.SetStateAction<null>>
  setOverrideDialog: React.Dispatch<
    React.SetStateAction<ScheduleCalendarOverrideDialogProps | null>
  >
}

// todo - move this interface to more generic file when those are converted to typescript
interface ScheduleCalendarOverrideDialogProps {
  variantOptions: string[]
  removeUserReadOnly: boolean
  defaultValue: {
    addUserID?: string
    removeUserID?: string
    start: string
    end: string
  }
}

export const ScheduleCalendarContext =
  React.createContext<ScheduleCalendarContextProps>({
    onNewTempSched: () => {},
    onEditTempSched: () => {},
    onDeleteTempSched: () => {},
    setOverrideDialog: () => {},
  })

interface ScheduleDetailsProps {
  scheduleID: string
}

export default function ScheduleDetails({
  scheduleID,
}: ScheduleDetailsProps): JSX.Element {
  const [showEdit, setShowEdit] = useState(false)
  const [showDelete, setShowDelete] = useState(false)
  const [configTempSchedule, setConfigTempSchedule] = useState<
    TempSchedValue | null | undefined
  >()
  const [deleteTempSchedule, setDeleteTempSchedule] = useState(null)
  const isMobile = useIsWidthDown('sm')

  const [slackEnabled] = useConfigValue('Slack.Enable')

  const onNewTempSched = useCallback(() => setConfigTempSchedule(null), [])
  const onEditTempSched = useCallback(
    (v: TempSchedValue) => setConfigTempSchedule(v),
    [],
  )
  const onDeleteTempSched = useCallback(setDeleteTempSchedule, [])
  const [overrideDialog, setOverrideDialog] =
    useState<ScheduleCalendarOverrideDialogProps | null>(null)

  const {
    data: _data,
    loading,
    error,
  } = useQuery(query, {
    variables: { id: scheduleID },
    returnPartialData: true,
  })

  const data = _.get(_data, 'schedule', null)

  if (loading && !data?.name) return <Spinner />
  if (error) return <GenericError error={error.message} />

  if (!data) {
    return showDelete ? <Redirect to='/schedules' push /> : <ObjectNotFound />
  }

  return (
    <React.Fragment>
      {showEdit && (
        <ScheduleEditDialog
          scheduleID={scheduleID}
          onClose={() => setShowEdit(false)}
        />
      )}
      {showDelete && (
        <ScheduleDeleteDialog
          scheduleID={scheduleID}
          onClose={() => setShowDelete(false)}
        />
      )}
      {configTempSchedule !== undefined && (
        <TempSchedDialog
          value={configTempSchedule}
          onClose={() => setConfigTempSchedule(null)}
          scheduleID={scheduleID}
        />
      )}
      {deleteTempSchedule && (
        <TempSchedDeleteConfirmation
          value={deleteTempSchedule}
          onClose={() => setDeleteTempSchedule(null)}
          scheduleID={scheduleID}
        />
      )}
      <DetailsPage
        avatar={<ScheduleAvatar />}
        title={data.name}
        subheader={`Time Zone: ${data.timeZone || 'Loading...'}`}
        details={data.description}
        pageContent={
          <ScheduleCalendarContext.Provider
            value={{
              onNewTempSched,
              onEditTempSched,
              onDeleteTempSched,
              setOverrideDialog,
            }}
          >
            {!isMobile && <ScheduleCalendarQuery scheduleID={scheduleID} />}
            {overrideDialog && (
              <ScheduleCalendarOverrideDialog
                defaultValue={overrideDialog.defaultValue}
                variantOptions={overrideDialog.variantOptions}
                scheduleID={scheduleID}
                onClose={() => setOverrideDialog(null)}
                removeUserReadOnly={overrideDialog.removeUserReadOnly}
              />
            )}
          </ScheduleCalendarContext.Provider>
        }
        primaryActions={[
          <CalendarSubscribeButton
            key='primary-action-subscribe'
            scheduleID={scheduleID}
          />,
        ]}
        secondaryActions={[
          {
            label: 'Edit',
            icon: <Edit />,
            handleOnClick: () => setShowEdit(true),
          },
          {
            label: 'Delete',
            icon: <Delete />,
            handleOnClick: () => setShowDelete(true),
          },
          <QuerySetFavoriteButton
            key='secondary-action-favorite'
            // todo - redo this prop when this component converted to typescript
            {...{ scheduleId: scheduleID }}
            // scheduleID={scheduleID}
          />,
        ]}
        links={[
          {
            label: 'Assignments',
            url: 'assignments',
            subText: 'Manage rules for rotations and users',
          },
          {
            label: 'Escalation Policies',
            url: 'escalation-policies',
            subText: 'Find escalation policies that link to this schedule',
          },
          {
            label: 'Overrides',
            url: 'overrides',
            subText: 'Add, remove, or replace a user temporarily',
          },
          {
            label: 'Shifts',
            url: 'shifts',
            subText: 'Review a list of past and future on-call shifts',
          },
        ].concat(
          slackEnabled
            ? [
                {
                  // only slack is supported ATM, so hide the link if disabled

                  label: 'On-Call Notifications',
                  url: 'on-call-notifications',
                  subText: 'Set up notifications to know who is on-call',
                },
              ]
            : [],
        )}
      />
    </React.Fragment>
  )
}
