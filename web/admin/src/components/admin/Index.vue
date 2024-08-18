<template>
  <div>
    <a-card style="background-image: url('https://img.oulu.me/006%20Lady%20Lips.png')">
      <div style="font-size:calc(20px + 1vw);display: flex;justify-content: center" >OuterBlog管理后台</div>
      <br>
      <div style="font-size:calc(10px + 0.5vw);display: flex;justify-content: center" >当前页面正在开发中</div>
    </a-card>
    <a-card>
      <a-list style="font-size:20px">
        <span>博客数据</span>
        <a-list-item>
          <div style="margin-left:20px;margin-top:20px">
            用户总数：
            <span style="color:orangered">{{this.totalList.userTotal}}</span>
          </div>
        </a-list-item>
        <a-list-item>
          <div style="margin-left:20px">
            文章总数：
            <span style="color:orangered">{{this.totalList.articleTotal}}</span>
          </div>
        </a-list-item>
        <a-list-item>
          <div style="margin-left:20px">
            分类总数：
            <span style="color:orangered">{{this.totalList.categoryTotal}}</span>
          </div>
        </a-list-item>
        <a-list-item>
          <div style="margin-left:20px">
            评论总数：
            <span style="color:orangered">{{this.totalList.commentTotal}}</span>
          </div>
        </a-list-item>
      </a-list>
    </a-card>
    <a-card>
      <div style="font-size:calc(10px + 0.5vw);display: flex;justify-content: center" >其实再怎么改也就这样了，不想动这个后台了</div>
    </a-card>
  </div>
</template>
<script>
export default {
  data(){
    return{
      totalList:{
        categoryTotal:0,
        articleTotal:0,
        userTotal:0,
        commentTotal:0,
      }
    }
  },
  created() {
    this.getCategoryList()
    this.getArticleList()
    this.getUserList()
    this.getCommentList()
  },
  methods:{
    async getCategoryList(){
      const { data : res } = await this.$http.get("categories",{
        params:{
          name:'',
          pagesize: 1000,
          pagenum: 1,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.totalList.categoryTotal = res.total
    },
    async getArticleList(){
      const { data : res } = await this.$http.get("articles", {
        params: {
          title: '',
          pagesize: 1000,
          pagenum: 1,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.totalList.articleTotal = res.total
    },
    async getUserList(){
      const { data : res } = await this.$http.get("users", {
        params: {
          username:'',
          pagesize:1000,
          pagenum: 1,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.totalList.userTotal = res.total
    },
    async getCommentList(){
      const { data : res } = await this.$http.get("comment", {
        params: {
          pagesize:1000,
          pagenum:1,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.totalList.commentTotal = res.total
    },
  }
}
</script>
