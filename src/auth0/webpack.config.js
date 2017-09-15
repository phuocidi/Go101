var webpack = require("webpack")
const path = require('path');
const Dotenv = require('dotenv-webpack');


module.exports = {
	externals: {},
	entry: "./static/index.js",
	output: {
		path: "/home/huutran/go/src/auth0/static/",
		filename: "js/bundle.js",
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
					presets: ["es2015", "react", "stage-0"]
				}
			},
		]
	},
	target: 'web',
	plugins: [ 
		new Dotenv({
      		path: './.env', // Path to .env file (this is the default) 
      		safe: false // load .env.example (defaults to "false" which does not use dotenv-safe) 
    	}),
		 new webpack.DefinePlugin({ 
  //       	'process.env.NODE_ENV': '"development"',
        	'global.GENTLY': false,
		 }),
	 ],
}
