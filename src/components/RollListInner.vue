<template>
    <div v-bind:id="ranId1">
        <ul v-bind:id="ranId2">
            <li v-for="item in vData">
                <p v-html="item"></p>
            </li>
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
        hasBorder: Boolean,
        vData: {},
    },
    data() {
        // 定义变量
        return {
            ranId1: 'rolldiv-' + this.random_id(),
            ranId2: 'ul1-' + this.random_id(),
            ranId3: 'ul2-' + this.random_id(),
            result: { data: [] },
            timer: 0
        };
    },
    methods: {
        random_id: function() {
            var tmpDate = new Date();
            var tmp = tmpDate.getTime();
            return tmp;
        },
        roll: function(t) {
            var ul1 = document.getElementById(this.ranId2);
            var ul2 = document.getElementById(this.ranId3);
            var ulbox = document.getElementById(this.ranId1);
            ul2.innerHTML = ul1.innerHTML;
            ulbox.scrollTop = 0; // 开始无滚动时设为0
            // 设置定时器，参数t用在这为间隔时间（单位毫秒），参数t越小，滚动速度越快
            this.timer = setInterval(this.rollStart, t);
            // 鼠标移入div时暂停滚动
            ulbox.onmouseover = function() {
                clearInterval(this.timer);
            };
            // 鼠标移出div后继续滚动
            ulbox.onmouseout = function() {
                this.timer = setInterval(this.rollStart, t);
            };
        },
        rollStart: function() { // 开始滚动函数
            // 上面声明的DOM对象为局部对象需要再次声明
            var ul1 = document.getElementById(this.ranId2);
            var ul2 = document.getElementById(this.ranId3);
            var ulbox = document.getElementById(this.ranId1);
            // scrollTop为空时进行重置
            if(ulbox === null || ulbox.scrollTop === null) ulbox.scrollTop=0;
            // 正常滚动不断给scrollTop的值+1,当滚动高度大于列表内容高度时恢复为0
            if (ulbox.scrollTop >= ul1.scrollHeight) {
                ulbox.scrollTop = 0;
            } else {
                ulbox.scrollTop++;
            }
        }
    },
    mounted() {
        this.roll(50);
    },
    beforeUpdate() { // 当data更新时触发
        clearInterval(this.timer);
        //var ulbox = document.getElementById(this.ranId1);
        //if(ulbox === null || ulbox.scrollTop === null) ulbox.scrollTop=0;
        this.roll(50);
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
* {
    margin: 0;
    padding: 0;
}
div {
    width: -webkit-max-content;
    height: 84px; /* 必须 */
    overflow: hidden; /* 必须 */
    margin: 1px auto;
    margin-left: 1px;
    /*border: 1px solid #f0f0f0;*/
    /*border:none;*/
    text-align: left;
}
ul {
    list-style: none;
}
li {
    width: 250px;
    font-size: 13px;
    display:block; /* NOTE: block属性触发里面的span布局生效. list-item属性增加列表标记 */
}
</style>
