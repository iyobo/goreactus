import { Link } from 'react-router';
import auth from '../auth/auth.js';

let MainLayout = React.createClass({
    getInitialState() {
        return {
            loggedIn: auth.loggedIn()
        }
    },

    updateAuth(loggedIn) {
        this.setState({
            loggedIn: loggedIn
        })
    },

    componentWillMount() {
        auth.onChange = this.updateAuth
        auth.login()
    },

    render() {
        return (
            <div className="nav">
                <h1>GoReactus</h1>
                <ul>
                    <li><Link to="/dashboard">Dashboard</Link></li>
                    <li><Link to="/about">About</Link></li>
                    <li>
                        {this.state.loggedIn ? (
                        <Link to="/logout">Log out</Link>
                            ) : (
                        <Link to="/login">Sign in</Link>
                            )}
                    </li>

                </ul>
                <div id="contentdiv">
                    {this.props.children}
                </div>
            </div>
        );
    }
});

export default MainLayout;