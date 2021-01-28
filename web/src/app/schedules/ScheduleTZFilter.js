import React from 'react'
import { gql, useQuery } from '@apollo/client'
import p from 'prop-types'
import { FormControlLabel, Switch } from '@material-ui/core'

import { useURLParam } from '../actions/hooks'

const tzQuery = gql`
  query($id: ID!) {
    schedule(id: $id) {
      id
      timeZone
    }
  }
`

export function ScheduleTZFilter(props) {
  const [zone, setZone] = useURLParam('tz', 'local')
  const { data, loading, error } = useQuery(tzQuery, {
    pollInterval: 0,
    variables: { id: props.scheduleID },
  })

  let label, tz
  if (error) {
    label = 'Error: ' + (error.message || error)
  } else if (!data && loading) {
    label = 'Fetching timezone information...'
  } else {
    tz = data.schedule.timeZone
    label = props.label ? props.label(tz) : `Show times in ${tz}`
  }

  return (
    <FormControlLabel
      data-cy='tz-switch'
      control={
        <Switch
          checked={zone !== 'local'}
          onChange={(e) => setZone(e.target.checked ? tz : 'local')}
          value={tz}
          disabled={Boolean(loading || error)}
        />
      }
      label={label}
    />
  )
}

ScheduleTZFilter.propTypes = {
  label: p.func,

  scheduleID: p.string.isRequired,

  // provided by connect
  zone: p.string,
  setZone: p.func,
}
