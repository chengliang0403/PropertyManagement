<template>
  <div :class='s.wrap'>
    <div :class='s.title'>
      <img src="~@/res/images/earth.png">
      Event handle Recs<!-- 居民物业纠纷处理 -->
    </div>
    <div :class='s.body'>
      <table>
        <!-- 提交人类别  提交人用户名  提交时间  提交文本说明  操作 -->
        <tr>
          <th>AuthorCategory</th>
          <th>AuthorName</th>
          <th>Time</th>
          <th>HandleInfo</th>
          <th>操作</th>
        </tr>
        <tr v-for='handle in eventHandles'>
          <td v-text='handle.AuthorCategory'></td>
          <td v-text='handle.AuthorName'></td>
          <td>{{ handle.Time | filterTime }}</td>
          <td v-text='handle.HandleInfo'></td>
          <td>
            <el-button type="primary" icon="search" @click='onHandleDetails(handle)'>详情</el-button>
          </td>
        </tr>
      </table>
      <el-dialog 
        v-model="showHandleDetails" 
        size="tiny" 
        title='Event Handle Details'>
        <details-event-handle 
          v-if='showHandleDetails'
          :eventHandle='showingHandle'></details-event-handle>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import DetailsEventHandle from '@/components/dialog/DetailsEventHandle'
import fetchpm from '@/fetchpm'
import filterTime from '@/filters/filterTime'
export default {
  components: { DetailsEventHandle },
  filters: { filterTime },
  data () {
    return {
      showHandleDetails: false,
      eventHandles: [],
      showingHandle:{}
    }
  },
  mounted () {
    var user = JSON.parse(sessionStorage.getItem('user')) || {}
    this.fetchEventHandles(user.UserName, user.type)
  },
  methods: {
    onHandleDetails (eventHandle) {
      this.showingHandle = eventHandle
      this.showHandleDetails = true
    },
    fetchEventHandles (AuthorName, AuthorCategory) {
      fetchpm(this, true, '/pm/eventHandle/kvs', {
        method: 'POST',
        body: {AuthorName: AuthorName, AuthorCategory: AuthorCategory}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchEventHandles', body)
        if (body.error === 0) {
          this.eventHandles = body.data || []
        }
        
      })
    }
  }
}
</script>

<style lang="less" module='s'>
.wrap{
  margin: 10px;
  border: solid 1px #4c87b9;
  margin: 10px;
  background-color: #fff;
  .title{
    color: #fff;
    font-size: 20px;
    background: #4c87b9;
    padding: 10px;
    display: flex;
    align-items: center;
    img{
      width: 20px;
      margin-right: 10px;
    }
  }
  .body{
    margin: 10px;
    table{
      width: 100%;
      font-size: 15px;
      color: #555;
      background-color: #fff;

      th, td {
        padding: 5px;
        border: solid 1px #ddd;
        text-align: center;
      }
      th{
        background-color: #f0f0f0;
      }
      tr{
        &:hover {
          background-color: #f0f0f0;
        }
      }
    }
  }
}
</style>