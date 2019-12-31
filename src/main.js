import Vue from 'vue'
import App from './App.vue'
//import App1 from './App1.vue'
import G2 from '@antv/g2'
//import DataSet from '@antv/data-set'
import $ from 'jquery'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'
import echarts from 'echarts';

import SubTopic from './components/SubTopic.vue'
// NOTE: 注册全局组件
Vue.component("sub-topic", SubTopic)

// NOTE: 定义全局变量
Vue.prototype.BG_COLOR='#C0C4CC';
// 依照环境切换Json Url的根
if (process.env.NODE_ENV === 'production') {
    Vue.prototype.JSON_BASE = "${base}/"
} else {
    Vue.prototype.JSON_BASE = "/"
}
// Json数据跟路径
Vue.prototype.DATA_BASE_URL = Vue.prototype.JSON_BASE + "data/data-view/"
// 各组件数据刷新频率(毫秒)
Vue.prototype.REFRESH_INTERVAL = 30 * 1000

// NOTE: 设置为开发环境或者生产环境: true开发模式; false生产模式
Vue.config.productionTip = false

new Vue({
    render: h => h(App),
}).$mount('#app')
