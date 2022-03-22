import React, { useState, ReactNode } from 'react'
import {
  Badge,
  Collapse,
  Grid,
  IconButton,
  Alert,
  AlertTitle,
  AlertProps,
} from '@mui/material'
import makeStyles from '@mui/styles/makeStyles'
import ExpandIcon from '@mui/icons-material/KeyboardArrowDown'
import CollapseIcon from '@mui/icons-material/KeyboardArrowUp'
import toTitleCase from '../util/toTitleCase'

const useStyles = makeStyles({
  alertAction: {
    marginRight: 0,
  },
  alertMessage: {
    width: '100%',
  },
  gridItem: {
    padding: '4px 0 4px 0',
  },
  lastGridItem: {
    padding: '4px 0 0 0',
  },
  secondGridItem: {
    padding: '8px 0 4px 0',
  },
})

export interface Notice {
  type: NoticeType
  message: string | JSX.Element
  details?: string | JSX.Element
  endNote?: string | JSX.Element
}

export type NoticeType = 'WARNING' | 'ERROR' | 'INFO' | 'OK'

interface NoticesProps {
  notices?: Notice[]
}

export default function Notices({
  notices = [],
}: NoticesProps): JSX.Element | null {
  const classes = useStyles()
  const [noticesExpanded, setNoticesExpanded] = useState(false)

  if (!notices.length) {
    return null
  }

  function renderShowAllToggle(): ReactNode {
    if (notices.length <= 1) return null
    return (
      <Badge
        color='primary'
        badgeContent={notices.length - 1}
        invisible={noticesExpanded}
      >
        <IconButton
          onClick={() => setNoticesExpanded(!noticesExpanded)}
          size='large'
        >
          {noticesExpanded ? <CollapseIcon /> : <ExpandIcon />}
        </IconButton>
      </Badge>
    )
  }

  /*
   * Spacing set manually on grid items to accommodate manual
   * accordion transitions for multiple notices
   */
  function getGridClassName(index: number): string {
    switch (index) {
      case 0:
        return ''
      case 1:
        return classes.secondGridItem
      case notices.length - 1:
        return classes.lastGridItem
      default:
        return classes.gridItem
    }
  }

  function renderNotice(notice: Notice, index: number): JSX.Element {
    return (
      <Grid key={index} className={getGridClassName(index)} item xs={12}>
        <Alert
          severity={notice.type.toLowerCase() as AlertProps['severity']}
          classes={{
            message: classes.alertMessage,
            action: classes.alertAction,
          }}
          action={index === 0 ? renderShowAllToggle() : null}
        >
          <AlertTitle>
            {toTitleCase(notice.type)}: {notice.message}
          </AlertTitle>
          {notice.details}
          {notice.endNote && (
            <div style={{ float: 'right' }}>{notice.endNote}</div>
          )}
        </Alert>
      </Grid>
    )
  }

  return (
    <Grid container>
      {renderNotice(notices[0], 0)}
      <Grid item xs={12}>
        <Collapse in={noticesExpanded}>
          <Grid container>
            {notices.slice(1).map((n, i) => renderNotice(n, i + 1))}
          </Grid>
        </Collapse>
      </Grid>
    </Grid>
  )
}
