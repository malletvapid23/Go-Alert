import React, { useState } from 'react'
import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core/styles'
import gql from 'graphql-tag'
import { chain, isEmpty } from 'lodash-es'
import AdminDialog from './AdminDialog'
import PageActions from '../util/PageActions'
import { Form } from '../forms'
import AdminSection from './AdminSection'
import { useQuery } from '@apollo/react-hooks'
import Spinner from '../loading/components/Spinner'

const query = gql`
  query getLimits {
    systemLimits {
      id
      description
      value
    }
  }
`
const mutation = gql`
  mutation($input: [SystemLimitInput!]!) {
    setSystemLimits(input: $input)
  }
`

const useStyles = makeStyles(theme => ({
  gridContainer: {
    [theme.breakpoints.up('md')]: {
      justifyContent: 'center',
    },
  },
  gridItem: {
    [theme.breakpoints.up('md')]: {
      maxWidth: '65%',
    },
  },
  groupTitle: {
    fontSize: '1.1rem',
  },
  saveDisabled: {
    color: 'rgba(255, 255, 255, 0.5)',
  },
}))

export default function AdminLimits() {
  const classes = useStyles()
  const [confirm, setConfirm] = useState(false)
  const [values, setValues] = useState({})

  const { data, loading } = useQuery(query)

  if (loading && !data) {
    return <Spinner />
  }

  const updateValue = (id, value) => {
    const newVal = { ...values }

    if (value === null) {
      delete newVal[id]
    } else {
      newVal[id] = value
    }

    setValues(newVal)
  }

  const renderPageActions = () => {
    return (
      <PageActions>
        <Button
          color='inherit'
          data-cy='reset'
          disabled={isEmpty(values)}
          onClick={() => setValues({})}
          classes={{
            label: isEmpty(values) ? classes.saveDisabled : null,
          }}
        >
          Reset
        </Button>
        <Button
          color='inherit'
          data-cy='save'
          disabled={isEmpty(values)}
          onClick={() => setConfirm(true)}
          classes={{
            label: isEmpty(values) ? classes.saveDisabled : null,
          }}
        >
          Save
        </Button>
      </PageActions>
    )
  }

  const render = () => {
    return (
      <div>
        <Grid container spacing={2} className={classes.gridContainer}>
          <Grid container item xs={12} className={classes.gridItem}>
            <Grid item xs={12}>
              <Typography
                component='h2'
                variant='subtitle1'
                color='textSecondary'
                classes={{ subtitle1: classes.groupTitle }}
              >
                System Limits
              </Typography>
            </Grid>
          </Grid>
          <Grid item xs={12} className={classes.gridItem}>
            <Grid item xs={12}>
              <Form>
                <Card>
                  <AdminSection
                    value={values}
                    onChange={(id, value) => updateValue(id, value)}
                    fields={data.systemLimits.map(f => ({
                      id: f.id,
                      type: 'integer',
                      description: f.description,
                      value: f.value.toString(),
                      label: chain(f.id.replace(/([a-z])([A-Z])/g, '$1 $2'))
                        .startCase()
                        .value(),
                      password: false,
                    }))}
                  />
                </Card>
              </Form>
            </Grid>
          </Grid>
        </Grid>
        {renderPageActions()}
        {confirm && (
          <AdminDialog
            mutation={mutation}
            values={data.systemLimits}
            fieldValues={values}
            onClose={() => setConfirm(false)}
            onComplete={() => {
              setValues({})
              setConfirm(false)
            }}
          />
        )}
      </div>
    )
  }

  return render()
}
