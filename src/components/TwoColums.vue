<template>
    <div class="users text-left" style="margin-bottom: 5px;">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center" >
                181 25.36%
            </div>
            <div class="col-md-6 col-lg-6 text-center" >
                34567
            </div>
        </div>
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center" >
                今天
            </div>
            <div class="col-md-6 col-lg-6 text-center" >
                累计用户
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'TwoColumns',
        props: {
            //接收父组件传递过来的参数
            msg: String,
            hasBorder: Boolean,
        },
    }
</script>

<style scoped>
    div {
        font-size: 14px;
    }
</style>
