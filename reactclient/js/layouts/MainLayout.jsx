import { Link } from 'react-router';
import auth from '../auth/auth.js';

let MainLayout = React.createClass({
    getInitialState() {
        return {
            loggedIn: auth.loggedIn(),
            user: "Anonymous"
        }
    },

    updateAuth(loggedIn,username) {
        var user = this.state.user;
        if(loggedIn)
            user = auth.getUsername();
        else
            user = "Anonymous";

        this.setState({
            loggedIn: loggedIn,
            user: user
        })
    },

    componentWillMount() {
        auth.onChange = this.updateAuth
    },

    activeUsername(){
        return auth.getUsername();
    },

    render() {
        return (
            <div className="nav">

                <div className="right spaceme">
                        <span className="boldify">Hello {this.state.user}!</span>
                </div>

                <h1 className="spaceme">GoReactus</h1>
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