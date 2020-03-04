import React, { useState } from 'react'
import { PropTypes as p } from 'prop-types'
import {
  TextField,
  makeStyles,
  MenuItem,
  ListItemIcon,
  Typography,
  Paper,
  Chip,
} from '@material-ui/core'
import { Autocomplete } from '@material-ui/lab'

const useStyles = makeStyles({
  listItemIcon: {
    minWidth: 0,
  },
  menuItem: {
    display: 'flex',
    flex: 1,
    justifyContent: 'space-between',
    wordBreak: 'break-word',
    whiteSpace: 'pre-wrap',
  },
  option: {
    padding: 0,
  },
  clearIndicator: {
    display: 'none',
  },
})

export default function MaterialSelect(props) {
  const classes = useStyles()
  const {
    disabled,
    error,
    isLoading,
    label,
    multiple,
    name,
    noOptionsText,
    onChange,
    onInputChange,
    options,
    placeholder,
    required,
    value: propsValue,
  } = props

  let value = propsValue
  if (!value) {
    if (multiple) value = []
    else value = { label: '', value: '' }
  }

  const [inputValue, setInputValue] = useState(multiple ? '' : value.label)

  return (
    <Autocomplete
      data-cy='material-select'
      data-cy-ready={!isLoading}
      classes={{
        option: classes.option,
        clearIndicator: classes.clearIndicator,
      }}
      value={multiple ? value : value.value}
      inputValue={inputValue}
      disableClearable={required}
      disabled={disabled}
      multiple={multiple}
      filterSelectedOptions
      noOptionsText={noOptionsText}
      onChange={(event, valueObj) => {
        if (valueObj !== null) {
          if (multiple) {
            setInputValue('')
          } else {
            setInputValue(valueObj.isCreate ? valueObj.value : valueObj.label)
          }
        }

        onChange(valueObj)
      }}
      onInputChange={(event, inputVal, reason) => {
        if (reason === 'clear' && !multiple) {
          setInputValue('')
        }
      }}
      onBlur={() => setInputValue(multiple ? '' : value.label)}
      loading={isLoading}
      getOptionLabel={option => option.label || ''}
      options={options}
      renderInput={params => {
        return (
          <TextField
            {...params}
            inputProps={{
              ...params.inputProps,
              name,
            }}
            InputProps={{
              ...params.InputProps,
              'data-cy': 'search-select-input',
            }}
            data-cy='search-select'
            fullWidth
            label={label}
            placeholder={placeholder}
            onChange={({ target }) => {
              const newInputVal = target.value
              setInputValue(newInputVal)
              if (onInputChange) onInputChange(newInputVal)
            }}
            error={error}
          />
        )
      }}
      renderOption={({ label, value, icon }) => (
        <MenuItem component='span' className={classes.menuItem}>
          <Typography noWrap>{label}</Typography>
          {icon && (
            <ListItemIcon className={classes.listItemIcon}>{icon}</ListItemIcon>
          )}
        </MenuItem>
      )}
      PaperComponent={params => <Paper data-cy='select-dropdown' {...params} />}
      renderTags={(value, getTagProps) =>
        value.map((option, index) => (
          <Chip
            key={index.toString()}
            data-cy='multi-value'
            label={option.label}
            {...getTagProps({ index })}
          />
        ))
      }
    />
  )
}

const valueShape = p.shape({
  label: p.string.isRequired,
  value: p.string.isRequired,
  icon: p.element,
})

// valueCheck ensures the type is `arrayOf(p.string)` if `multiple` is set
// and `p.string` otherwise.
function valueCheck(props, ...args) {
  if (props.multiple) return p.arrayOf(valueShape).isRequired(props, ...args)
  return valueShape(props, ...args)
}

MaterialSelect.propTypes = {
  multiple: p.bool, // allow selecting multiple values
  required: p.bool,
  onChange: p.func.isRequired,
  onInputChange: p.func,
  options: p.arrayOf(valueShape).isRequired,
  placeholder: p.string,
  value: valueCheck,
}

MaterialSelect.defaultProps = {
  options: [],
}
