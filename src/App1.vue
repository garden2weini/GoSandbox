<template>
    <div id="app1" class="container">
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center">
                <chart-pie msg="设备量"/>ok?
            </div>
            <div class="col-md-6 col-lg-6 text-center">
                <chart-double-line msg="活跃设备日均销售额"/>
            </div>
        </div>
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center">
                <chart-pie msg="近30天退款金额" />
            </div>
            <div class="col-md-6 col-lg-6 text-center">
                <spectaculars msg="近30天用户复购"/>
            </div>
        </div>
    </div>
</template>

<script>
    //import HelloWorld from './components/HelloWorld.vue'
    import ChartDoubleLine from './components/chart/DoubleLine.vue'
    import ChartPie from './components/chart/Pie.vue'
    import Spectaculars from './components/chart/Spectaculars.vue'

    export default {
        name: 'app',
        components: {
            "chart-double-line" : ChartDoubleLine,
            "chart-pie" : ChartPie,
            "Spectaculars" : Spectaculars
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
