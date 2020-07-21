const HtmlWebpackPlugin = require('html-webpack-plugin')

const path = require('path')

module.exports = {
  entry: {
    app: './index.js'
  },
  output: {
    filename: 'index_bundle.js',
    path: __dirname + '/static'
  },
  plugins: [
    new HtmlWebpackPlugin()
  ]
}
