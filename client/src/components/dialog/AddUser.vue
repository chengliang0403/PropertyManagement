<template>
  <div :class='s.wrap'>
    <div :class='s.warn' v-if='warn !== ""' v-text='warn'></div>
    <div :class='s.item'>
      <div :class='s.red'>*</div>
      UserName
      <!-- 用户名 -->
      <el-input :class='s.elInput' v-model="user.UserName" placeholder="请输入" @focus='onFocus'></el-input>
    </div>
    <div :class='s.item' v-if='USER_TYPE == 3'>
      <div :class='s.red'>*</div>
      StreetName
      <!-- 用户名 -->
      <el-select :class='s.elInput' v-model="user.StreetID" placeholder="请选择">
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
      <el-button @click='onSave' type='primary'>提交</el-button>
      <el-button @click='onCancel'>取消</el-button>
    </div>
  </div>
</template>

<script>
import fetchpm from '@/fetchpm'
export default {
  props: {
    USER_TYPE: Number
  },
  data () {
    return {
      warn:'',
      user: {
        UserName: '',
        Password: '',
        RealName: '',
        StreetID:'',
        Tel: '',
        Type: -1
      },
      streets:[]
    }
  },
  mounted () {
    this.user.Type = this.USER_TYPE
    this.fetchAllStreets()
  },
  methods: {
    resetData () {
      this.user = {
        UserName: '',
        Password: '',
        RealName: '',
        Tel: '',
        Type: this.USER_TYPE
      }
    },
    onFocus () {
      this.warn = ''
    },
    onSave () {
      if(!this.checkUser()) return
      this.addUser()
    },
    checkUser () {
      if (this.user.UserName == ''){
        this.warn = '用户名不能为空'
        return false
      }
      if (this.user.Password === ''){
        this.warn = '密码不能为空'
        return false
      }
      return true
    },
    onCancel () {
      this.$emit('cancel')
    },
    addUser () {
      fetchpm(this, true, '/pm/user/add', {
        method: 'POST',
        body: this.user
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('addUser', body)
        if (body.error === 1) {
          this.warn = body.data
        } else {
          this.$emit('addSucc')
          this.warn = '添加成功'
          this.resetData()
        }
      })
    },
    fetchAllStreets () {
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