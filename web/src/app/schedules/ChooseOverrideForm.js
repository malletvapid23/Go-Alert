import React from 'react'
import p from 'prop-types'
import { FormContainer } from '../forms'
import {
  Grid,
  RadioGroup,
  FormControlLabel,
  Radio,
  FormHelperText,
} from '@material-ui/core'
import { variantDetails } from './ScheduleOverrideCreateDialog'

export default function ChooseOverrideForm(props) {
  const { onVariantChange, activeVariant, variantOptions, ...formProps } = props

  return (
    <FormContainer {...formProps} optionalLabels>
      <Grid item xs={12}>
        <RadioGroup
          required
          aria-label='Choose an override action'
          name='variant'
          onChange={(e) => onVariantChange(e.target.value)}
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
    </FormContainer>
  )
}

ChooseOverrideForm.propTypes = {
  value: p.shape({
    addUserID: p.string.isRequired,
    removeUserID: p.string.isRequired,
    start: p.string.isRequired,
    end: p.string.isRequired,
  }).isRequired,

  disabled: p.bool.isRequired,
  activeVariant: p.string.isRequired,
  onVariantChange: p.func.isRequired,
  variantOptions: p.arrayOf(p.string),
}
