import React, { useLayoutEffect } from 'react'
import classnames from 'classnames'
import MUIListItem, {
  ListItemProps as MUIListItemProps,
} from '@mui/material/ListItem'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemSecondaryAction from '@mui/material/ListItemSecondaryAction'
import ListItemText from '@mui/material/ListItemText'
import makeStyles from '@mui/styles/makeStyles'
import AppLink from '../util/AppLink'
import { FlatListItem } from './FlatList'

const useStyles = makeStyles(() => ({
  listItem: {
    width: '100%',
  },
  listItemDisabled: {
    opacity: 0.6,
    width: '100%',
  },
  secondaryText: {
    whiteSpace: 'pre-line',
  },
}))

export interface FlatListItemProps extends MUIListItemProps {
  item: FlatListItem
  index: number
  showOptions?: boolean
}

export default function FlatListItem(props: FlatListItemProps): JSX.Element {
  const classes = useStyles()

  const showOptions = props?.showOptions ?? true
  const {
    disabled,
    highlight,
    icon,
    secondaryAction,
    scrollIntoView,
    subText,
    title,
    url,
    ...muiListItemProps
  } = props.item

  const ref = React.useRef<HTMLLIElement>(null)
  useLayoutEffect(() => {
    if (scrollIntoView) {
      ref.current?.scrollIntoView({ block: 'center' })
    }
  }, [scrollIntoView])

  let linkProps = {}
  if (url) {
    linkProps = {
      component: AppLink,
      to: url,
      button: true,
    }
  }

  return (
    <MUIListItem
      key={props.index}
      {...linkProps}
      {...muiListItemProps}
      className={classnames({
        [classes.listItem]: true,
        [classes.listItemDisabled]: disabled,
      })}
      selected={highlight}
    >
      {icon && <ListItemIcon>{icon}</ListItemIcon>}
      <ListItemText
        primary={title}
        secondary={subText}
        secondaryTypographyProps={{
          className: classnames({
            [classes.secondaryText]: true,
            [classes.listItemDisabled]: disabled,
          }),
        }}
      />
      {secondaryAction && showOptions && (
        <ListItemSecondaryAction sx={{ zIndex: 9002 }}>
          {secondaryAction}
        </ListItemSecondaryAction>
      )}
    </MUIListItem>
  )
}
