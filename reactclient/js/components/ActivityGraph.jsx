var RadarChart = require("react-chartjs").Radar;

let ActivityChart = React.createClass({

    getInitialState() {
        return {
            data: {
                labels: ["Biking", "Coding", "Eating", "Plotting", "Swimming", "Yelling", "Yoga"],
                datasets: [
                    {
                        label: "Activity Dataset",
                        fillColor: "rgba(151,187,205,0.2)",
                        strokeColor: "rgba(151,187,205,1)",
                        pointColor: "rgba(151,187,205,1)",
                        pointStrokeColor: "#fff",
                        pointHighlightFill: "#fff",
                        pointHighlightStroke: "rgba(151,187,205,1)",
                        data: [0, 0, 0, 0, 0, 0, 0]
                    }
                ]
            },
            options: {},
            count: 7
        };
    },
    refresh(result){
        //curate the data
        var data = this.state.data

        var keys = []
        var values = []
        for (var i in result) {
            keys.push(i);
            values.push(result[i]);
        }

        data.labels = keys
        data.datasets[0].data = values;

        this.setState({
            data: data
        });
    },

    componentWillMount() {
        $.get(SERVER_URL + "/v1/activity", function (result) {

                this.refresh(result);

        }.bind(this));
    },

    handleClick(e){
        var name = $(e.target).attr("data-activityname");
        if (name == undefined)
            name = $(e.target.parentElement).attr("data-activityname"); //React adds annoying spans everywhere

        console.log(name);

        //log activity on server
        $.post(SERVER_URL + "/v1/activity", {name: name}, function (result) {
            this.refresh(result);

        }.bind(this));
    },

    render() {

        var buttons = [];
        for (var i = 0; i < this.state.count; i++) {
            var label = this.state.data.labels[i];
            buttons.push(
                <span key={"key_"+label}>
                    <button className="actButtons" onClick={this.handleClick} data-activityname={label}> {label}</button>
                </span>
            );
        }

        return (
            <div>
                <h3>Activity Graph</h3>
                <p>Iyobo Eki would like to know which of these is the most popular Activity in your company. Click an
                    activity as many times as you did it today. (e.g. if you ate 3 times today, click "Eating" 3
                    times)</p>

                {buttons}
                <RadarChart data={this.state.data} options={this.state.options}/>
            </div>
        )
    }
});

export default ActivityChart;