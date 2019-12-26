<template>
    <div class="refund text-left" style="margin-bottom: 5px;">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center" >
                退款占比：{{data.refundRate}}%
            </div>
            <div class="col-md-3 col-lg-3 text-left" >
                退款：{{data.refundCnt}}
            </div>
            <div class="col-md-3 col-lg-3 text-left" >
                ¥{{data.refundAmt}}
            </div>
        </div>
        <div class="row">
            <ratio-bar exStyle="width: 320px;height:100px;" :vData="data"/>
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
                chartData: {
                    types: ['掉错商品', '没有掉货', '商品损坏', '订单超时'],
                    colors: ['#E8A61F', '#E67C26', '#0CD8A7', '#ACD8A7'],
                    values: [53.1, 40, 23, 10]
                },
            }
        }
    }
</script>

<style scoped>
    div {
        padding: 0px 1px 1px 0px; /* 上内边距;右内边距;下内边距;左内边距 */
        font-size: 12px;
    }
</style>
