import React from 'react'
import { PropTypes as p } from 'prop-types'
import { FormContainer, FormField } from '../forms'
import { UserSelect } from '../selection'
import { Grid } from '@mui/material'

export default function UserForm(props) {
  return (
    <FormContainer {...props}>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <FormField
            component={UserSelect}
            disabled={false}
            fieldName='users'
            fullWidth
            label='Select User(s)'
            multiple
            name='users'
            required
            value={props.value.users}
          />
        </Grid>
      </Grid>
    </FormContainer>
  )
}

UserForm.propTypes = {
  errors: p.array,
  onChange: p.func,
  disabled: p.bool,
  value: p.shape({
    users: p.arrayOf(p.string),
  }).isRequired,
}
