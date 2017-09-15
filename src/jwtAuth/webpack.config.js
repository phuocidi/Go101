var webpack = require("webpack")
const path = require('path');


module.exports = {
	entry: "./static/index.js",
	output: {
		path: "/home/huutran/go/src/jwtAuth/static/",
		filename: "./js/bundle.js",
		publicPath: "./static"
	},

	resolve: {
        modules: ['static','node_modules'],
    },
	module: {
		loaders: [
			{
				test: /\.js$/,
				exclude: /(node_modules)/,
				loader: "babel-loader",
				query: {
					compact: true,
					presets: [ "stage-0","react","es2015"]
				}
			},
		]
	}
}
