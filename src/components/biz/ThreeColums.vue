<template>
    <div class="m-block users text-left">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div class="row">
            <div class="col-md-4 col-lg-4 text-center" >
                ${{result.orders.today}}
            </div>
            <div class="col-md-4 col-lg-4 text-center" >
                ${{result.orders.last7}}
            </div>
            <div class="col-md-4 col-lg-4 text-center" >
                ${{result.orders.last30}}
            </div>
        </div>
        <div class="row">
            <div class="col-md-4 col-lg-4 text-center" >
                今天
            </div>
            <div class="col-md-4 col-lg-4 text-center" >
                 近7天
            </div>
            <div class="col-md-4 col-lg-4 text-center" >
                近30天
            </div>
        </div>
    </div>
</template>

<script>
    import $ from 'jquery'
    export default {
        name: 'ThreeColumns',
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
        padding: 0px 1px 1px 0px; /* 上内边距;右内边距;下内边距;左内边距 */
        font-size: 14px;
    }
</style>
