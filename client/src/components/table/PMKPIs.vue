<template>
  <div :class='s.wrap'>
    <div :class='s.content'>
      <div :class='s.title'>
        <img src="~@/res/images/earth.png">
        PM KPI<!-- 物业考核查询-->
      </div>
      <div :class='s.body'>
        <table>
          <tr>
            <td :class='s.key'>Year</td>
            <td>
              <el-date-picker
                v-model="selectedYear"
                type="year"
                placeholder="选择年份">
              </el-date-picker>
            </td>
            <!-- 季度 -->
            <td :class='s.key'>Quarter</td>
            <td>
              <el-select v-model="selectedQuarter" filterable placeholder="select Quarter">
                <el-option
                  v-for="item in quarters"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>Company Name</td>
            <td>
              <el-input v-model="searchCompanyName" placeholder="请输入Company Name"></el-input>
            </td>
          </tr>
          <tr>
            <td :class='s.key'>Street</td>
            <td>
              <el-select v-model="selectedStreetID" filterable placeholder="select Street">
                <el-option
                  v-for="item in streets"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>Community</td>
            <td>
              <el-select v-model="selectedCommunityID" filterable placeholder="select Community">
                <el-option
                  v-for="item in communities"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
            <td :class='s.key'>XQ</td>
            <td>
              <el-select v-model="selectedXQID" filterable placeholder="select XQ">
                <el-option
                  v-for="item in xqs"
                  :key="item.ID"
                  :label="item.Name"
                  :value="item.ID">
                </el-option>
              </el-select>
            </td>
          </tr>
        </table>
        <div :class='s.searchWrap'>
          <el-button type="primary" @click='onCurQuarter'>Cur Quarter</el-button>
          <el-button type="primary" icon="search" @click='onInputSearch'>查询</el-button>
        </div>
        <!-- search result -->
        <table>
        <!-- 警告类型 事件编号  开始时间  所在小区  事件状态  事件等级  事件类别  操作 -->
          <tr>
            <th>Year</th>
            <th>Quarter</th>
            <th>Company Name</th>
            <th>Red Warning Num</th>
            <th>Yellow Warning Num</th>
            <th>Important Warning Num</th>
            <th>Score</th>
          </tr>
          <tr v-for='item in KPIs'>
            <td v-text='item.Year'></td>
            <td v-text='item.Quarter'></td>
            <td v-text='item.Name'></td>
            <td v-text='item.YWNo'></td>
            <td v-text='item.RWNo'></td>
            <td v-text='item.IWNo'></td>
            <td v-text='item.Score'></td>
          </tr>
          <tr v-if='KPIs.length === 0'>
            <td colspan="7">无记录</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script type="text/javascript">
import filterEventStatus from '@/filters/filterEventStatus'
import filterEventLevel from '@/filters/filterEventLevel'
import fetchpm from '@/fetchpm'
export default {
  filters: {filterEventStatus, filterEventLevel},
  data () {
    return {
      host: '//localhost:3000',
      quarters: [{value:0, label:'全部'}, {value:1, label:'1'},{value:2,label:'2'},{value:3, label:'3'},{value:4, label: '4'}],
      selectedYear: 0,
      selectedQuarter: 0,
      searchCompanyName: '',
      selectedStreetID:'',
      selectedCommunityID: '',
      selectedXQID: '',
      xqs:[],
      communities: [],
      streets: [],
      KPIs:[]
    }
  },
  mounted () {
    this.selectedYear = new Date()
    this.fetchKPIs('')
    this.fetechAllStreets()
  },
  watch: {
    selectedStreetID: function (value) {
      this.selectedCommunityID = ''
      this.fetchCommunitiesByStreetID(value)
    },
    selectedCommunityID: function (value) {
      this.selectedXQID = ''
      this.fetchXQByCommunityID(value)
    }
  },
  methods: {
    onCurQuarter () {
      var today = new Date()

      var data = {
        Year: today.getFullYear(),
        Quarter: (parseInt(today.getMonth() / 3) + 1)
      }
      this.onSearch(data)
    },
    onInputSearch () {
      this.onSearch(this.formatSearchData())
    },
    fetechAllStreets() {
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
    },
    fetchCommunitiesByStreetID (streetID) {
      if (!streetID) return
      this.isLoadingInput = true
      this.communities = []
      this.inputCommunityID = ''
      fetchpm(this, true, '/pm/community/kvs', {
        method: 'POST',
        body: {streetID: streetID}
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('fetchCommunitiesByStreetName', body)
        if(body.error !== 0) {
          console.error("Error: search CommunitiesByStreetName. Reason:" + body.data)
        } else if (body.data !== null ) {
          this.communities = body.data
        }
        this.isLoadingInput = false
      })
    },
    fetchXQByCommunityID (communityID) {
      if (!communityID) return
      this.isLoadingInput = true
      fetchpm(this, true, '/pm/xq/kvs', {
        method: 'POST',
        body: {communityID: communityID}
      }).then(resp => {
        return resp.json()
      }).then( body => {
        console.info('fetchXQByCommunityName', body)
        if(body.error !== 0) console.error("Error: search fetchXQByCommunityName. Reason:" + body.data)
        else if (body.data !== null ) {
          this.xqs = body.data
        }
        this.isLoadingInput = false
      })
    },
    onSearch (data) {
      fetchpm(this, true, '/pm/pmkpi/kvs', {
        method: 'POST',
        body: data
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('onSearch', body)
        if (body.error === 0) {
          this.KPIs = body.data || []
        } else {
          this.KPIs = []
        }
      })
    },
    formatSearchData () {
      var searchData = {}
      if (this.selectedYear !== -1 && this.selectedYear !== '') searchData.Year = this.selectedYear.getFullYear()
      if (this.selectedQuarter !== 0) searchData.Quarter = this.selectedQuarter
      if (this.searchCompanyName !== '') searchData.Name = this.searchCompanyName
      if (this.selectedStreetID !== '') searchData.StreetID = this.selectedStreetID
      if (this.selectedCommunityID !== '') searchData.CommunityID = this.selectedCommunityID
      if (this.selectedXQID !== '') searchData.XQID = this.selectedXQID
      console.info('formatSearchData', searchData)
      return searchData

    },
    fetchKPIs (name) {
      if (!name ) {
        console.info('name', name)
      }
      fetchpm(this,true,'/pm/pmkpi', {
        method: 'POST',
        body: {name: name}
      }).then(resp => {
        return resp.json()
      }).then(body => {
        console.info('fetchKPIs', body)
        this.KPIs = body.data
      })
    },
    toDetails (event) {
      event.index = 1
      console.info('toDetails', event.Index)
      this.$router.push({path:'/street/detail/', params:{index: '1'}})
      // this.$router.push({path:'/street/detail/' + event.Index})
    }
  }
}
</script>

<style lang='less' module='s'>
.wrap{
  .content{
    border: solid 1px #4c87b9;
    margin-top: 50px;
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
      .searchWrap{
        display: flex;
        justify-content: flex-end;
        margin: 10px;
      }
      table{
        width: 100%;
        margin-top: 10px;
        font-size: 15px;
        color: #555;
        background-color: #fff;
        th, td {
          padding: 5px;
          border: solid 1px #ddd;
          text-align: center;
          font-size: 15px;
        }
        th{
          background-color: #f0f0f0;
        } 
        tr{
          &:hover {
            background-color: #f0f0f0;
          }
        }
        .key{
          background-color: #f0f0f0;
        }
      }
    }    
  }
}
</style>