var webpack = require('webpack');

module.exports = {
    entry: './js/index.js',
    output: {        
        filename: './js/bundle.js'
    },
    module: {
      rules: [
        {
          test: /\.js$/,
          exclude: /node_modules/,
          use: {
            loader: 'babel-loader',
            options: {
              presets: ['env','stage-0'],
              plugins: [require('babel-plugin-transform-object-rest-spread')]
            }
          }
        }
      ]
    },
    plugins: [
        new webpack.ProvidePlugin({
           $: "jquery",
           jQuery: "jquery",
          'window.jQuery': 'jquery',
                 'jquery': 'jquery',
          'window.jquery': 'jquery',
                 '$'     : 'jquery',
          'window.$'     : 'jquery'
       })
      ],
    //   resolve: {
    //     alias: {
    //         'jquery': require.resolve('jquery'),
    //     }
    // }
} 