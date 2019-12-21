<template>
    <!-- 基础柱状图 -->
    <div class="column text-left">
        <sub-topic :msg="msg" :hasBorder="hasBorder"/>
        <div v-bind:id="ranId" style="width: 300px;height:150px;"></div>
    </div>
</template>

<script>

export default {
    name: 'ChartColumn',
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
            ranId: 'column-' + this.random_id(),
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
            
            this.chart = new G2.Chart({
                container: this.ranId,
                forceFit: true,
                height: _result.height,
                padding: [ 20, 20, 20, 20 ]
            });
            
            // 模拟window.resize时才会触发forceFit: true，以便自适应屏幕的宽度
            const e = document.createEvent('Event');
            e.initEvent('resize', true, true);
            window.dispatchEvent(e);
            
            this.chart.source(_data);
            this.chart.scale('value', {
              alias: '用户数'
            });
            this.chart.axis('time', {
              label: {
                textStyle: {
                  fill: '#aaaaaa'
                }
              },
              tickLine: {
                alignWithLabel: false,
                length: 0
              }
            });
            
            this.chart.axis('value', {
              label: {
                textStyle: {
                  fill: '#aaaaaa'
                }
              }
            });
            
            this.chart.tooltip({
              share: true
            });
            this.chart.interval().position('time*value')
              .opacity('time', val => {
                /*
                if (val === '30+') {
                  return 0.4;
                }*/
                return 1;
              })
              .style('time', {
                lineWidth: val => {
                  /*
                  if (val === '30+') {
                    return 1;
                  }*/
                  return 0;
                },
                stroke: '#636363',
                lineDash: [ 3, 2 ]
              });
            
            this.chart.render();
        },
        getResult: function() {
            this.result = {
                data: [ // 
                  { time: '1-5', value: 30 },
                  { time: '6-10', value: 90 },
                  { time: '11-20', value: 50 },
                  { time: '21-30', value: 30 },
                  { time: '30+', value: 70 }
                ]
            };
            this.result.height = 120; //this.$el.offsetHeight; // 获取组件高度
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
