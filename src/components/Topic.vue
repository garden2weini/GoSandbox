<template>
    <!-- 每个图例的标题 -->
    <div>
        <el-button style="font-size: 20px;border-radius: 4px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)">
            {{ msg }}
        </el-button>
    </div>
</template>

<script>
    import ElementUI from 'element-ui';
    import 'element-ui/lib/theme-chalk/index.css';
    
    export default {
        name: 'Topic',
        props: {
            //接收父组件传递过来的参数
            msg: String,
        },
    }
</script>

<style>
</style>

