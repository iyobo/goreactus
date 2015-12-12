import React from 'react';
import ReactDOM from 'react-dom';
import { DefaultRoute, Link, Route, Router, RouteHandler, IndexRoute } from 'react-router';

import LoginPage from './pages/Login.jsx';

let App = React.createClass({
    render() {
        return (
            <div className="nav">
                <h1>App</h1>
                <ul>
                    <li><Link to="/login">Login</Link></li>
                    <li><Link to="/about">About Us</Link></li>
                </ul>
                {this.props.children}
            </div>
        );
    }
});

let Login = React.createClass({
    render() {
        return (
            <div className="nav">
                <div> We're Login in!</div>
            </div>
        );
    }
});
let About = React.createClass({
    render() {
        return (
            <div className="nav">
                <div> We are the best.</div>
            </div>
        );
    }
});
let Error404Page = React.createClass({
    render() {
        return (
            <div className="nav">
                <div> Nothing here</div>
            </div>
        );
    }
});

ReactDOM.render((
    <Router>
        <Route path="/" component={App}>
            <IndexRoute component={LoginPage} />
            <Route path="login" component={LoginPage}/>
            <Route path="about" component={About}/>
            <Route path="*" component={Error404Page}/>
        </Route>

    </Router>
), document.getElementById("mainframe"))