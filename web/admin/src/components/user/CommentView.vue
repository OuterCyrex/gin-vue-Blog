<template>
  <div>
    <div class="mainTitle">评论列表</div>
    <a-card>
      <a-table rowKey="id" :columns="columns" :pagination="pagination" :dataSource="CommentList" @change="handleTableChange" bordered style="margin-top:40px;">
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="danger" icon="delete" @click="deleteComment(data.ID)" style="margin:0 10px;">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>
<script>
const columns = [
  {
    title:'ID',
    dataIndex:'ID',
    width:'5%',
    key:'id',
    align:"center",
  },
  {
    title:"用户名",
    dataIndex:"User.username",
    width:"10%",
    key:"name",
    align:"center",
  },
  {
    title:"文章标题",
    dataIndex:"Article.title",
    width:"20%",
    key:"title",
    align:"center",
  },
  {
    title:"评论内容",
    dataIndex:"comments",
    width:"20%",
    key:"comments",
    align:"center",
  },
  {
    title:"操作",
    width:"20%",
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
      CommentList:[],
      columns,
      visible:false,
    }
  },
  created(){
    this.getCommentList()
  },

  methods:{
    async getCommentList(){
      const { data : res } = await this.$http.get("comment", {
        params: {
          pagesize:this.queryParam.pagesize,
          pagenum:this.queryParam.pagenum,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.CommentList= res.data
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
      this.getCommentList()
    },

    async deleteComment(id){
      this.$confirm({
        title:"确定要删除吗？",
        content:"一旦删除将不可恢复，真的会消失很久！",
        onOk:async()=>{
          const res = await this.$http.delete(`comment/${id}`)
          if(res.data.status !== 200) {
            this.$message.error("操作失败")
          }else{
            this.$message.success("操作成功")
          }
          this.getCommentList()
        },
        onCancel:()=>{
          this.$message.info("取消成功")
        },
      })
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
