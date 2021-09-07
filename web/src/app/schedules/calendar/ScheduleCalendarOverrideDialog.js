import React, { useContext, useEffect, useState } from 'react'
import { gql, useMutation } from '@apollo/client'
import p from 'prop-types'
import FormDialog from '../../dialogs/FormDialog'
import { DateTime } from 'luxon'
import {
  fieldErrors as getFieldErrors,
  nonFieldErrors,
} from '../../util/errutil'
import useOverrideNotices from '../useOverrideNotices'
import { ScheduleCalendarContext } from '../ScheduleDetails'
import { variantDetails } from '../ScheduleOverrideCreateDialog'
import ScheduleOverrideForm from '../ScheduleOverrideForm'
import {
  FormControlLabel,
  FormHelperText,
  Grid,
  Radio,
  RadioGroup,
} from '@material-ui/core'

const mutation = gql`
  mutation ($input: CreateUserOverrideInput!) {
    createUserOverride(input: $input) {
      id
    }
  }
`
export default function ScheduleCalendarOverrideDialog(props) {
  const { variantOptions = ['replace', 'remove', 'add', 'temp'] } = props

  const initialValue = {
    addUserID: '',
    removeUserID: '',
    start: DateTime.local().startOf('hour').toISO(),
    end: DateTime.local().startOf('hour').plus({ hours: 8 }).toISO(),
    ...props.defaultValue,
  }

  const [step, setStep] = useState(0)
  const [value, setValue] = useState(initialValue)
  const [fieldErrors, setFieldErrors] = useState([])
  const [activeVariant, setActiveVariant] = useState(variantOptions[0])

  const { onNewTempSched } = useContext(ScheduleCalendarContext)

  const notices = useOverrideNotices(props.scheduleID, value)

  const [mutate, { loading, error }] = useMutation(mutation, {
    variables: {
      input: {
        ...value,
        scheduleID: props.scheduleID,
      },
    },
    onCompleted: props.onClose,
  })

  useEffect(() => {
    setFieldErrors(getFieldErrors(error))
  }, [error])

  const onNext = () => {
    if (activeVariant === 'temp') {
      onNewTempSched()
      props.onClose()
    } else {
      setValue(initialValue)
      setFieldErrors([])
      setStep(step + 1)
    }
  }

  return (
    <FormDialog
      onClose={props.onClose}
      title={
        step === 0
          ? variantDetails.choose.title
          : variantDetails[activeVariant].title
      }
      subTitle={
        step === 0
          ? variantDetails.choose.desc
          : variantDetails[activeVariant].desc
      }
      errors={nonFieldErrors(error)}
      notices={step === 0 ? [] : notices} // create and edit dialogue
      loading={loading}
      form={
        <React.Fragment>
          {/* Step 0: Choose override variant page */}
          {step === 0 && (
            <Grid item xs={12}>
              <RadioGroup
                required
                aria-label='Choose an override action'
                name='variant'
                onChange={(e) => setActiveVariant(e.target.value)}
                value={activeVariant}
              >
                {variantOptions.map((variant) => (
                  <FormControlLabel
                    key={variant}
                    data-cy={`variant.${variant}`}
                    value={variant}
                    control={<Radio />}
                    label={
                      <React.Fragment>
                        {variantDetails[variant].name}
                        <FormHelperText>
                          {variantDetails[variant].helperText}
                        </FormHelperText>
                      </React.Fragment>
                    }
                  />
                ))}
              </RadioGroup>
            </Grid>
          )}
          {/* Step 1: Schedule override based on variant selected */}
          {step === 1 && (
            <ScheduleOverrideForm
              add={activeVariant !== 'remove'}
              remove={activeVariant !== 'add'}
              scheduleID={props.scheduleID}
              disabled={loading}
              errors={fieldErrors}
              value={value}
              onChange={(newValue) => setValue(newValue)}
              removeUserReadOnly={props.removeUserReadOnly}
            />
          )}
        </React.Fragment>
      }
      onSubmit={() => mutate()}
      onNext={step < 1 ? onNext : null}
      onBack={step > 0 ? () => setStep(step - 1) : null}
    />
  )
}

ScheduleCalendarOverrideDialog.defaultProps = {
  defaultValue: {},
}

ScheduleCalendarOverrideDialog.propTypes = {
  scheduleID: p.string.isRequired,
  onClose: p.func,
  removeUserReadOnly: p.bool,
  defaultValue: p.shape({
    addUserID: p.string,
    removeUserID: p.string,
    start: p.string,
    end: p.string,
  }),
  variantOptions: p.arrayOf(p.string),
}
