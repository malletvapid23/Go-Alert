import { makeStyles } from '@material-ui/core'
import React from 'react'

const karakurenai = '#cd1831'

const useStyles = makeStyles({
  button: {
    color: karakurenai,
    all: 'unset',
    backgroundColor: 'transparent',
    border: 'none',
    cursor: 'pointer',
    textDecoration: 'none',
    display: 'inline',
    margin: 0,
    padding: 0,
    '&:hover': {
      textDecoration: 'underline',
    },
    '&:focus': {
      textDecoration: 'underline',
    },
  },
})

interface ClickableTextProps {
  text: string
  onClick: () => void
}

// ClickableText
// since anchor elements have become entrenched with navigation by convention
// it is recommended that clickable text be implemented using a button element
// see https://github.com/jsx-eslint/eslint-plugin-jsx-a11y/blob/master/docs/rules/anchor-is-valid.md
function ClickableText({ text, onClick }: ClickableTextProps): JSX.Element {
  const classes = useStyles()
  return (
    <button
      onClick={(e) => {
        e.preventDefault()
        onClick()
      }}
      className={classes.button}
    >
      {text}
    </button>
  )
}

export default ClickableText
