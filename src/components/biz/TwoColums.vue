<template>
    <div class="m-block users text-left" style="margin-bottom: 5px;">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div class="row">
            <div class="col-md-6 col-lg-6 text-center" >
                {{result.users.today}} {{result.users.rate}}%
            </div>
            <div class="col-md-6 col-lg-6 text-center" >
                {{result.users.total}}
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
    import $ from 'jquery'
    export default {
        name: 'TwoColumns',
        props: {
            //接收父组件传递过来的参数
            msg: String,
            hasBorder: Boolean,
        },
        data() {
            // 定义变量
            return {
                result: {
                    "users": {
                        "today": 0,
                        "rate": 0,
                        "total": 0
                    },
                    "orders": {
                        "today": 0,
                        "last7": 0,
                        "last30": 0
                    }
                },
            }
        },
        mounted() {
            var self = this;
            function updateSummary() {
                // NOTE 远程获取rest数据
                $.getJSON(self.DATA_BASE_URL + 'others.json', (sourceData) => {
                      self.result = sourceData;
                    });
            }
            updateSummary();
            setInterval(updateSummary, this.REFRESH_INTERVAL);
        },
    }
</script>

<style scoped>
    div {
        font-size: 14px;
    }
</style>
