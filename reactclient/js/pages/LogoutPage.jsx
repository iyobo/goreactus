import auth from '../auth/auth.js'

let LogoutPage = React.createClass({

  componentDidMount() {
    auth.logout();
      this.setState({"logged":"out"})
  },

  render() {
    return <p>You are now logged out</p>
  }
});

export default LogoutPage;