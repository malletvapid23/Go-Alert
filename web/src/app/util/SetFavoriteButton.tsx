import React from 'react'
import IconButton from '@material-ui/core/IconButton'
import FavoriteFilledIcon from '@material-ui/icons/Star'
import FavoriteBorderIcon from '@material-ui/icons/StarBorder'
import Tooltip from '@material-ui/core/Tooltip'
import Spinner from '../loading/components/Spinner'

interface SetFavoriteButtonProps {
  typeName: 'rotation' | 'service' | 'schedule'
  isFavorite?: boolean
  loading: boolean
  onClick: Function

  tooltipOptions?: {
    enabled: boolean
    setMessage?: string // The tooltip displayed when the favorite has already been set
    unsetMessage?: string // The tooltip displayed when the favorite has not yet been set
  }
}

export function SetFavoriteButton({
  typeName,
  isFavorite,
  loading,
  onClick,
  tooltipOptions,
}: SetFavoriteButtonProps): JSX.Element {
  let icon = isFavorite ? <FavoriteFilledIcon /> : <FavoriteBorderIcon />
  if (loading) {
    icon = <Spinner />
  }

  const content = (
    <form
      onSubmit={(e) => {
        e.preventDefault()
        onClick()
      }}
    >
      <IconButton
        aria-label={
          isFavorite
            ? `Unset as a Favorite ${typeName}`
            : `Set as a Favorite ${typeName}`
        }
        type='submit'
        color='inherit'
        data-cy='set-fav'
      >
        {icon}
      </IconButton>
    </form>
  )

  const { enabled, setMessage, unsetMessage } = tooltipOptions ?? {}
  if (enabled) {
    return (
      <Tooltip title={isFavorite ? unsetMessage : setMessage}>
        {content}
      </Tooltip>
    )
  }
  return content
}
