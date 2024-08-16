<template>
  <div>
    <div class="mainTitle">友链列表</div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button type="primary" @click="addFriendLinkVisible = true">新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="username" :columns="columns" :pagination="pagination" :dataSource="FriendList" @change="handleTableChange" bordered style="margin-top:40px;">
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="danger" icon="delete" @click="deleteFriend(data.id)" style="margin:0 10px;">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!--新增标签-->
    <a-modal
        title="新增用户"
        :visible="addFriendLinkVisible"
        @ok="addFriendOk"
        @cancel="addFriendCancel"
        destroyOnClose >
      <a-form-model :model="newFriend" :rules="newFriendRule" ref="addFriendRef">
        <a-form-model-item has-feedback label="网站名" prop="name">
          <a-input v-model="newFriend.name"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="链接" prop="link">
          <a-input v-model="newFriend.link"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>
<script>
const columns = [
  {
    title:'ID',
    dataIndex:'id',
    width:'10%',
    key:'id',
    align:"center",
  },
  {
    title:"网站名",
    dataIndex:"name",
    width:"20%",
    key:"name",
    align:"center",
  },
  {
    title:"链接",
    dataIndex:"link",
    width:"20%",
    key:"link",
    align:"center",
  },
  {
    title:"操作",
    width:"30%",
    key:"action",
    align:"center",
    scopedSlots:{customRender:'action'},
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
      queryParam:{
        pagesize:5,
        pagenum:1,
      },
      FriendList:[],
      columns,
      newFriend:{
        name:'',
        link:'',
      },
      visible:false,
      addFriendLinkVisible:false,
      newFriendRule:{
        name:[{validator:(rule,value,callback)=>{
            if(this.newFriend.name === ""){
              callback(new Error("请输入网站名"))
            }
            callback()
          },trigger:"blur"}],
        link:[{validator:(rule,value,callback)=>{
            if(this.newFriend.link === ""){
              callback(new Error("请输入链接"))
            }
            callback()
          },trigger:"blur"}],
      },
    }
  },
  created(){
    this.getFriendList()
  },

  methods:{
    async getFriendList(){
      const { data : res } = await this.$http.get("friend", {
        params: {
          pagesize:this.queryParam.pagesize,
          pagenum:this.queryParam.pagenum,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.FriendList= res.data
      this.pagination.total = res.total
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
      this.getFriendList()
    },

    async deleteFriend(id){
      this.$confirm({
        title:"确定要删除吗？",
        content:"一旦删除将不可恢复，真的会消失很久！",
        onOk:async()=>{
          const res = await this.$http.delete(`friend/${id}`)
          if(res.data.status !== 200) {
            this.$message.error("操作失败")
          }else{
            this.$message.success("操作成功")
          }
          this.getFriendList()
        },
        onCancel:()=>{
          this.$message.info("取消成功")
        },
      })
    },


    addFriendOk(){
      this.$refs.addFriendRef.validate(async (valid)=>{
        if(!valid)return this.$message.error("请填写正确信息")
        const { data:res } = await this.$http.post('friend/add', {
          name: this.newFriend.name,
          link:this.newFriend.link
        })
        if(res.status !== 200) return this.$message.error(res.message)
        this.$message.success("创建友链成功")
        this.addFriendCancel()
        this.getFriendList()
      })
    },

    addFriendCancel(){
      this.$refs.addFriendRef.resetFields()
      this.addFriendLinkVisible = false
    },
  }
}
</script>

<style scoped>
.mainTitle{
  font-size: calc(20px + 2vw);
  text-align: center;
  align-content: center;
  background-color: white;
  font-family:"方正粗黑宋简体",sans-serif;
  user-select: none;
}
</style>
