<template>
  <div>
    <div class="mainTitle">文章列表</div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search v-model="queryParam.title" placeholder="搜索标签" enter-button allowClear @search="getArticleList"/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="$router.push('/admin/addarticle')">新增</a-button>
        </a-col>
        <a-col :span="6" :offset="4">
          <a-select allowClear style="width:200px;" placeholder="请选择标签" @change="CategoryChange" >
            <a-select-option v-for="item in categoryList" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
          </a-select>
        </a-col>
      </a-row>
      <a-table rowKey="ID" :columns="columns" :pagination="pagination" :dataSource="articleList" @change="handleTableChange" bordered style="margin-top:40px;">
        <span slot="category" slot-scope="category">
          <a-tag :color=" category.length > 8 ? 'orange' : 'blue' " size="big">{{ category }}</a-tag>
        </span>
        <span class="articleImg" slot="img" slot-scope="img">
          <img :src="img">
        </span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin:0 10px;" size="small" @click="$router.push(`/admin/addarticle/${data.ID}`)">编辑</a-button>
            <a-button type="danger" icon="delete" @click="deleteArticle(data.ID)" style="margin:0 10px;" size="small">删除</a-button>
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
    key:'ID',
    align:"center",
  },
  {
    title:"标题",
    dataIndex:"title",
    width:"20%",
    key:"title",
    align:"center",
  },
  {
    title:"分类",
    dataIndex: "Category.name",
    width:"10%",
    key:"category",
    align:"center",
    scopedSlots:{customRender:'category'},
  },
  {
    title:"描述",
    dataIndex:"desc",
    width:"20%",
    key:"desc",
    align:"center",
  },
  {
    title:"缩略图",
    dataIndex: "img",
    width:"10%",
    key:"img",
    align:"center",
    scopedSlots:{customRender:'img'},
  },
  {
    title:"操作",
    width:"10%",
    key:"action",
    align:"center",
    scopedSlots:{customRender:'action'},
  }
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
      articleList:[],
      categoryList:[],
      columns,
      queryParam:{
        title:'',
        pagesize:5,
        pagenum:1,
      },
    }
  },
  created(){
    this.getArticleList()
    this.getCategoryList()
  },

  methods:{
    async getArticleList(){
      const { data : res } = await this.$http.get("articles", {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.articleList = res.data
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
      this.getArticleList()
    },

    async deleteArticle(id){
      this.$confirm({
        title:"确定要删除吗？",
        content:"一旦删除将不可恢复，真的会消失很久！",
        onOk:async()=>{
          const res = await this.$http.delete(`article/${id}`)
          if(res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.getArticleList()
          this.$message.success("操作成功")
        },
        onCancel:()=>{
          this.$message.info("取消成功")
        },
      })
    },

    async getCategoryList(){
      const { data : res } = await this.$http.get("categories",{
      params:{
        name:'',
        pagesize: 1000,
        pagenum: 1,
      }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.categoryList = res.data
      this.pagination.total = res.total
    },

    CategoryChange(value){
      this.getArticleByCategory(value)
    },

    async getArticleByCategory(id){
      const {data:res} = await this.$http.get(`article/category/${id}`,{
        params: {
          title:'',
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        }
      })
      if (res.status !== 200)return this.$message.error(res.message)
      this.articleList = res.data
      this.pagination.total = res.total
    }
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
.articleImg img{
  width: 150px;
  height:100px;
}
</style>
