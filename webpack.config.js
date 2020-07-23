const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path')

var webpack = require('webpack');

module.exports = {
  entry: {
    app: './index.js'
  },
  output: {
    filename: 'index_bundle_[contenthash].js',
    path: __dirname + '/static'
  },
  watch: true,
  watchOptions: {
    aggregateTimeout: 300
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
