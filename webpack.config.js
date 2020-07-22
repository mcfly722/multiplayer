const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path')

var webpack = require('webpack');

module.exports = {
  entry: {
    app: './index.js'
  },
  output: {
    filename: 'index_bundle.js',
    path: __dirname + '/static'
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'multiplayer'
    }),
    new webpack.ProvidePlugin({
      $: "jquery",
      jQuery: "jquery"
    })
  ]
}
