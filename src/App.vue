<template>
    <div id="app" class="container" style="margin-top: 10px;margin-left: 10px;margin-right: 20px;">
        <div class="row">
            <div class="col-md-3 col-lg-3">
                <roll-list msg="近30天TopN设备" :hasBorder="true" :data="last30DeviceList"/>
                <roll-list msg="近30天 热销SKU" :hasBorder="true" :data="last30SkuList"/>
                <roll-list msg="Top毛利*周转SKU" :hasBorder="true" :data="topSkuList"/>
                <biz-pie msg="设备量" :hasBorder="true" :data="deviceData"/>
            </div>
            <div class="col-md-5 col-lg-5">
                <topic msg="智能零售数据运营大屏"/>
                <el-divider></el-divider>
                <sale-summary></sale-summary>
                <el-divider></el-divider>
                <chart-double-line msg="活跃设备日均销售额" :hasBorder="false"/>
            </div>
            <div class="col-md-4 col-lg-4">
                <chart-column msg="近30天 用户复购" :hasBorder="true"/>
                <two-colums msg="服务用户" :hasBorder="true"/>
                <three-colums msg="客单价" :hasBorder="true"/>
                <biz-ratio-bar msg="支付占比" :hasBorder="true"/>
                <refund msg="近30天退款金额" :hasBorder="true"/>
            </div>
        </div>
        <div class="row">
           <div style="position: absolute; bottom: 5px;left: 50%;transform: translate(-50%, -50%);">
               <timer/>
           </div> 
        </div>
    </div>
</template>

<script>
    import $ from 'jquery'
    import ElementUI from 'element-ui';
    import 'element-ui/lib/theme-chalk/index.css';
    
    // Common components.
    import Topic from './components/Topic.vue'
    import RollList from './components/RollList.vue'
    import TwoColums from './components/TwoColums.vue'
    import ThreeColums from './components/ThreeColums.vue'
    // Biz components.
    import BizPie from './components/biz/BizPie.vue'
    import BizRatioBar from './components/biz/BizRatioBar.vue'
    import SaleSummary from './components/biz/Summary.vue'
    import Refund from './components/biz/Refund.vue'
    // Raw Chart components.
    import ChartDoubleLine from './components/chart/DoubleLine.vue'
    import ChartColumn from './components/chart/Column.vue'
    import Timer from './components/chart/Timer.vue'
    
    export default {
        name: 'app',
        components: {
            "chart-double-line" : ChartDoubleLine,
            BizPie,
            ChartColumn,
            Timer,
            Topic,
            RollList,
            TwoColums,
            ThreeColums,
            BizRatioBar,
            SaleSummary,
            Refund,
        },
        data() {
            // 定义变量
            return {
                deviceData: null,
                last30DeviceList: null,
                last30SkuList: null,
                topSkuList: null,
                }
        },
        mounted() {
            var self = this;
            // NOTE 远程获取rest数据
            function updateViewData() {
                $.getJSON(self.DATA_BASE_URL + 'device.json', (sourceData) => {
                      self.deviceData = sourceData;
                });
                $.getJSON(self.DATA_BASE_URL + 'last30-device.json', (sourceData) => {
                      var list = sourceData.list;
                      var results = new Array(list.length);
                      // 格式化List内容
                      for(var i=0; i< list.length; i++) {
                          results[i] = list[i].name + ":" + list[i].amt + "/" + list[i].cnt + "笔";
                      }
                      self.last30DeviceList = results;
                });
                $.getJSON(self.DATA_BASE_URL + 'last30-sku.json', (sourceData) => {
                      var list = sourceData.list;
                      var results = new Array(list.length);
                      // 格式化List内容
                      for(var i=0; i< list.length; i++) {
                          results[i] = list[i].name + ":" + list[i].rate;
                      }
                      self.last30SkuList = results;
                });
                $.getJSON(self.DATA_BASE_URL + 'top-sku.json', (sourceData) => {
                      var list = sourceData.list;
                      var results = new Array(list.length);
                      // 格式化List内容
                      for(var i=0; i< list.length; i++) {
                          results[i] = list[i];
                      }
                      self.topSkuList = results;
                });
            }
            function updateListData() {
                $.getJSON(self.DATA_BASE_URL + 'last30-device.json', (sourceData) => {
                      var list = sourceData.list;
                      var results = new Array(list.length);
                      // 格式化List内容
                      for(var i=0; i< list.length; i++) {
                          results[i] = list[i].name + ":" + list[i].amt + "/" + list[i].cnt + "笔";
                      }
                      self.last30DeviceList = results;
                });
                $.getJSON(self.DATA_BASE_URL + 'last30-sku.json', (sourceData) => {
                      var list = sourceData.list;
                      var results = new Array(list.length);
                      // 格式化List内容
                      for(var i=0; i< list.length; i++) {
                          results[i] = list[i].name + ":" + list[i].rate;
                      }
                      self.last30SkuList = results;
                });
                $.getJSON(self.DATA_BASE_URL + 'top-sku.json', (sourceData) => {
                      var list = sourceData.list;
                      var results = new Array(list.length);
                      // 格式化List内容
                      for(var i=0; i< list.length; i++) {
                          results[i] = list[i];
                      }
                      self.topSkuList = results;
                });
            }
            updateViewData();
            updateListData();
            setInterval(updateViewData, this.REFRESH_INTERVAL);
            setInterval(updateListData, 5000); // 60 * 1000
        }
    }
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }
</style>
