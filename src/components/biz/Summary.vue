<template>
    <!-- 大屏Summary数据 -->
    <div class="summary" style="margin-bottom: 5px;">
        <div class="row" style="margin-left: 0px;margin-right: 0px;">
            <div class="m-block left col-md-6 col-lg-6">
                <div class="row">
                    <div class="col-md-6 col-lg-6 text-left">总成交金额</div>
                    <div class="col-md-6 col-lg-6 text-right">{{ result.total.cnt }}笔</div>
                </div>
                <div class="row">
                    <div class="col-md-12 col-lg-12 text-center">
                        <h3>¥{{ result.total.amt }}</h3>
                    </div>
                </div>
                <div class="row">
                    <div class="col-md-12 col-lg-12 text-center">
                        总毛利： ￥{{ result.total.grossAmt }} | {{ result.total.rate }}%
                        </div>
                </div>
            </div>
            <div class="m-block right col-md-6 col-lg-6">
                <div class="row">
                    <div class="col-md-6 col-lg-6 text-left">今日成交额</div>
                    <div class="col-md-6 col-lg-6 text-right">{{ result.today.cnt }}笔</div>
                </div>
                <div class="row">
                    <div class="col-md-12 col-lg-12 text-center" style="margin-top: 7px;margin-bottom: 7px;">
                        <font class="amt">¥{{ result.today.amt }}&nbsp; &nbsp;</font>
                        <font class="text-success" v-if="result.today.upRate >= 0">
                            {{result.today.upRate}}%
                        </font>
                        <font class="text-danger" v-else>
                            {{result.today.upRate}}%
                        </font>
                        
                    </div>
                </div>
                <div class="row">
                    <div class="col-md-12 col-lg-12 text-center">
                        毛利： ￥{{ result.today.grossAmt }} | {{ result.today.rate }}%
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import $ from 'jquery';
export default {
    name: 'Summary',
    data() {
        // 定义变量
        return {
            result: {
                total: { amt: 0, cnt: 0, grossAmt: 0, rate: 0 },
                today: { amt: 0, cnt: 0, grossAmt: 0, rate: 0 , upRate: 0}
            }
        };
    },
    methods: {},
    mounted() {
        var self = this;
        function updateSummary() {
            // NOTE 远程获取rest数据
            $.getJSON(self.DATA_BASE_URL + 'summary.json', sourceData => {
                self.result = sourceData;
            });
        }
        updateSummary();
        setInterval(updateSummary, this.REFRESH_INTERVAL);
    },
    watch: {}
};
</script>

<style scoped>
div {
    font-size: 14px;
}
div.left {
    padding-top: 2px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
}
div.right {
    padding-top: 2px;
    padding-bottom: 5px;
    padding-left: 10px;
    padding-right: 10px;
}
font.amt {
    font-size: 18px;
    font:bold;
}
</style>
