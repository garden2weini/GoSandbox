<template>
    <div v-bind:id="ranId1">
        <ul v-bind:id="ranId2">
            <li>点位地址1...................¥8719/3600笔</li>
            <li>点位地址2...................¥8719/3600笔</li>
            <li>点位地址3...................¥8719/3600笔</li>
            <li>点位地址4...................¥8719/3600笔</li>
            <li>点位地址5...................¥8719/3600笔</li>
            <li>点位地址6...................¥8719/3600笔</li>
            <li>点位地址7...................¥8719/3600笔</li>
            <li>点位地址8...................¥8719/3600笔</li>
        </ul>
        <ul v-bind:id="ranId3"></ul>
    </div>
</template>

<script>
import $ from 'jquery';
export default {
    name: 'RollListInner',
    props: {
        //接收父组件传递过来的参数
        msg: String,
        hasBorder: Boolean
    },
    data() {
        // 定义变量
        return {
            ranId1: 'rolldiv-' + this.random_id(),
            ranId2: 'ul1-' + this.random_id(),
            ranId3: 'ul2-' + this.random_id(),
            result: { data: [] }
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        }
    },
    mounted() {
        var self = this;
        function roll(t) {
            var ul1 = document.getElementById(self.ranId2);
            var ul2 = document.getElementById(self.ranId3);
            var ulbox = document.getElementById(self.ranId1);
            ul2.innerHTML = ul1.innerHTML;
            ulbox.scrollTop = 0; // 开始无滚动时设为0
            // 设置定时器，参数t用在这为间隔时间（单位毫秒），参数t越小，滚动速度越快
            var timer = setInterval(rollStart, t);
            // 鼠标移入div时暂停滚动
            ulbox.onmouseover = function() {
                clearInterval(timer);
            };
            // 鼠标移出div后继续滚动
            ulbox.onmouseout = function() {
                timer = setInterval(rollStart, t);
            };
        }

        // 开始滚动函数
        function rollStart() {
            // 上面声明的DOM对象为局部对象需要再次声明
            var ul1 = document.getElementById(self.ranId2);
            var ul2 = document.getElementById(self.ranId3);
            var ulbox = document.getElementById(self.ranId1);
            // scrollTop为空时进行重置
            if(!ulbox.scrollTop) ulbox.scrollTop=0;
            // 正常滚动不断给scrollTop的值+1,当滚动高度大于列表内容高度时恢复为0
            if (ulbox.scrollTop >= ul1.scrollHeight) {
                ulbox.scrollTop = 0;
            } else {
                ulbox.scrollTop++;
            }
        }
        roll(50);
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
* {
    margin: 0;
    padding: 0;
}
div {
    width: 260px;
    height: 84px; /* 必须 */
    overflow: hidden; /* 必须 */
    margin: 1px auto;
    margin-left: 1px;
    border: 1px solid #f0f0f0;
    /*border:none;*/
    text-align: left;
}
ul {
    list-style: none;
}
li {
    width: 250px;
    font-size: 13px;
}
</style>
