import React, { useState } from 'react'
import { useQuery, gql } from '@apollo/client'
import { Box, Grid, Typography } from '@mui/material'
import makeStyles from '@mui/styles/makeStyles'
import { GenericError } from '../../error-pages'
import Spinner from '../../loading/components/Spinner'
import OutgoingLogCard from './OutgoingLogCard'
import Search from '../../util/Search'
import OutgoingLogsFilter from './OutgoingLogsFilter'
import OutgoingLogDetails from './OutgoingLogDetails'
import { theme } from '../../mui'
import { DebugMessage } from '../../../schema'

const debugMessageLogsQuery = gql`
  query debugMessageLogsQuery {
    debugMessages {
      id
      createdAt
      updatedAt
      type
      status
      userID
      userName
      source
      destination
      serviceID
      serviceName
      alertID
      providerID
    }
  }
`

const useStyles = makeStyles<typeof theme>((theme) => ({
  gridContainer: {
    [theme.breakpoints.up('md')]: {
      justifyContent: 'center',
    },
  },
  groupTitle: {
    fontSize: '1.1rem',
  },
  saveDisabled: {
    color: 'rgba(255, 255, 255, 0.5)',
  },
  card: {
    margin: theme.spacing(1),
    cursor: 'pointer',
  },
}))

export default function AdminOutgoingLogs(): JSX.Element {
  const classes = useStyles()
  const [selectedLog, setSelectedLog] = useState<DebugMessage | null>(null)

  const { data, loading, error } = useQuery(debugMessageLogsQuery)

  if (error) return <GenericError error={error.message} />
  if (loading && !data) return <Spinner />

  return (
    <React.Fragment>
      <OutgoingLogDetails
        open={Boolean(selectedLog)}
        onClose={() => setSelectedLog(null)}
        log={selectedLog}
      />
      <Grid container spacing={2} className={classes.gridContainer}>
        <Grid container item xs={12}>
          <Grid item xs={12}>
            <Typography
              component='h2'
              variant='subtitle1'
              color='textSecondary'
              classes={{ subtitle1: classes.groupTitle }}
            >
              Outgoing Messages
            </Typography>
          </Grid>
          <Grid item xs={12}>
            <Box
              display='flex'
              flexDirection='row'
              alignItems='center'
              justifyContent='space-between'
            >
              <div>
                <OutgoingLogsFilter />
              </div>
              <div>
                <Search />
              </div>
            </Box>
          </Grid>
          <Grid item xs={12}>
            <Box
              display='flex'
              flexDirection='column'
              alignItems='stretch'
              width='full'
            >
              {/* TODO: change card's outline color in list when selected */}
              {data.debugMessages.map((debugMessage) => (
                <OutgoingLogCard
                  key={debugMessage.id}
                  debugMessage={debugMessage}
                  onClick={() => setSelectedLog(debugMessage)}
                />
              ))}
            </Box>
          </Grid>
        </Grid>
      </Grid>
    </React.Fragment>
  )
}
