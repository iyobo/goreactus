
let LogoutPage = React.createClass({

  componentDidMount() {
    auth.logout();
  },

  render() {
    return <p>You are now logged out</p>
  }
});

export default LogoutPage;