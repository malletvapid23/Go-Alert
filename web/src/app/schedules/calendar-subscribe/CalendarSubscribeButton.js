import React, { useState } from 'react'
import { PropTypes as p } from 'prop-types'
import { Button, Grid, makeStyles, Typography } from '@material-ui/core/index'
import CalendarIcon from 'mdi-material-ui/Calendar'
import CalendarSubscribeDialog from './CalendarSubscribeDialog'

const useStyles = makeStyles(theme => ({
  calIcon: {
    marginRight: theme.spacing(1),
  },
  captionContainer: {
    display: 'grid',
  },
}))

export default function CalendarSubscribeButton(props) {
  const [showDialog, setShowDialog] = useState(false)
  const classes = useStyles()

  const caption =
    'Subscribe to your shifts on this calendar from your preferred calendar app'
  // const caption = 'You have 1 active subscription for this schedule'

  return (
    <React.Fragment>
      <CalendarSubscribeDialog
        open={showDialog}
        onClose={() => setShowDialog(false)}
        scheduleID={props.scheduleID}
      />
      <Grid container spacing={1}>
        <Grid item xs={12}>
          <Button
            color='primary'
            onClick={() => setShowDialog(true)}
            variant='contained'
          >
            <CalendarIcon className={classes.calIcon} />
            Subscribe
          </Button>
        </Grid>
        <Grid item xs={12} className={classes.captionContainer}>
          <Typography variant='caption' color='textSecondary'>
            {caption}
          </Typography>
        </Grid>
      </Grid>
    </React.Fragment>
  )
}

CalendarSubscribeButton.propTypes = {
  scheduleID: p.string,
}
