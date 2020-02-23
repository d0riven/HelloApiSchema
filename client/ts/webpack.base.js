const {CleanWebpackPlugin} = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: {
    main: './src/main.ts',
  },
  output: {
    path: `${__dirname}/dist`,
    filename: '[name]-[chunkhash].js'
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'HelloApiSchema',
      filename: 'index.html',
    }),
    new CleanWebpackPlugin({}),
  ],
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: [
          { loader: 'babel-loader' },
          { loader: 'ts-loader' },
        ],
      },
    ],
  },
  resolve: {
    extensions: ['.ts', '.js', '.json']
  }
};