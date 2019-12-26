<template>
    <!--种类占比, ref: https://gallery.echartsjs.com/editor.html?c=xWv8FheY8-->
    <!-- 为 ECharts 准备一个具备大小（宽高）的 DOM -->
    <div v-bind:id="ranId" :style="exStyle" v-bind:vData="vData"></div>
</template>

<script>
import echarts from 'echarts';

export default {
    name: 'RawRatioBar',
    props: {
        //接收父组件传递过来的参数
        exStyle: String, // NOTE: 用来配置组件的高度和宽度
        vData: undefined,
    },
    data() {
        // 定义变量
        return {
            option: null,
            chart: null,
            ranId: 'ratiobar-' + this.random_id(),
            result: { data: [] },
            rawOption: { //指定图表的配置项和数据模版
                //backgroundColor: '#000E1B',
                backgroundColor: 'white',
                legend: {
                    icon: 'circle',
                    bottom: '20%', // label的下相对位置
                    left: '10%', // label的左相对位置
                    itemWidth: 7,
                    itemHeight: 7,
                    itemGap: 5, // label的间隔
                    textStyle: {
                        //color: '#89A7AF'
                        color: 'black'
                    },
                    data: [] // NOTE 待添加的类型数据
                },
                xAxis: [
                    {
                        type: 'value',
                        axisTick: {
                            show: false
                        },
                        axisLine: {
                            show: false
                        },
                        axisLabel: {
                            show: false
                        },
                        splitLine: {
                            show: false
                        }
                    }
                ],
                yAxis: [
                    {
                        //type: 'category',
                        data: [''],
                        axisTick: {
                            show: false
                        },
                        axisLine: {
                            show: false
                        },
                        axisLabel: {
                            textStyle: {
                                color: '#fff'
                            }
                        }
                    }
                ],
                series: [] // NOTE 待补充
            },
            item1_tmp: {
                name: '油车',
                type: 'bar',
                barWidth: 16,
                stack: '种类占比',
                label: {
                    normal: {
                        borderWidth: 10,
                        distance: 20,
                        align: 'center',
                        verticalAlign: 'middle',
                        borderRadius: 1,
                        borderColor: '#E8A61F',
                        backgroundColor: '#E8A61F',
                        show: true,
                        position: 'top',
                        formatter: '{c}%',
                        color: '#000'
                    }
                },
                itemStyle: {
                    color: '#E8A61F'
                },
                data: [
                    {
                        value: 53.1,
                        itemStyle: {
                            normal: {
                                color: {
                                    type: 'bar',
                                    colorStops: [
                                        {
                                            offset: 0,
                                            color: '#E8A61F' // 0% 处的颜色
                                        },
                                        {
                                            offset: 1,
                                            color: '#E8A61F' // 100% 处的颜色
                                        }
                                    ],
                                    globalCoord: false // 缺省为 false
                                }
                            }
                        }
                    }
                ]
            },
            item2_tmp: {
                name: '1nd-triangle',
                type: 'line',
                barWidth: 0,
                markPoint: {
                    symbol: 'triangle',
                    symbolRotate: '180',
                    itemStyle: {
                        color: {
                            type: 'linear',
                            x: 0,
                            y: 0,
                            x2: 1,
                            y2: 0,
                            colorStops: [
                                {
                                    offset: 0,
                                    color: '#E8A61F' // 0% 处的颜色
                                },
                                {
                                    offset: 1,
                                    color: '#E8A61F' // 100% 处的颜色
                                }
                            ],
                            globalCoord: false // 缺省为 false
                        }
                    },
                    symbolSize: [6, 5], // 容器大小
                    symbolOffset: [0, -15], //位置偏移
                    data: [
                        {
                            coord: [53.11 / 2]
                        }
                    ],
                    label: {
                        normal: {
                            show: false
                        },
                        offset: [0, 0]
                    }
                }
            }
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        },
        init_chart_option: function(types, colors, values) {
            var self = this;
            this.option = JSON.parse(JSON.stringify(self.rawOption));
            var tmp_value = 0;
            for (let i = 0; i < types.length; i++) {
                // 数组追加一个元素
                this.option.legend.data[this.option.legend.data.length] = { name: types[i] };

                // 复制json
                var item1 = JSON.parse(JSON.stringify(this.item1_tmp));
                item1.name = types[i];
                item1.itemStyle.color = colors[i];
                item1.label.normal.borderColor = colors[i];
                item1.label.normal.backgroundColor = colors[i];
                item1.data[0].value = values[i]; //53.1;
                item1.data[0].itemStyle.normal.color.colorStops[0].color = colors[i];
                item1.data[0].itemStyle.normal.color.colorStops[1].color = colors[i];

                this.option.series[this.option.series.length] = item1;

                var item2 = JSON.parse(JSON.stringify(this.item2_tmp));
                item2.name = i + 1 + 'nd-triangle';
                item2.markPoint.itemStyle.color.colorStops[0].color = colors[i];
                item2.markPoint.itemStyle.color.colorStops[1].color = colors[i];
                item2.markPoint.data[0].coord = [tmp_value + values[i] / 2];
                tmp_value = tmp_value + values[i];
                this.option.series[this.option.series.length] = item2;
            }
        }
    },
    mounted() {
        // 基于准备好的dom，初始化echarts实例
        this.chart = echarts.init(document.getElementById(this.ranId));

    },
    beforeUpdate() { // 当data更新时触发
        this.init_chart_option(
            this.vData.data.types, 
            this.vData.data.colors, 
            this.vData.data.values);
        
        // 使用刚指定的配置项和数据显示图表。
        this.chart.setOption(this.option, true);
        
        
    },
};
</script>

<style></style>
