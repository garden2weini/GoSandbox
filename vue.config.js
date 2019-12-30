const path = require('path')

module.exports = {
    //publicPath: '/resources/operation/', // 基本路径
    //outputDir = 'dist'; // 输出文件目录
    publicPath: process.env.NODE_ENV === 'production'
        ? '${base}/resources/operation/'
        : './',
    outputDir: process.env.NODE_ENV === 'production'
        ? 'dist'
        : 'dist-dev',
    
    configureWebpack: (config) => {
        if (process.env.NODE_ENV === 'production') {
            // 为生产环境修改配置...
            config.mode = 'production';
        } else {
            // 为开发环境修改配置...
            config.mode = 'development';
        }
    }
}
