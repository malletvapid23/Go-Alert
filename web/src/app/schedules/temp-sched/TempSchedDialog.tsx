import React, { useState, useEffect } from 'react'
import { useMutation, gql } from '@apollo/client'
import Checkbox from '@material-ui/core/Checkbox'
import DialogContentText from '@material-ui/core/DialogContentText'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import FormHelperText from '@material-ui/core/FormHelperText'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core'
import Alert from '@material-ui/lab/Alert'
import AlertTitle from '@material-ui/lab/AlertTitle'
import _ from 'lodash'
import { DateTime, Interval } from 'luxon'

import { fieldErrors, nonFieldErrors } from '../../util/errutil'
import FormDialog from '../../dialogs/FormDialog'
import { contentText, fmtLocal, Shift, Value } from './sharedUtils'
import { FormContainer, FormField } from '../../forms'
import TempSchedAddNewShift from './TempSchedAddNewShift'
import { isISOAfter, parseInterval } from '../../util/shifts'
import { getNextWeekday } from '../../util/luxon-helpers'
import { useScheduleTZ } from './hooks'
import TempSchedShiftsList from './TempSchedShiftsList'
import { ISODateTimePicker } from '../../util/ISOPickers'
import { getCoverageGapItems } from './shiftsListUtil'

const mutation = gql`
  mutation ($input: SetTemporaryScheduleInput!) {
    setTemporarySchedule(input: $input)
  }
`

function shiftEquals(a: Shift, b: Shift): boolean {
  return a.start === b.start && a.end === b.end && a.userID === b.userID
}

const useStyles = makeStyles((theme) => ({
  contentText,
  avatar: {
    backgroundColor: theme.palette.primary.main,
  },
  formContainer: {
    height: '100%',
  },
  noCoverageError: {
    marginTop: '.5rem',
    marginBottom: '.5rem',
  },
  rightPane: {
    [theme.breakpoints.down('md')]: {
      marginTop: '1rem',
    },
    overflow: 'hidden',
  },
  sticky: {
    position: 'sticky',
    top: 0,
  },
  tzNote: {
    fontStyle: 'italic',
  },
}))

type TempScheduleDialogProps = {
  onClose: () => void
  scheduleID: string
  value?: Value
}

