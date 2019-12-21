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
// 注册全局组件
Vue.component("sub-topic", SubTopic)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')