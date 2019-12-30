<template>
    <div class="line2 text-left">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div v-bind:id="ranId" v-bind:vData="vData"></div>
    </div>
</template>

<script>
import DataSet from '@antv/data-set';

export default {
    name: 'ChartDoubleLine',
    props: {
        msg: String,
        hasBorder: Boolean,
        vData: {}
    },
    data() {
        // 定义变量
        return {
            chart: null,
            ranId: 'doubleline-' + this.random_id(),
            data: [
                {
                    day: '12-10',
                    '销售金额': 0,
                    '退款金额': 0
                },
                {
                    day: '12-11',
                    '销售金额': 10,
                    '退款金额': 10
                }
            ],
            dataView: null
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        },
        initCharts: function() {
            var ds = new DataSet();
            this.dataView = ds.createView().source(this.data);
            // fold 方式完成了行列转换，如果不想使用 DataSet 直接手工转换数据即可
            this.dataView.transform({
                type: 'fold',
                fields: ['销售金额', '退款金额'], // 展开字段集
                key: 'type', // key字段
                value: 'amt' // value字段
            });
            this.chart = new G2.Chart({
                container: this.ranId,
                forceFit: true,
                height: 200
            });

            this.chart.source(this.dataView);
            
            this.chart.tooltip({
                crosshairs: {
                    type: 'line'
                }
            });
            this.chart.axis('amt', {
                label: {
                    formatter: val => {
                        return '¥' + val;
                    }
                }
            });
            this.chart
                .line()
                .position('day*amt')
                .color('type')
                .size(2)
                .shape('smooth');
            this.chart
                .point()
                .position('day*amt')
                .color('type')
                .size(4)
                .shape('circle')
                .style({
                    stroke: '#fff',
                    lineWidth: 1
                });
            this.chart.render();
        }
    },
    mounted() {
        this.initCharts();
    },
    beforeUpdate() {
        // 当data更新时触发
        this.data = this.vData.data;
        if (this.chart) {
            // 如果存在的话就删除图表再重新生成
            this.chart.destroy();
        }
        this.initCharts();
        
    }
};
</script>

<style></style>