export default function TempSchedDialog({
  onClose,
  scheduleID,
  value: _value,
}: TempScheduleDialogProps): JSX.Element {
  const classes = useStyles()
  const edit = Boolean(_value)
  const { q, zone, isLocalZone } = useScheduleTZ(scheduleID)
  const [now] = useState(DateTime.utc().startOf('minute').toISO())
  const [showForm, setShowForm] = useState(false)
  const [value, setValue] = useState({
    start: _value?.start ?? '',
    end: _value?.end ?? '',
    shifts: (_value?.shifts ?? []).map((s) =>
      _.pick(s, 'start', 'end', 'userID'),
    ),
  })
  const [shift, setShift] = useState(null as Shift | null)
  const [allowNoCoverage, setAllowNoCoverage] = useState(false)
  const [hasSubmitted, setHasSubmitted] = useState(false)

  useEffect(() => {
    // set default start, end times when zone is ready
    if (!value.start && !value.end && !q.loading && zone) {
      const nextMonday = getNextWeekday(1, DateTime.now(), zone)
      const nextFriday = nextMonday.plus({ days: 5 }) // thru to the end of Friday
      setValue({
        ...value,
        start: nextMonday.toISO(),
        end: nextFriday.toISO(),
      })
    }
  }, [q.loading, zone])

  function validate(): Error | null {
    if (isISOAfter(value.start, value.end)) {
      return new Error('Start date/time cannot be after end date/time.')
    }
    return null
  }

  const hasInvalidShift = (() => {
    if (q.loading) return false
    const schedInterval = parseInterval(value, zone)
    return value.shifts.some(
      (s) => !schedInterval.engulfs(parseInterval(s, zone)),
    )
  })()

  const shiftErrors = hasInvalidShift
    ? [
        {
          message:
            'One or more shifts extend beyond the start and/or end of this temporary schedule',
        },
      ]
    : []

  function handleCoverageGapClick(coverageGap: Interval): void {
    if (!showForm) setShowForm(true)
    setShift({
      userID: shift?.userID ?? '',
      start: coverageGap?.start.toISO(),
      end: coverageGap?.end.toISO(),
    })
  }

  const hasCoverageGaps = (() => {
    if (q.loading) return false
    const schedInterval = parseInterval(value, zone)
    return (
      getCoverageGapItems(
        schedInterval,
        value.shifts,
        zone,
        handleCoverageGapClick,
      ).length > 0
    )
  })()

  const [submit, { loading, error }] = useMutation(mutation, {
    onCompleted: () => onClose(),
    variables: {
      input: {
        ...value,
        scheduleID,
      },
    },
  })

  const handleSubmit = (): void => {
    setHasSubmitted(true)

    if (hasCoverageGaps && !allowNoCoverage) {
      return
    }

    submit()
  }

  const nonFieldErrs = nonFieldErrors(error).map((e) => ({
    message: e.message,
  }))
  const fieldErrs = fieldErrors(error).map((e) => ({
    message: `${e.field}: ${e.message}`,
  }))
  const noCoverageErrs =
    hasSubmitted && hasCoverageGaps && !allowNoCoverage
      ? [new Error('This temporary schedule has gaps in coverage.')]
      : []
  const errs = nonFieldErrs
    .concat(fieldErrs)
    .concat(shiftErrors)
    .concat(noCoverageErrs)

  return (
    <FormDialog
      fullHeight
      maxWidth='lg'
      title='Define a Temporary Schedule'
      onClose={onClose}
      loading={loading}
      errors={errs}
      notices={
        !value.start ||
        DateTime.fromISO(value.start, { zone }) >
          DateTime.utc().minus({ hour: 1 }) ||
        edit
          ? []
          : [
              {
                type: 'WARNING',
                message: 'Start time occurs in the past',
                details:
                  'Any shifts or changes made to shifts in the past will be ignored when submitting.',
              },
            ]
      }
      form={
        <FormContainer
          optionalLabels
          disabled={loading}
          value={value}
          onChange={(newValue: Value) => setValue(newValue)}
        >
          <Grid
            container
            className={classes.formContainer}
            justifyContent='space-between'
          >
            {/* left pane */}
            <Grid
              item
              xs={12}
              md={6}
              container
              alignContent='flex-start'
              spacing={2}
            >
              <Grid item xs={12}>
                <DialogContentText className={classes.contentText}>
                  The schedule will be exactly as configured here for the entire
                  duration (ignoring all assignments and overrides).
                </DialogContentText>
              </Grid>

              <Grid item xs={12}>
                <Typography color='textSecondary' className={classes.tzNote}>
                  Configuring in {zone}
                </Typography>
              </Grid>

              <Grid item xs={12} md={6}>
                <FormField
                  fullWidth
                  component={ISODateTimePicker}
                  required
                  name='start'
                  label='Schedule Start'
                  min={now}
                  max={DateTime.fromISO(now, { zone })
                    .plus({ year: 1 })
                    .toISO()}
                  validate={() => validate()}
                  timeZone={zone}
                  disabled={q.loading || edit}
                  hint={isLocalZone ? '' : fmtLocal(value.start)}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <FormField
                  fullWidth
                  component={ISODateTimePicker}
                  required
                  name='end'
                  label='Schedule End'
                  min={value.start}
                  max={DateTime.fromISO(value.start, { zone })
                    .plus({ month: 3 })
                    .toISO()}
                  validate={() => validate()}
                  timeZone={zone}
                  disabled={q.loading || edit}
                  hint={isLocalZone ? '' : fmtLocal(value.end)}
                />
              </Grid>

              <Grid item xs={12} className={classes.sticky}>
                <TempSchedAddNewShift
                  value={value}
                  onChange={(shifts: Shift[]) => setValue({ ...value, shifts })}
                  scheduleID={scheduleID}
                  edit={edit}
                  showForm={showForm}
                  setShowForm={setShowForm}
                  shift={shift}
                  setShift={setShift}
                />
              </Grid>
            </Grid>

            {/* right pane */}
            <Grid
              item
              xs={12}
              md={6}
              container
              spacing={2}
              className={classes.rightPane}
            >
              <Grid item xs={12}>
                <Typography variant='subtitle1' component='h3'>
                  Shifts
                </Typography>

                {hasSubmitted && hasCoverageGaps && (
                  <Alert severity='error' className={classes.noCoverageError}>
                    <AlertTitle>Gaps in coverage</AlertTitle>
                    <FormHelperText>
                      There are gaps in coverage. During these gaps, nobody on
                      the schedule will receive alerts. If you still want to
                      proceed, check the box below and retry.
                    </FormHelperText>
                    <FormControlLabel
                      label='Allow gaps in coverage'
                      labelPlacement='end'
                      control={
                        <Checkbox
                          data-cy='no-coverage-checkbox'
                          checked={allowNoCoverage}
                          onChange={(e) => setAllowNoCoverage(e.target.checked)}
                          name='allowCoverageGaps'
                        />
                      }
                    />
                  </Alert>
                )}

                <TempSchedShiftsList
                  scheduleID={scheduleID}
                  value={value.shifts}
                  start={value.start}
                  end={value.end}
                  onRemove={(shift: Shift) => {
                    setValue({
                      ...value,
                      shifts: value.shifts.filter(
                        (s) => !shiftEquals(shift, s),
                      ),
                    })
                  }}
                  edit={edit}
                  handleCoverageGapClick={handleCoverageGapClick}
                />
              </Grid>
            </Grid>
          </Grid>
        </FormContainer>
      }
      onSubmit={handleSubmit}
    />
  )
}
