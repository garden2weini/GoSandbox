<template>
    <!-- 种类占比 -->
    <div class="m-block users text-left" style="margin-bottom: 5px;">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <ratio-bar exStyle="width: 350px;height:100px;" :vData="data"/>
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
            
        }
    }
};
</script>

<style></style>
