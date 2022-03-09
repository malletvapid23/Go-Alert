import React, { ReactNode, ReactElement, forwardRef } from 'react'
import Avatar from '@mui/material/Avatar'
import List from '@mui/material/List'
import ListItem from '@mui/material/ListItem'
import ListItemText from '@mui/material/ListItemText'
import Typography from '@mui/material/Typography'
import ListItemAvatar from '@mui/material/ListItemAvatar'
import makeStyles from '@mui/styles/makeStyles'
import InfiniteScroll from 'react-infinite-scroll-component'
import { useIsWidthDown } from '../util/useWidth'
import { FavoriteIcon } from '../util/SetFavoriteButton'
import { ITEMS_PER_PAGE } from '../config'
import Spinner from '../loading/components/Spinner'
import { CheckboxItemsProps } from './ControlledPaginatedList'
import AppLink, { AppLinkProps } from '../util/AppLink'
import { debug } from '../util/debug'
import useStatusColors from '../theme/useStatusColors'

// gray boxes on load
// disable overflow
// can go to last page + one if loading & hasNextPage
// delete on details -> update list (cache, refetch?)
// - on details, don't have accesses to search param

const useStyles = makeStyles(() => ({
  infiniteScrollFooter: {
    display: 'flex',
    justifyContent: 'center',
    padding: '0.25em 0 0.25em 0',
  },
  itemAction: {
    paddingLeft: 14,
  },
  itemText: {
    wordBreak: 'break-word',
  },
  favoriteIcon: {
    backgroundColor: 'transparent',
  },
}))

const loadingStyle = {
  color: 'lightgrey',
  background: 'lightgrey',
  height: '10.3333px',
}

const useLoadingStyles = makeStyles({
  item: {
    display: 'block',
    minHeight: (dense) => (dense ? 57 : 71),
  },
  lineOne: {
    ...loadingStyle,
    width: '50%',
  },
  lineTwo: {
    ...loadingStyle,
    width: '35%',
    margin: '5px 0 5px 0',
  },
  lineThree: {
    ...loadingStyle,
    width: '65%',
  },
})

// LoadingItem is used as a placeholder for loading content
function LoadingItem(props: { dense?: boolean }): JSX.Element {
  const classes = useLoadingStyles(props.dense)

  return (
    <ListItem className={classes.item} dense={props.dense}>
      <ListItemText className={classes.lineOne} />
      <ListItemText className={classes.lineTwo} />
      <ListItemText className={classes.lineThree} />
    </ListItem>
  )
}

export interface PaginatedListProps {
  items: PaginatedListItemProps[] | CheckboxItemsProps[]
  itemsPerPage?: number

  pageCount?: number
  page?: number

  isLoading?: boolean
  loadMore?: (numberToLoad?: number) => void

  // disables the placeholder display during loading
  noPlaceholder?: boolean

  // provide a custom message to display if there are no results
  emptyMessage?: string

  // if set, loadMore will be called when the user
  // scrolls to the bottom of the list. appends list
  // items to the list rather than rendering a new page
  infiniteScroll?: boolean
}

export interface PaginatedListItemProps {
  url?: string
  title: string
  subText?: string
  isFavorite?: boolean
  icon?: ReactElement // renders a list item icon (or avatar)
  action?: ReactNode
  status?: 'ok' | 'warn' | 'err'
}

export function PaginatedList(props: PaginatedListProps): JSX.Element {
  const {
    items = [],
    itemsPerPage = ITEMS_PER_PAGE,
    pageCount,
    page,
    infiniteScroll,
    isLoading,
    loadMore,
    emptyMessage = 'No results',
    noPlaceholder,
  } = props

  const classes = useStyles()
  const statusColors = useStatusColors()
  const fullScreen = useIsWidthDown('md')

  function renderNoResults(): ReactElement {
    return (
      <ListItem>
        <ListItemText
          disableTypography
          secondary={<Typography variant='caption'>{emptyMessage}</Typography>}
        />
      </ListItem>
    )
  }

  function renderItem(item: PaginatedListItemProps, idx: number): ReactElement {
    let favIcon = null
    if (item.isFavorite) {
      favIcon = (
        <div className={classes.itemAction}>
          <Avatar className={classes.favoriteIcon}>
            <FavoriteIcon />
          </Avatar>
        </div>
      )
    }

    const borderColor = (s?: string): string => {
      switch (s) {
        case 'ok':
        case 'warn':
        case 'err':
          return statusColors[s]

        default:
          return 'transparent'
      }
    }

    const AppLinkListItem = forwardRef<HTMLAnchorElement, AppLinkProps>(
      (props, ref) => (
        <li>
          <AppLink ref={ref} {...props} />
        </li>
      ),
    )
    AppLinkListItem.displayName = 'AppLinkListItem'

    // must be explicitly set when using, in accordance with TS definitions
    const urlProps = item.url && {
      component: AppLinkListItem,

      // NOTE button: false? not assignable to true
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      button: true as any,
      to: item.url,
    }

    return (
      <ListItem
        sx={{
          borderLeft: `3px solid ${borderColor(item.status)}`,
        }}
        dense={!fullScreen}
        key={'list_' + idx}
        {...urlProps}
      >
        {item.icon && <ListItemAvatar>{item.icon}</ListItemAvatar>}
        <ListItemText
          className={classes.itemText}
          primary={item.title}
          secondary={item.subText}
        />
        {favIcon}
        {item.action && <div className={classes.itemAction}>{item.action}</div>}
      </ListItem>
    )
  }

  function renderListItems(): ReactElement | ReactElement[] {
    if (pageCount === 0 && !isLoading) return renderNoResults()

    let newItems: Array<PaginatedListItemProps> = items.slice()
    if (!infiniteScroll && page !== undefined) {
      newItems = items.slice(page * itemsPerPage, (page + 1) * itemsPerPage)
    }
    const renderedItems: ReactElement[] = newItems.map(renderItem)

    // Display full list when loading
    if (!noPlaceholder && isLoading) {
      while (renderedItems.length < itemsPerPage) {
        renderedItems.push(
          <LoadingItem
            dense={!fullScreen}
            key={'list_' + renderedItems.length}
          />,
        )
      }
    }

    return renderedItems
  }

  function renderList(): ReactElement {
    return <List data-cy='apollo-list'>{renderListItems()}</List>
  }

  function renderAsInfiniteScroll(): ReactElement {
    const len = items.length

    return (
      <InfiniteScroll
        hasMore={Boolean(loadMore)}
        next={
          loadMore ||
          (() => {
            debug('next callback missing from InfiniteScroll')
          })
        }
        scrollableTarget='content'
        endMessage={
          len === 0 ? null : (
            <Typography
              className={classes.infiniteScrollFooter}
              color='textSecondary'
              variant='body2'
            >
              Displaying all results.
            </Typography>
          )
        }
        loader={
          <div className={classes.infiniteScrollFooter}>
            <Spinner text='Loading...' />
          </div>
        }
        dataLength={len}
      >
        {renderList()}
      </InfiniteScroll>
    )
  }

  return (
    <React.Fragment>
      {infiniteScroll ? renderAsInfiniteScroll() : renderList()}
    </React.Fragment>
  )
}
