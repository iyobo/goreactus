/**
 * Created by Admin on 2015-12-05.
 */
var webpack = require('webpack');
module.exports = {
    entry: ['webpack/hot/dev-server', "./js/app.jsx"],
    output: {
        path: __dirname,
        filename: "bundle.js"
    },
    module: {
        loaders: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loaders: ['react-hot', 'babel?presets[]=es2015&presets[]=react'],
            },
            {
                test: /\.jsx?$/,
                //loader: 'babel',
                loaders: ['react-hot', 'babel?presets[]=es2015&presets[]=react'],
                exclude: /node_modules/,
            },
            {
                test: /\.css$/,
                loader: "style!css"
            },
        ]
    },
    plugins: [
        //new webpack.NoErrorsPlugin(),
        new webpack.HotModuleReplacementPlugin(),
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery",
            "window.jQuery": "jquery",
            React: "react",
            reactMixin: "react-mixin"
        })
    ]
};