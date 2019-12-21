<template>
    <div style="margin-top: -35px;">
    <!-- TODO: 双Cycle -->
    <div v-bind:id="ranId" v-bind:chartData="chartData"></div>
    </div>
</template>

<script>
    import insertCss from 'insert-css';
    
    // 可将样式添加到自己的样式文件中
    insertCss(`
        .g2-guide-html {
            width: 50px;
            height: 50px;
            vertical-align: middle;
            text-align: center;
            line-height: 0.1
        }
    
        .g2-guide-html .title {
            font-size: 12px;
            color: #8c8c8c;
            font-weight: 300;
        }
    
        .g2-guide-html .value {
            font-size: 18px;
            color: #000;
            font-weight: bold;
        }
    `);
    
    export default {
        name: 'TwinCycle',
        props: {
            //接收父组件传递过来的参数
            chartData: undefined,
        },
        data() {
            // 定义变量
            return {
                ranId: 'twincycle-' + this.random_id(),
                chart : null,
                data : [
                  { type: '男性', value: 56.4 },
                  { type: '快销品', value: 43.6 }
                ],
            }
        },
        methods:{
            random_id: function() {
                var tmpDate = new Date();
                var tmp = tmpDate.getTime();
                return tmp;
            },
        },
        mounted() {
            this.chart = new G2.Chart({
              container: this.ranId,
              forceFit: true,
              height: 200,
              padding: 'auto'
            });
            this.chart.source(this.data);
            this.chart.legend(false);
            this.chart.facet('rect', {
              fields: [ 'type' ],
              padding: 20,
              showTitle: false,
              eachView: function eachView(view, facet) {
                const data1 = facet.data;
                let color;
                if (data1[0].type === '男性') {
                  color = '#0a9afe';
                } else {
                  color = '#f0657d';
                }
                data1.push({ type: '其他', value: 100 - data1[0].value });
                view.source(data1);
                view.coord('theta', {
                  radius: 1.0,
                  innerRadius: 0.5
                });
                view.intervalStack()
                  .position('value')
                  .color('type', [ color, '#eceef1' ])
                  .opacity(1);
                view.guide().html({ 
                  position: [ '50%', '55%' ], // 内部html相对于外部的位置比例
                  html: `
                    <div class="g2-guide-html">
                      <p class="title">${data1[0].type}</p>
                      <p class="value">${data1[0].value}%</p>
                    </div>
                  `
                });
              }
            });
            this.chart.render();
        }
    }
    
    
    
</script>

<style>
</style>
