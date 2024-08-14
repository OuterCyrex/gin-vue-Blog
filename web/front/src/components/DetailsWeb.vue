<template>
  <div>
    <div class="d-flex justify-center text-h4 font-weight-bold pt-3">{{articleInfo.title}}</div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-4" elevation="1" color="indigo" dark border="left" outlined>{{articleInfo.desc}}</v-alert>
    <v-card-text>
      <div class="ArticleHTML" v-html="articleInfo.html"></div>
    </v-card-text>
  </div>
</template>

<script>
export default {
  props:['id'],
  data(){
    return{
      articleInfo:{},
    }
  },
  created(){
    this.getArticleInfo()
  },
  methods:{
    async getArticleInfo(){
      const {data:res}= await this.$http.get('article/' + this.id)
      this.articleInfo = res.data
    }
  }
}
</script>

<style scoped>
.ArticleHTML >>> pre{
  font-size:20px;
  font-family:"Cascadia Code",sans-serif;
  margin:10px;
  color:white;
  background-color:#333;
  padding:20px;
  width:90%;
  border-radius:6px;
}
.ArticleHTML >>> .hljs-keyword{
  color:orange;
}
.ArticleHTML >>> .hljs-title{
  color:cornflowerblue;
}
.ArticleHTML >>> .hljs-type{
  color:greenyellow;
}
.ArticleHTML >>> .hljs-string{
  color:orangered;
}
.ArticleHTML >>> .hljs-literal{
  color:pink;
}
.ArticleHTML >>> .hljs-params{
  color:mediumpurple;
}
.ArticleHTML >>> .hljs-comment{
  color:green;
}
.ArticleHTML >>> .hljs-number{
  color:yellow;
}
.ArticleHTML >>> span{
  font-family:"Cascadia Code",sans-serif;
}
.ArticleHTML >>> h1{
  margin:20px 0;
}
.ArticleHTML >>> h2{
  margin:15px 0;
}
.ArticleHTML >>> h3{
  margin:10px 0;
}
.ArticleHTML >>> code{
  font-family: "Cascadia Code",sans-serif;
  font-size:12px;
}
</style>

