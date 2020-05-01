import React from 'react'
import p from 'prop-types'
import CircularProgress from '@material-ui/core/CircularProgress'
import Typography from '@material-ui/core/Typography'

import { DEFAULT_SPIN_DELAY_MS, DEFAULT_SPIN_WAIT_MS } from '../../config'

/*
 * Show a loading spinner in the center of the container.
 */
export default class Spinner extends React.PureComponent {
  static propTypes = {
    // Wait `delayMs` milliseconds before rendering a spinner.
    delayMs: p.number,

    // Wait `waitMs` before calling onReady.
    waitMs: p.number,

    // onSpin is called when the spinner starts spinning.
    onSpin: p.func,

    // onReady is called once the spinner has spun for `waitMs`.
    onReady: p.func,

    // text indicates being used as a text placeholder
    text: p.string,
  }

  static defaultProps = {
    delayMs: DEFAULT_SPIN_DELAY_MS,
    waitMs: DEFAULT_SPIN_WAIT_MS,
  }

  state = {
    spin: false,
  }

  componentDidMount() {
    this._spin = setTimeout(() => {
      this._spin = null
      this.setState({ spin: true })
      if (this.props.onSpin) this.props.onSpin()

      if (this.props.waitMs && this.props.onReady) {
        this._spin = setTimeout(this.props.onReady, this.props.waitMs)
      }
    }, this.props.delayMs)
  }

  componentWillUnmount() {
    clearTimeout(this._spin)
  }

  render() {
    if (this.props.delayMs && !this.state.spin) return null

    const style = this.props.text
      ? {
          height: '1.5em',
          color: 'gray',
          display: 'flex',
          alignItems: 'center',
        }
      : { position: 'absolute', top: '50%', left: '50%' }

    return (
      <div style={style}>
        <CircularProgress size={this.props.text ? '1em' : '40px'} />
        &nbsp;<Typography variant='body2'>{this.props.text}</Typography>
      </div>
    )
  }
}
