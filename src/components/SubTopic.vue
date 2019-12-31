<template>
    <!-- 每个图例的标题 -->
    <div>
        <el-button class="border" v-if="hasBorder">
            {{ msg }}
        </el-button>
        <el-button class="no-border" v-else>
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
            hasBorder:Boolean,
        },
        
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    /* 修饰el-button风格 */
    button {
        font-size: 15px;
        padding-top: 5px;
        padding-bottom: 5px;
        padding-left: 10px;
        background-color: #10707F;
        color:#00FF00;
    }
    button.border {
         border-radius: 4px;
         box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    }
    button.no-border {
          border:none;
    }
</style>
