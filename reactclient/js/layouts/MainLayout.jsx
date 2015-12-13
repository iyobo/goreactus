import React from 'react';
import { Link } from 'react-router';


let MainLayout = React.createClass({
    render() {
        return (
            <div className="nav">
                <h1>GoReactus</h1>
                <ul>
                    <li><Link to="/dashboard">Dashboard</Link></li>
                    <li><Link to="/about">About</Link></li>
                    <li><Link to="/logout">Log Out</Link></li>
                </ul>
                <div id="contentdiv">
                    {this.props.children}
                </div>
            </div>
        );
    }
});

export default MainLayout;