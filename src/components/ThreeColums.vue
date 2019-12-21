<template>
    <div class="users text-left" style="margin-bottom: 5px;">
        
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div class="row">
            <div class="col-md-3 col-lg-3 text-center" >
                $8.27
            </div>
            <div class="col-md-3 col-lg-3 text-center" >
                $12.38
            </div>
            <div class="col-md-3 col-lg-3 text-center" >
                $13.84
            </div>
        </div>
        <div class="row">
            <div class="col-md-3 col-lg-3 text-center" >
                今天
            </div>
            <div class="col-md-3 col-lg-3 text-center" >
                 近7天
            </div>
            <div class="col-md-3 col-lg-3 text-center" >
                近30天
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'ThreeColumns',
        props: {
            //接收父组件传递过来的参数
            msg: String,
            hasBorder: Boolean,
        },
    }
</script>

<style scoped>
    div {
        padding: 0px 1px 1px 0px; /* 上内边距;右内边距;下内边距;左内边距 */
        font-size: 14px;
    }
</style>
