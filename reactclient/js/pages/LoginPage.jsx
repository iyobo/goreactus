import auth from '../auth/auth.js'

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
                    error: "Invalid Credentials"
                });
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
            </div>
        );
    }
});

export default LoginPage;