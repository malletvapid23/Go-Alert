import { combineReducers } from 'redux'
import alerts from './alerts'
import auth from './auth'
import main from './main'
import { connectRouter } from 'connected-react-router'

export default history =>
  combineReducers({
    router: connectRouter(history),

    auth, // auth status
    alerts, // reducer for filters on alerts list
    main, // reducer for new user setup flag
  })
