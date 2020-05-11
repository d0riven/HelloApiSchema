const path = require('path');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: {
    main: './src/main.tsx',
  },
  output: {
    path: `${__dirname}/dist`,
    filename: '[name]-[chunkhash].js'
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'HelloApiSchema',
      template: "src/index.ejs"
    }),
    new CleanWebpackPlugin({}),
  ],
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        use: [
          { loader: 'babel-loader' },
          { loader: 'ts-loader' },
        ],
      },
    ],
  },
  resolve: {
    modules: [
      path.resolve('./src'),
      path.resolve('./node_modules'),
    ],
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
  }
};