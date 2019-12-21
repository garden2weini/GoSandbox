<template>
    <div class="pie text-left">
        <!-- 饼图+图示 -->
        <sub-topic :msg="msg" :hasBorder="hasBorder"/>
        <chart-pie/>
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
    },
    data() {
        // 定义变量
        return {
            chart: null,
            
        };
    },
};
</script>

<style>

</style>
