<template>
    <!-- 饼图：百分比 -->
    <div class="pie text-left">
        <sub-topic :msg="msg" :hasBorder="hasBorder" />
        <div v-bind:id="ranId" class=""></div>
    </div>
</template>

<script>
    import DataSet from '@antv/data-set';
import insertCss from 'insert-css';

// 我们用 insert-css 演示引入自定义样式
// 推荐将样式添加到自己的样式文件中
// 若拷贝官方代码，别忘了 npm install insert-css
insertCss(`
        .g2-guide-html {
            width: 100px;
            height: 80px;
            vertical-align: middle;
            text-align: center;
            line-height: 0.4
        }
    
        .g2-guide-html .title {
            font-size: 12px;
            color: #8c8c8c;
            font-weight: 300;
        }
    
        .g2-guide-html .value {
            font-size: 30px;
            color: #000;
            font-weight: bold;
        }
    `);

export default {
    name: 'ChartPie2',
    props: {
        //接收父组件传递过来的参数
        msg: String,
        //chartId: String
        hasBorder: Boolean
    },
    data() {
        // 定义变量
        return {
            chart: null,
            ranId: 'pie2-' + this.random_id(),
            result: { data: [] },
            r: null,
            OFFSET: null,
            center: null,
            canvasHeight: null,
            LINEHEIGHT: null,
            APPEND_OFFSET: null,
            labelGroup: null,
            labels: null,
            canvasWidth: null,
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        },
        
        getEndPoint: function(center, angle, r) {
            return {
                x: center.x + r * Math.cos(angle),
                y: center.y + r * Math.sin(angle)
            };
        },
        drawLabel: function(label) {
            const _anchor = label._anchor,
                _router = label._router,
                fill = label.fill,
                y = label.y;

            const labelAttrs = {
                y,
                fontSize: 12, // 字体大小
                fill: '#808080',
                text: label._data.type + '\n' + label._data.value,
                textBaseline: 'bottom'
            };
            const lastPoint = {
                y
            };

            if (label._side === 'left') {
                // 具体文本的位置
                lastPoint.x = this.APPEND_OFFSET;
                labelAttrs.x = this.APPEND_OFFSET; // 左侧文本左对齐并贴着画布最左侧边缘
                labelAttrs.textAlign = 'left';
            } else {
                lastPoint.x = this.canvasWidth - this.APPEND_OFFSET;
                labelAttrs.x = this.canvasWidth - this.APPEND_OFFSET; // 右侧文本右对齐并贴着画布最右侧边缘
                labelAttrs.textAlign = 'right';
            }

            // 绘制文本
            const text = this.labelGroup.addShape('Text', {
                attrs: labelAttrs
            });
            this.labels.push(text);
            // 绘制连接线
            let points = void 0;
            if (_router.y !== y) {
                // 文本位置做过调整
                points = [[_anchor.x, _anchor.y], [_router.x, y], [lastPoint.x, lastPoint.y]];
            } else {
                points = [[_anchor.x, _anchor.y], [_router.x, _router.y], [lastPoint.x, lastPoint.y]];
            }

            this.labelGroup.addShape('polyline', {
                attrs: {
                    points,
                    lineWidth: 1,
                    stroke: fill
                }
            });
        },
        antiCollision: function(half, isRight) {
            var self = this;
            const startY = this.center.y - this.r - this.OFFSET - this.LINEHEIGHT;
            let overlapping = true;
            let totalH = this.canvasHeight;
            let i = void 0;

            let maxY = 0;
            let minY = Number.MIN_VALUE;
            const boxes = half.map(function(label) {
                const labelY = label.y;
                if (labelY > maxY) {
                    maxY = labelY;
                }
                if (labelY < minY) {
                    minY = labelY;
                }
                return {
                    size: self.LINEHEIGHT,
                    targets: [labelY - startY]
                };
            });
            if (maxY - startY > totalH) {
                totalH = maxY - startY;
            }

            while (overlapping) {
                // eslint-disable-next-line no-loop-func
                boxes.forEach(box => {
                    const target = (Math.min.apply(minY, box.targets) + Math.max.apply(minY, box.targets)) / 2;
                    box.pos = Math.min(Math.max(minY, target - box.size / 2), totalH - box.size);
                });

                // detect overlapping and join boxes
                overlapping = false;
                i = boxes.length;
                while (i--) {
                    if (i > 0) {
                        const previousBox = boxes[i - 1];
                        const box = boxes[i];
                        if (previousBox.pos + previousBox.size > box.pos) {
                            // overlapping
                            previousBox.size += box.size;
                            previousBox.targets = previousBox.targets.concat(box.targets);

                            // overflow, shift up
                            if (previousBox.pos + previousBox.size > totalH) {
                                previousBox.pos = totalH - previousBox.size;
                            }
                            boxes.splice(i, 1); // removing box
                            overlapping = true;
                        }
                    }
                }
            }

            // step 4: normalize y and adjust x
            i = 0;
            boxes.forEach(function(b) {
                let posInCompositeBox = startY; // middle of the label
                b.targets.forEach(function() {
                    half[i].y = b.pos + posInCompositeBox + self.LINEHEIGHT / 2;
                    posInCompositeBox += self.LINEHEIGHT;
                    i++;
                });
            });

            // (x - cx)^2 + (y - cy)^2 = totalR^2
            half.forEach(function(label) {
                const rPow2 = label.r * label.r;
                const dyPow2 = Math.pow(Math.abs(label.y - self.center.y), 2);
                if (rPow2 < dyPow2) {
                    label.x = self.center.x;
                } else {
                    const dx = Math.sqrt(rPow2 - dyPow2);
                    if (!isRight) {
                        // left
                        label.x = self.center.x - dx;
                    } else {
                        // right
                        label.x = self.center.x + dx;
                    }
                }
                self.drawLabel(label);
            });
        },
        addPieLabel: function(_dv, _startAngle, _center) {
            var self = this;
            const halves = [[], []];
            
            const data = _dv.rows;
            let angle = _startAngle;
        
            for (let i = 0; i < data.length; i++) {
                const percent = data[i].percent;
                const targetAngle = angle + Math.PI * 2 * percent;
                const middleAngle = angle + (targetAngle - angle) / 2;
                angle = targetAngle;
                const edgePoint = this.getEndPoint(_center, middleAngle, this.r);
                const routerPoint = this.getEndPoint(_center, middleAngle, this.r + this.OFFSET);
                // label
                const label = {
                    _anchor: edgePoint,
                    _router: routerPoint,
                    _data: data[i],
                    x: routerPoint.x,
                    y: routerPoint.y,
                    r: this.r + this.OFFSET,
                    fill: '#bfbfbf'
                };
                // 判断文本的方向
                if (edgePoint.x < _center.x) {
                    label._side = 'left';
                    halves[0].push(label);
                } else {
                    label._side = 'right';
                    halves[1].push(label);
                }
            } // end of for
        
            const maxCountForOneSide = parseInt(this.canvasHeight / this.LINEHEIGHT, 10);
            halves.forEach(function(half, index) {
                // step 2: reduce labels
                if (half.length > maxCountForOneSide) {
                    half.sort(function(a, b) {
                        return b._percent - a._percent;
                    });
                    half.splice(maxCountForOneSide, half.length - maxCountForOneSide);
                }
        
                // step 3: distribute position (x and y)
                half.sort(function(a, b) {
                    return a.y - b.y;
                });
                self.antiCollision(half, index);
            });
        },
        initCharts: function(_result) {
            var self = this;
            let _data = _result.data;
            let _color = _result.color;

            let startAngle = -Math.PI / 2 - Math.PI / 4;

            const ds = new DataSet();
            const dv = ds.createView().source(_data);
            dv.transform({
                type: 'percent',
                field: 'value',
                dimension: 'type',
                as: 'percent'
            });
            this.chart = new G2.Chart({
                container: this.ranId,
                forceFit: true,
                height: _result.height,
                padding: 'auto'
            });
            this.chart.source(dv);
            this.chart.legend(false);
            this.chart.coord('theta', {
                radius: 0.75,
                innerRadius: 0.5,
                startAngle,
                endAngle: startAngle + Math.PI * 2
            });
            this.chart
                .intervalStack()
                .position('value')
                .color('type', _color)
                .opacity(1)
                .label('percent', {
                    offset: -20,
                    textStyle: {
                        fill: 'white',
                        fontSize: 12,
                        shadowBlur: 2,
                        shadowColor: 'rgba(0, 0, 0, .45)'
                    },
                    formatter: val => {
                        return parseInt(val * 100) + '%';
                    }
                });
            this.chart.guide().html({
                position: ['50%', '50%'],
                html: '<div class="g2-guide-html"><p class="title">总计</p><p class="value">'+this.result.total+'</p></div>'
            });
            this.chart.render();
            // draw label
            this.OFFSET = 20;
            this.APPEND_OFFSET = 50;
            this.LINEHEIGHT = 60;
            const coord = this.chart.get('coord'); // 获取坐标系对象
            this.center = coord.center; // 极坐标圆心坐标
            this.r = coord.radius; // 极坐标半径
            const canvas = this.chart.get('canvas');
            this.canvasWidth = this.chart.get('width');
            this.canvasHeight = this.chart.get('height');
            this.labelGroup = canvas.addGroup();
            this.labels = [];
            this.addPieLabel(dv, startAngle, this.center);
            canvas.draw();
            this.chart.on('afterpaint', function() {
                self.addPieLabel(dv, startAngle, self.center);
            });
        },
        getResult: function() {
            this.result = {
                data: [
                    { type: '朝阳', value: 140 },
                    { type: '海淀', value: 875 },
                    { type: '东城', value: 267 },
                    { type: '西城', value: 853 },
                    { type: '丰台', value: 685 },
                    { type: '崇文', value: 179 },
                    { type: '宣武', value: 88 },
                    { type: '门头沟', value: 583 }
                ]
            };
            this.result.total = 9671;
            this.result.color = ['#0a4291', '#0a57b6', '#1373db', '#2295ff', '#48adff', '#6fc3ff', '#96d7ff', '#bde8ff'];
            this.result.height = 350; //this.$el.offsetHeight; // 获取组件高度
        }
    },
    created() {
        // 生命周期 - 创建完成（可以访问当前this实例）
    },
    mounted() {
        // 生命周期 - 载入后, Vue 实例挂载到实际的 DOM 操作完成
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
    watch: {
        // 处理数据变化时的刷新动作
    }
};
</script>

<style></style>
