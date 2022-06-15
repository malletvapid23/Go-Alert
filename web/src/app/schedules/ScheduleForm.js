import React from 'react'
import p from 'prop-types'
import { FormContainer, FormField } from '../forms'
import { TextField, Grid } from '@mui/material'
import { TimeZoneSelect } from '../selection'

export default function ScheduleForm(props) {
  return (
    <FormContainer optionalLabels {...props}>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <FormField
            fullWidth
            component={TextField}
            name='name'
            label='Name'
            required
          />
        </Grid>
        <Grid item xs={12}>
          <FormField
            fullWidth
            component={TextField}
            multiline
            name='description'
            label='Description'
          />
        </Grid>
        <Grid item xs={12}>
          <FormField
            fullWidth
            component={TimeZoneSelect}
            name='time-zone'
            fieldName='timeZone'
            label='Time Zone'
            required
          />
        </Grid>
      </Grid>
    </FormContainer>
  )
}

ScheduleForm.propTypes = {
  value: p.shape({
    name: p.string.isRequired,
    description: p.string.isRequired,
    timeZone: p.string.isRequired,
  }).isRequired,

  errors: p.arrayOf(
    p.shape({
      field: p.oneOf(['name', 'description', 'timeZone']).isRequired,
      message: p.string.isRequired,
    }),
  ),

  onChange: p.func.isRequired,
  disabled: p.bool,
}
