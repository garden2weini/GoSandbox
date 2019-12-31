const path = require('path')

module.exports = {
    // The base URL your application bundle will be deployed.
    publicPath: process.env.NODE_ENV === 'production'
        ? '../resources/operation/'
        //? '/resources/operation/'
        : './',
    // 构建的输出目录
    outputDir: process.env.NODE_ENV === 'production'
        ? 'dist'
        : 'dist-dev',
    // SourceMap决定部署代码是否与源代码保留映射关系(打包js时生成的map)
    productionSourceMap: process.env.NODE_ENV === 'production' ? false : true,
    
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
