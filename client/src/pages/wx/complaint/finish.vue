<template>
  <div>
    <action-bar :OPTIONS='headerOptions'></action-bar>
    <div :class='s.body'>
      <i class='iconfont icon-succ'></i>
      <p>提交成功</p>
      <h4>等待人员处理</h4>
      <!-- <div :class='s.btn' @click='onFinish'>好的</div> -->
    </div>
  </div>
</template>

<script>
import ActionBar from '@/components/mobile/ActionBar'
export default {
  components: {ActionBar},
  data () {
    return {
      headerOptions: {
        title: '完成'
      }
    }
  },
  mounted () {},
  methods: {
  }
}
</script>

<style lang="less" module='s'>
.body{
  margin-top: 80px;
  font-size: 30px;
  padding-top: 10%;
  text-align: center;
  i{
    color: #1eb504;
    font-size: 60px;
  }
  .btn{
    background-color: #1eb504;
    font-size: 30px;
    color: #fff;
    text-align: center;
    padding: 15px 0;
    margin: 20px;
  }
}
</style>