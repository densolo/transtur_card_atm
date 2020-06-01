const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin')


module.exports = {
    entry: './src/main.transtur_atm.tsx',
    mode: 'development',
    output: {
        filename: 'main.transtur_atm.bundle.[hash].js',
        path: path.resolve(__dirname, 'dist/resources/app')
    },
    devServer: {
            contentBase: path.resolve(__dirname, 'dist/resources/app'),
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: 'htmlpage/index.html.template',
            inject: false
        })
    ],
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.html$/i,
                loader: 'html-loader',
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            },
        ],
    },
    resolve: {
        modules: [
            path.resolve(__dirname, 'src'),
            path.resolve(__dirname, 'node_modules')
        ],
        extensions: [ '.tsx', '.ts', '.js' ],
    }
};
