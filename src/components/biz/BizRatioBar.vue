<template>
    <!-- 种类占比 -->
    <div class="users text-left" style="margin-bottom: 5px;">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <ratio-bar exStyle="width: 300px;height:100px;" :vData="data"/>
    </div>
</template>

<script>
import RatioBar from '../chart/RatioBar.vue'

export default {
    name: 'BizRatioBar',
    components:{
        RatioBar,
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
            chartData: {
                types: ['支付宝', '微信', '京东支付'],
                colors: ['#E8A61F', '#E67C26', '#0CD8A7'],
                values: [53.1, 23 ,1]
            },
        }
    }
};
</script>

<style></style>
