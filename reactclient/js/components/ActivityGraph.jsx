var RadarChart = require("react-chartjs").Radar;

let ActivityChart = React.createClass({

    getInitialState() {
        return {
            data: {
                labels: ["Eating", "Sleeping", "Biking", "Swimming", "Sitting(hrs)", "Running", "Yelling","Yoga"],
                datasets: [
                    {
                        label: "My Second dataset",
                        fillColor: "rgba(151,187,205,0.2)",
                        strokeColor: "rgba(151,187,205,1)",
                        pointColor: "rgba(151,187,205,1)",
                        pointStrokeColor: "#fff",
                        pointHighlightFill: "#fff",
                        pointHighlightStroke: "rgba(151,187,205,1)",
                        data: [0, 0, 0, 0, 0, 0, 0,0]
                    }
                ]
            },
            options: {},
        };
    },

    componentDidMount() {
        $.get(SERVER_URL+"/v1/activity", function(result) {
            if (this.isMounted()) {

                //curate the data


                this.setState({
                    data: data,
                });
            }
        }.bind(this));
    },

    render() {
        return (
            <div>
                <h3>Activity Graph</h3>
                <p>Iyobo Eki would like to know which of these is the most popular Activity in your company. Click an activity as many times as you did it today. (e.g. if you ate 3 times today, click "Eating" 3 times)</p>

                <RadarChart data={this.state.data} options={this.state.options}/>
            </div>
        )
    }
});

export default ActivityChart;