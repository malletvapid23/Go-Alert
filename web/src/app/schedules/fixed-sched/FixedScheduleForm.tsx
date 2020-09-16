import React, { useState, ReactNode } from 'react'
import SwipeableViews from 'react-swipeable-views'
import { virtualize, bindKeyboard } from 'react-swipeable-views-utils'
import { FormContainer } from '../../forms'
import { FieldError } from '../../util/errutil'
import SchedulesTimesStep from './ScheduleTimesStep'
import AddShiftsStep from './AddShiftsStep'
import ReviewStep from './ReviewStep'
import SuccessStep from './SuccessStep'

// allows changing the index programatically
const VirtualizeAnimatedViews = bindKeyboard(virtualize(SwipeableViews))

interface FixedScheduleFormProps {
  activeStep: number
  setStep: (step: number) => void
  value: any
  onChange: (val: any) => any
  disabled: boolean
  errors: FieldError[]
}

export default function FixedScheduleForm({
  activeStep,
  setStep,
  ...otherProps
}: FixedScheduleFormProps) {
  const bodyStyle = {
    display: 'flex',
    alignItems: 'center', // vertical align
    justifyContent: 'center', // horizontal align
    height: '100%',
    width: '100%',
  }

  const containerStyle = {
    width: '35%', // ensures form fields don't shrink down too small
    marginBottom: '10%', // slightly raise higher than center of screen
  }

  interface SlideRenderer {
    index: number
    key: number
  }
  function renderSlide({ index, key }: SlideRenderer): ReactNode {
    switch (index) {
      case 0:
        return <SchedulesTimesStep key={key} />
      case 1:
        return <AddShiftsStep key={key} />
      case 2:
        return <ReviewStep key={key} />
      case 3:
        return <SuccessStep key={key} />
      default:
        return null
    }
  }

  return (
    <div style={bodyStyle}>
      <div style={containerStyle}>
        <FormContainer optionalLabels {...otherProps}>
          <VirtualizeAnimatedViews
            index={activeStep}
            onChangeIndex={(i: number) => setStep(i)}
            slideRenderer={renderSlide}
          />
        </FormContainer>
      </div>
    </div>
  )
}
