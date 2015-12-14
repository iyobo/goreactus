import auth from '../auth/auth.js';
import ActivityGraph from '../components/ActivityGraph.jsx'


let DashboardPage = React.createClass({

    render() {
        return (
            <div>
                <h1>Dashboard</h1>

                <ActivityGraph />
            </div>
        );
    }
});

export default DashboardPage;