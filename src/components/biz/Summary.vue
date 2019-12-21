<template>
    <!-- 大屏Summary数据 -->
    <div class="summary" style="margin-bottom: 5px;">
        <div class="row">
            <div class="col-md-6 col-lg-6" style="padding: 5px;">
                <div class="row tt">
                    <div class="col-md-6 col-lg-6 text-left" style="font-size: 14px;">
                        总成交金额
                    </div>
                    <div class="col-md-6 col-lg-6 text-right" style="font-size: 14px;">
                        {{result.total.cnt}}笔
                    </div>
                </div>
                <div class="row tt">
                    <div class="col-md-12 col-lg-12 text-center">
                        <h3>¥ {{result.total.amt}}</h3>
                    </div>
                </div>
                <div class="row tt"><div class="col-md-12 col-lg-12 text-center" style="font-size: 14px;">
                    总毛利： ￥{{result.total.grossAmt}} | {{result.total.rate}}%
                </div>
                </div>
            </div>
            <div class="col-md-6 col-lg-6" style="padding: 5px;">
                <div class="row tt">
                    <div class="col-md-6 col-lg-6 text-left" style="font-size: 14px;">
                        今日成交额
                    </div>
                    <div class="col-md-6 col-lg-6 text-right" style="font-size: 14px;">
                        {{result.today.cnt}}笔
                    </div>
                </div>
                <div class="row tt">
                    <div class="col-md-12 col-lg-12 text-center">
                        <h3>¥ {{result.today.amt}}</h3>
                    </div>
                </div>
                <div class="row tt">
                    <div class="col-md-12 col-lg-12 text-center" style="font-size: 14px;">
                        毛利： ￥{{result.today.grossAmt}} | {{result.today.rate}}%
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Summary',
        data() {
            // 定义变量
            return {
                result: {
                    total: {"amt": 14567.92, "cnt": 1111, "grossAmt": 123456.91, "rate": 75.96},
                    today: {"amt": 8567.92, "cnt": 555, "grossAmt": 4456.91, "rate": 55.96},
                },
            }
        },
        mounted() {
            // TODO 远程获取rest数据
        }
    }
</script>

<style scoped>
    div.tt {
        /*background-color: #F0F0F0;*/
    }
</style>
