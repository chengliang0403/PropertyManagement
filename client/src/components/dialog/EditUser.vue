<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      UserName
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.UserName" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item' v-if='user.Type == 3'>
      <div :class='s.red'>*</div>
      Street
      <!-- 用户名 -->
        <el-select :class='s.elInput' v-model="selectedStreetID" placeholder="请选择">
          <el-option
            v-for="item in streets"
            :key="item.ID"
            :label="item.Name"
            :value="item.ID">
          </el-option>
        </el-select>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      RealName
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.RealName" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      Password
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.Password" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      OfficeNumber
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.Tel" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item'>
      <el-button @click='onSave' type='primary'>提 交</el-button>
      <el-button @click='onCancel'>取 消</el-button>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  props: {
    USER: Object
  },
  data () {
    return {
      warn: '',
      user: {
        UserName: '',
        RealName: '',
        Password: '',
        StreetID: '',
        Tel: ''
      },
      selectedStreetID: '',
      streets: []
    }
  },
  mounted () {
    Object.assign(this.user, this.USER)
    this.getAllStreets()
    this.selectedStreetID = this.user.StreetID
  },
  methods: {
    onSave () {
      if (this.selectedStreetID != '') this.user.StreetID = this.selectedStreetID
      if(!this.checkUser()) return
      fetchpm(this, true, '/pm/user/update', {
        method: 'POST',
        body: this.user
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onSave', body)
        if (body.error == 1) this.warn = body.data
        else {
          this.warn = '修改成功'
          this.$emit('editSucc')
        }
      })
    },
    onCancel () {
      this.$emit('cancel')
    },
    onFocus () {
      this.warn = ''
    },
    checkUser () {
      if (this.user.UserName == '') {
        this.warn = '用户名不能为空'
        return false
      } 
      if (this.user.Password == '') {
        this.warn = '密码不能为空'
        return false
      }
      return true
    },
    getAllStreets () {
      fetchpm(this, true, '/pm/street',{
          method: 'POST'
      }).then(resp => {
        console.info(resp)
        return resp.json()
      }).then( data => {
        if (data.error === 0) {
          console.info (data)
          this.streets = data.data
        }
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  width: 100%;
  .warn{
    text-align: center;
    color: red;
  }
  .item{
    display: flex;
    justify-content: flex-end;
    align-items: center;
    padding: 10px 0px;
    font-size: 18px;
    .red{
      color: red;
    }
    .elInput{
      width: 75%;
      margin-left: 10px
    }
  }
}  
</style>