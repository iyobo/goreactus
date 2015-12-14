import auth from '../auth/auth.js'
import history from '../history.jsx'

let LoginPage = React.createClass({

    getInitialState() {
        return {
            username: "",
            password: "",
            error: ""
        };
    },

    // This will be called when the user clicks on the login button
    handleLogin(e) {
        e.preventDefault();
        var me = this;
        console.log("Authenticating with server...");
        // Here, we call an external AuthService. We"ll create it in the next step
        auth.login(this.state.username, this.state.password, function (status) {
            console.log("loggedin: " + status);
            if (!status) {
                me.setState({
                    error: "Login: Invalid Credentials"
                });
            }else{
                me.setState({
                    error: ""
                });
                //history.replaceState(null, '#/dashboard')
                window.location = '#/dashboard'
            }
        });
    },

    handleUsernameChange(e){
        this.setState({
            username: e.target.value
        });
    },

    handlePasswordChange(e){
        this.setState({
            password: e.target.value
        });
    },

    render() {
        return (
            <div id="loginform">
                <h1>Login</h1>
                <form role="form" onSubmit={this.handleLogin}>
                    {this.state.error==="" ? (
                    <span></span>
                        ) : (
                    <div className="alert alert-warning" role="alert">{this.state.error}</div>
                        )}

                    <div className="form-group">
                        <div className="form-group">
                            <label htmlFor="usr">Username:</label>
                            <input className="form-control" type="text" value={this.state.username}
                                   onChange={this.handleUsernameChange} placeholder="Username"/>
                        </div>
                        <div className="form-group">
                            <label htmlFor="usr">Password:</label>
                            <input className="form-control" type="password" value={this.state.password}
                                   onChange={this.handlePasswordChange} placeholder="Password"/>
                        </div>
                    </div>
                    <button className="btn" type="submit">Login</button>
                </form>

                <div>
                    <br/>
                    <h4>Try...</h4>
                    <p>iyobo / password</p>
                    <p>wolverine / password</p>
                    <p>phoenix / password</p>
                    <p>jack / password</p>
                    <p>adam / password</p>
                </div>
            </div>
        );
    }
});

export default LoginPage;