<template>
  <div>
    <h3>用户列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="搜索用户" enter-button/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="username" :columns="columns" :pagination="pagination" :dataSource="userList" @change="handleTableChange" style="margin-top:40px;"></a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title:'ID',
    dataIndex:'ID',
    width:'10%',
    key:'id',
  },
  {
    title:"用户名",
    dataIndex:"username",
    width:"20%",
    key:"username",
  },
  {
    title:"权限",
    dataIndex:"role",
    width:"20%",
    key:"role",
  },
  {
    title:"操作",
    width:"30%",
    key:"action",
  },
]

export default {
  data(){
    return{
      pagination:{
        pageSizeOptions: ['5','10','20'],
        pageSize:5,
        total:0,
        showSizeChanger:true,
        showTotal:(total)=>`共${total}条`,
      },
      userList:[],
      columns,
      queryParam:{
        pagesize:5,
        pagenum:1,
      },
      visible:false,
    }
  },
  created(){
    this.getUserList()
  },
  methods:{
    async getUserList(){
      const { data : res } = await this.$http.get("users", {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.userList = res.data
      this.pagination.total = res.total
      console.log(res)
    },

    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
  }
}
</script>
