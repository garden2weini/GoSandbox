<template>
    <div class="totel-number">
        <div class="panel-top">
            <div class="panel-top_Left"></div>
            <div class="panel-top_Text flex-center">数量</div>
            <div class="panel-top_Right"></div>
        </div>
        <div class="panel-body">
            <div class="counter-wrap">
                <ul class="counter-board">
                    <li class="num-board" v-for="item in totalNumber">
                        <span class="num u">
                            <i class="w">{{ item }}</i>
                        </span>
                        <span class="num b">
                            <i class="w">-{{ item }}</i>
                        </span>
                    </li>
                </ul>
            </div>
        </div>
        
    </div>
</template>

<!--
两种情况不会触发视图更新，需要换种变通写法：
- 当你利用索引直接设置一个项时，例如： vm.items[indexOfItem] = newValue
- 当你修改数组的长度时，例如： vm.items.length = newLength
-->
<script>
    import $ from 'jquery';
    import * as d3 from 'd3';
    export default {
        name: "Timer",
        data() {
            return {
                totalNumber: new Array(3)
            };
        },
        methods: {
            
        },
        beforeMount() {
            var self = this;
            function getTotalNumber() {
                let newArray = new Array(3);
                for (let i = 0; i < self.totalNumber.length; i++) {
                    newArray[i] = Math.ceil(Math.random() * 10) - 1;
                    //newArray[i] = 123;
                }
                self.totalNumber = newArray;
            }
            //设置定时器，每3秒刷新一次
            getTotalNumber();
            setInterval(getTotalNumber, 1000);
        }
    };
</script>
