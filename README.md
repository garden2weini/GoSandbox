数据大屏调研
===================
基于Vue开发Data View
-------------------

# 快速开始
1. clone the respository from git.
1. npm install(to get node_modules).
1. npm run serve

# 运行准备
- <b>重点：</b><font color="red">(虽然走了点弯路才发现)不用再照着G2/Echarts重新改写各框架的Chart代码了！</font>
- See [Viser Demo for React/Vue/Angular](https://viserjs.github.io/demo.html).

```
# Viser (for chart)
npm install viser-vue
# ElementUI
npm i element-ui -S
# DataV g2(+insert-css)
npm install --save @antv/g2
npm install --save @antv/data-set
npm install insert-css
# Echarts
npm install echarts --save
# 安装JQuery
npm install jquery --save
# 安装Bootstrap
npm install bootstrap --save
npm install --save popper.js
```
在main.js文件中添加如下内容
```
import $ from 'jquery'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'
```

# 基础知识-Project setup
```
npm install
```

## Compiles and hot-reloads for development
```
npm run serve
```

## Compiles and minifies for production
```
npm run build
```

## Run your tests
```
npm run test
```

## Lints and fixes files
```
npm run lint
```

# Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
