<template>
    <!-- 饼图：百分比 -->
    <div class="pie text-left">
        <sub-topic :msg="msg" :hasBorder="hasBorder"/>
        <div v-bind:id="ranId" class=""></div>
    </div>
</template>

<script>
import DataSet from '@antv/data-set';
import insertCss from 'insert-css';

insertCss(`
  .g2-label-item-inner {
    text-align: center;
    font-size: 12px;
    color: #ffffff;
    text-shadow: 0px 0px 2px #595959;
  }

  .g2-label-item-outer {
    width:60px;
    font-size: 12px;
    color: #595959;
  }
`);

export default {
    name: 'ChartPie1',
    props: {
        //接收父组件传递过来的参数
        msg: String,
        //chartId: String
        hasBorder: Boolean,
    },
    data() {
        // 定义变量
        return {
            chart: null,
            ranId: 'pie-' + this.random_id(),
            result: { data: [] }
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        },
        initCharts: function(_result) {
            let _data = _result.data;
            let _color = _result.color;
            const ds = new DataSet();
            const dv = ds.createView().source(_data);
            dv.transform({
                type: 'percent',
                field: 'y',
                dimension: 'x',
                as: 'percent'
            });

            this.chart = new G2.Chart({
                container: this.ranId,
                forceFit: true,
                height: _result.height,
                padding: 'auto' // 为了防止小图时图表变形
            });
            // 模拟window.resize时才会触发forceFit: true，以便自适应屏幕的宽度
            const e = document.createEvent('Event');
            e.initEvent('resize', true, true);
            window.dispatchEvent(e);
            
            this.chart.source(dv);
            this.chart.legend(false);
            this.chart.coord('theta', {
                radius: 0.75
            });
            this.chart
                .intervalStack()
                .position('y')
                .color('x', _color)
                .opacity(1)
                .label('y', function(val) {
                    const offset = val > 0.02 ? -30 : 30;
                    const label_class = val > 0.02 ? 'g2-label-item-inner' : 'g2-label-item-outer';
                    return {
                        offset,
                        useHtml: true,
                        htmlTemplate: (text, item) => {
                            const d = item.point;
                            const percent = String(parseInt(d.percent * 100)) + '%';
                            return '<div class=' + label_class + '>' + d.x + percent + '</div>';
                        }
                    };
                });
            this.chart.render();
        },
        getResult: function() {
            this.result = {
                data: [
                    { x: '朝阳', y: 0.4 },
                    { x: '海淀', y: 0.11 },
                    { x: '东城', y: 0.1 },
                    { x: '西城', y: 0.01 },
                    { x: '丰台', y: 0.07 },
                    { x: '崇文', y: 0.13 },
                    { x: '宣武', y: 0.08 },
                    { x: '门头沟', y: 0.1 }
                ]
            };
            this.result.color = ['#2593fc', '#38c060', '#27c1c1', '#705dc8', '#3b4771', '#f9cb34', '#ab5771', '#b7ab34'];
            this.result.height = 200; //this.$el.offsetHeight; // 获取组件高度
        }
    },
    created() { // 生命周期 - 创建完成（可以访问当前this实例）
        
    },
    mounted() { // 生命周期 - 载入后, Vue 实例挂载到实际的 DOM 操作完成
        // 准备定时更新逻辑
        var self = this;
        // 更新Chart数据的逻辑
        function reloadData() {
            self.getResult();
            self.chart.source(self.result.data);  
        }
        self.getResult();
        self.initCharts(self.result);
        //设置定时器，每60秒刷新一次
        setInterval(reloadData, 60 * 1000);
    
    },
    watch: { // 处理数据变化时的刷新动作
        
    }
};
</script>

<style></style>
