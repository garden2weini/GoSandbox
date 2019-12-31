<template>
    <div class="m-block rolllist text-left" style="margin-bottom: 5px;">
        <!-- 饼图+图示 -->
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <roll-list-inner :vData="data" />
    </div>
</template>

<script>
import RollListInner from './RollListInner.vue';
export default {
    name: 'RollList',
    components: {
        RollListInner
    },
    props: {
        //接收父组件传递过来的参数
        msg: String,
        hasBorder: Boolean,
        data: {},
    }
};
</script>

<style></style>
