<template>
    <div class="m-block pie text-left">
        <!-- 饼图+图示 -->
        <sub-topic :msg="msg" :hasBorder="hasBorder"/>
        <chart-pie :vData="data"/>
    </div>
</template>

<script>
import ChartPie from '../chart/Pie.vue';

export default {
    name: 'BizPie',
    components: {
        ChartPie,
    },
    props: {
        //接收父组件传递过来的参数
        msg: String,
        hasBorder: Boolean,
        data: {},
    },
    data() {
        // 定义变量
        return {
            //chart: null,
        };
    },
    mounted() {
        //console.log("BizPie:");
        //console.log(new Date());
        //console.log(this.data);
    }
};
</script>

<style>

</style>
