import React, { useState } from 'react'
import Grid from '@material-ui/core/Grid'
import ToggleIcon from '@material-ui/icons/CompareArrows'
import { DateTime } from 'luxon'
import { FormField } from '../../forms'
import { UserSelect } from '../../selection'
import ClickableText from '../../util/ClickableText'
import { ISODateTimePicker } from '../../util/ISOPickers'
import { Value } from './sharedUtils'
import NumberField from '../../util/NumberField'
import { useScheduleTZ } from './hooks'

export default function TempSchedAddShiftForm({
  min,
  scheduleID,
}: {
  min?: string
  scheduleID: string
}): JSX.Element {
  const [manualEntry, setManualEntry] = useState(false)
  const [now] = useState(DateTime.utc().startOf('minute').toISO())
  const { q, zone, isLocalZone, zoneAbbr } = useScheduleTZ(scheduleID)

  return (
    <React.Fragment>
      <Grid item>
        <FormField
          fullWidth
          component={UserSelect}
          label='Select a User'
          name='userID'
        />
      </Grid>
      <Grid item>
        <FormField
          fullWidth
          component={ISODateTimePicker}
          label={'Shift Start' + (isLocalZone ? '' : ` (${zoneAbbr})`)}
          name='start'
          min={min ?? now}
          mapOnChangeValue={(value: string, formValue: Value) => {
            if (!manualEntry) {
              const diff = DateTime.fromISO(value).diff(
                DateTime.fromISO(formValue.start),
              )
              formValue.end = DateTime.fromISO(formValue.end).plus(diff).toISO()
            }
            return value
          }}
          timeZone={zone}
          disabled={q.loading}
        />
      </Grid>
      <Grid item>
        {manualEntry ? (
          <FormField
            fullWidth
            component={ISODateTimePicker}
            label={'Shift End' + (isLocalZone ? '' : ` (${zoneAbbr})`)}
            name='end'
            min={min ?? now}
            hint={
              <ClickableText
                data-cy='toggle-duration-on'
                onClick={() => setManualEntry(false)}
                endIcon={<ToggleIcon />}
              >
                Configure as duration
              </ClickableText>
            }
            timeZone={zone}
            disabled={q.loading}
          />
        ) : (
          <FormField
            fullWidth
            component={NumberField}
            label='Shift Duration (hours)'
            name='end'
            float
            // value held in form input
            mapValue={(nextVal: string, formValue: Value) => {
              const nextValDT = DateTime.fromISO(nextVal)
              if (!formValue || !nextValDT.isValid) return ''
              return nextValDT
                .diff(DateTime.fromISO(formValue.start), 'hours')
                .hours.toString()
            }}
            // value held in state
            mapOnChangeValue={(nextVal: string, formValue: Value) => {
              if (!nextVal) return ''
              return DateTime.fromISO(formValue.start)
                .plus({ hours: parseFloat(nextVal) })
                .toISO()
            }}
            step='any'
            min={0}
            hint={
              <ClickableText
                data-cy='toggle-duration-off'
                onClick={() => setManualEntry(true)}
                endIcon={<ToggleIcon />}
              >
                Configure as date/time
              </ClickableText>
            }
          />
        )}
      </Grid>
    </React.Fragment>
  )
}
