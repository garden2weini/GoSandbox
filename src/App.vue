<template>
    <div id="app" class="container" style="margin-top: 10px;margin-left: 10px;margin-right: 20px;">
        <div class="row">
            <div class="col-md-3 col-lg-3">
                <roll-list msg="近30天TopN设备" :hasBorder="true"/>
                <roll-list msg="近30天 热销SKU" :hasBorder="true" />
                <roll-list msg="Top毛利*周转SKU" :hasBorder="true" />
                <chart-pie msg="设备量" :hasBorder="true" />
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
    import ElementUI from 'element-ui';
    import 'element-ui/lib/theme-chalk/index.css';
    
    import ChartDoubleLine from './components/chart/DoubleLine.vue'
    import ChartPie2 from './components/chart/Pie2.vue'
    import ChartPie1 from './components/chart/Pie1.vue'
    import ChartPie from './components/biz/BizPie.vue'
    import ChartColumn from './components/chart/Column.vue'
    import BizRatioBar from './components/biz/BizRatioBar.vue'
    import Timer from './components/chart/Timer.vue'
    import Topic from './components/Topic.vue'
    import RollList from './components/RollList.vue'
    import SaleSummary from './components/biz/Summary.vue'
    import TwoColums from './components/TwoColums.vue'
    import ThreeColums from './components/ThreeColums.vue'
    import Refund from './components/biz/Refund.vue'
    
    export default {
        name: 'app',
        components: {
            "chart-double-line" : ChartDoubleLine,
            ChartPie,
            ChartPie1,
            ChartPie2,
            ChartColumn,
            Timer,
            Topic,
            RollList,
            TwoColums,
            ThreeColums,
            BizRatioBar,
            SaleSummary,
            Refund,
           
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
