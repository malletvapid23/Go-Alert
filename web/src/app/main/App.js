import React, { useState } from 'react'
import AppBar from '@material-ui/core/AppBar'
import Hidden from '@material-ui/core/Hidden'
import Toolbar from '@material-ui/core/Toolbar'
import ToolbarTitle from './components/ToolbarTitle'
import ToolbarAction from './components/ToolbarAction'
import ErrorBoundary from './ErrorBoundary'
import routeConfig, { renderRoutes } from './routes'
import { Switch, Route } from 'react-router-dom'
import Grid from '@material-ui/core/Grid'
import { useSelector } from 'react-redux'
import { authSelector } from '../selectors'
import { PageActionContainer, PageActionProvider } from '../util/PageActions'
import { PageNotFound as LazyPageNotFound } from '../error-pages/Errors'
import LazySideBarDrawerList from './components/SideBarDrawerList'
import SwipeableDrawer from '@material-ui/core/SwipeableDrawer'
import LazyWideSideBar, { drawerWidth } from './WideSideBar'
import LazyNewUserSetup from './components/NewUserSetup'
import Login from './components/Login'
import { SkipToContentLink } from '../util/SkipToContentLink'
import { SearchContainer, SearchProvider } from '../util/AppBarSearchContainer'
import { makeStyles } from '@material-ui/core'
import { useIsWidthDown } from '../util/useWidth'
import { isIOS } from '../util/browsers'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    zIndex: 1,
    position: 'relative',
    display: 'flex',
    backgroundColor: 'lightgrey',
    height: '100%',
  },
  main: {
    width: '100%',
    overflowY: 'auto',
    marginTop: '64px',
  },
  mainContainer: { position: 'relative', height: '100%' },
  appBar: {
    [theme.breakpoints.up('md')]: { width: `calc(100% - ${drawerWidth})` },
    zIndex: theme.zIndex.drawer + 1,
  },
  containerClass: {
    padding: '1em',
    [theme.breakpoints.up('md')]: { width: '75%' },
    [theme.breakpoints.down('sm')]: { width: '100%' },
  },
}))

export default function App() {
  const classes = useStyles()
  const [showMobile, setShowMobile] = useState(false)
  const fullScreen = useIsWidthDown('sm')
  const marginLeft = fullScreen ? 0 : drawerWidth
  const authValid = useSelector(authSelector)

  if (!authValid) {
    return <Login />
  }

  let cyFormat = 'wide'
  if (fullScreen) cyFormat = 'mobile'
  return (
    <div className={classes.root}>
      <PageActionProvider>
        <SearchProvider>
          <AppBar
            position='fixed'
            className={classes.appBar}
            data-cy='app-bar'
            data-cy-format={cyFormat}
          >
            <SkipToContentLink />
            <Toolbar>
              <ToolbarAction
                showMobileSidebar={showMobile}
                openMobileSidebar={() => setShowMobile(true)}
              />
              <ToolbarTitle />

              <PageActionContainer />
              <SearchContainer />
            </Toolbar>
          </AppBar>

          <Hidden smDown>
            <LazyWideSideBar>
              <LazySideBarDrawerList
                closeMobileSidebar={() => setShowMobile(false)}
              />
            </LazyWideSideBar>
          </Hidden>
          <Hidden mdUp>
            <SwipeableDrawer
              disableDiscovery={isIOS}
              open={showMobile}
              onOpen={() => setShowMobile(true)}
              onClose={() => setShowMobile(false)}
            >
              <LazySideBarDrawerList
                closeMobileSidebar={() => setShowMobile(false)}
              />
            </SwipeableDrawer>
          </Hidden>

          <main id='content' className={classes.main} style={{ marginLeft }}>
            <ErrorBoundary>
              <LazyNewUserSetup />
              <Grid
                container
                justifyContent='center'
                className={classes.mainContainer}
              >
                <Grid className={classes.containerClass} item>
                  <Switch>
                    {renderRoutes(routeConfig)}
                    <Route component={() => <LazyPageNotFound />} />
                  </Switch>
                </Grid>
              </Grid>
            </ErrorBoundary>
          </main>
        </SearchProvider>
      </PageActionProvider>
    </div>
  )
}
