
import React from 'react';
import ReactDOM from 'react-dom';
import { DefaultRoute, Link, Route, Router, RouteHandler, IndexRoute } from 'react-router';

import MainLayout from './layouts/MainLayout.jsx';
import LoginPage from './pages/LoginPage.jsx';
import LogoutPage from './pages/LogoutPage.jsx';
import AboutPage from './pages/AboutPage.jsx';
import Error404Page from './pages/Error404Page.jsx';
import DashboardPage from './pages/DashboardPage.jsx';
import auth from './auth/auth.js';
import history from './history.jsx';



function requireAuth(nextState, replaceState) {
    if (!auth.loggedIn())
        replaceState({ nextPathname: nextState.location.pathname }, '/login')
}

ReactDOM.render((
    <Router history={history}>
        <Route path="/" component={MainLayout}>
            <IndexRoute component={DashboardPage} onEnter={requireAuth} />
            <Route path="login" component={LoginPage}/>
            <Route path="logout" component={LogoutPage}/>
            <Route path="about" component={AboutPage}/>
            <Route path="dashboard" component={DashboardPage} onEnter={requireAuth} />
            <Route path="*" component={Error404Page}/>
        </Route>

    </Router>
), document.getElementById("mainframe"));