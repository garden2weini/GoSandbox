<template>
    <div  style="margin-top: -15px;">
        <!-- 饼图+图示 -->
        <div v-bind:id="ranId" v-bind:vData="vData"></div>
    </div>
</template>

<script>
export default {
    name: 'ChartPie',
    props: {
        //接收父组件传递过来的参数
        vData: {},
    },
    data() {
        // 定义变量
        return {
            chart: null,
            ranId: 'pie-' + this.random_id(),
            colors: null,
            data: null,
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        },
        initCharts: function(_offsetHeight) {
            this.chart = new G2.Chart({
                container: this.ranId,
                forceFit: true,
                height: _offsetHeight,
                padding: 'auto'
            });
            this.chart.source(this.data);
            this.chart.legend({
                position: 'right-center', // 确定Label放置的位置
                offsetX: -20,
            });
            this.chart.coord('theta', {
                radius: 0.75
            });
            this.chart
                .intervalStack()
                .position('value')
                .color('name', this.colors)
                .style({
                    stroke: 'white',
                    lineWidth: 0//1
                })
                .label('value', val => {
                    if (val < 3) {
                        return null;
                    }
                    return {
                        offset: -18, //  饼图各区域中Label的位置偏移量
                        textStyle: {
                            fill: 'white',
                            fontSize: 12,
                            shadowBlur: 2,
                            shadowColor: 'rgba(0, 0, 0, .45)'
                        },
                        formatter: text => {
                            //return text + '%';
                            return text;
                        }
                    };
                });
            this.chart.render();
        }
    },
    // 生命周期 - 创建完成（可以访问当前this实例）
    created() {
        
    },
    mounted() { // 生命周期 - 载入后, Vue 实例挂载到实际的 DOM 操作完成，一般在该过程进行 Ajax 交互
        var self = this;
        /*
        this.data = [
                {name: '海淀',value: 44.9},
                {name: '西城',value: 19.7},
                {name: '东城',value: 17.3},
                {name: '崇文',value: 14.4},
                {name: '宣武',value: 2.5},
                {name: '顺义',value: 2.5}];
            
        this.colors = ["#1890ff", "#37c661", "#fbce1e", "#2b3b79", "#8a4be2", "#1dc5c5"];
        */
        // NOTE: chart label color有默认值
        let height = 250 // 获取父级高度
        this.initCharts(height);
    },
    beforeUpdate() { // 当data更新时触发
        //console.log("beforeUpdate:");
        //console.log(this.vData);
        this.data = this.vData.data;
        this.colors = this.vData.colors;
        //this.chart
        this.chart.changeData(this.data);
        
        
    },
    updated() {
        //console.log("updated:");
        //console.log(this.vData);
    },
    watch: { // 处理数据变化时的刷新动作
        
    }
};
</script>

<style>

</style>
