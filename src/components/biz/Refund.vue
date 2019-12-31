<template>
    <div class="m-block refund text-left">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center" >
                退款占比：{{data1.refundRate}}%
            </div>
            <div class="col-md-3 col-lg-3 text-left" >
                退款：{{data1.refundCnt}}
            </div>
            <div class="col-md-3 col-lg-3 text-left" >
                ¥{{data1.refundAmt}}
            </div>
        </div>
        <div class="row">
            <ratio-bar exStyle="width: 350px;height:100px;" :vData="data"/>
        </div>
    </div>
    
</template>

<script>
    import RatioBar from '../chart/RatioBar.vue'
    export default {
        name: 'Refund',
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
				data1: {
					refundRate:0,
					refundCnt: 0,
					refundAmt: 0},
            }
        },
		beforeUpdate() { // 当data更新时触发
		    this.data1 = this.data;
		}
    }
</script>

<style scoped>
    div {
        padding: 1px 0px 0px 0px; /* 上内边距;右内边距;下内边距;左内边距 */
        font-size: 13px;
        
    }
    div.row {
        margin-left: 0px;
        margin-right: 0px;
    }
</style>
