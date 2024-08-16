<template>
  <div>
  <v-card class="pb-10">
    <div class="d-flex justify-center text-h4 font-weight-bold pt-3 pb-10">{{articleInfo.title}}</div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-4" elevation="1" color="indigo" dark border="left" outlined>{{articleInfo.desc}}</v-alert>
    <v-card-text>
      <div class="ArticleHTML" v-html="articleInfo.html"></div>
    </v-card-text>

    <v-divider></v-divider>
    <v-card class="ma-10">
      <v-card class="pa-0" color="#f4f4f4" outlined>
        <div class="d-flex justify-center" style="background-color: #00E5FF; color:white;font-size:calc(20px + 1vw);padding:10px">文章评论区</div>
        <CommentWeb :id="this.id"></CommentWeb>
      </v-card>

    </v-card>
  </v-card>
  </div>
</template>

<script>
import CommentWeb from "@/components/CommentWeb.vue";

export default {
  components: {CommentWeb},
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
.ArticleHTML >>> td{
  padding:10px;
  border: 1px solid black;
  border-collapse: collapse;
}
.ArticleHTML >>> table{
  margin:40px auto;
  padding:4px;
  border: 1px solid black;
  border-collapse: collapse;
}
.ArticleHTML >>> th{
  padding:10px;
  border: 1px solid black;
  border-collapse: collapse;
}
.ArticleHTML >>> img{
  width:90%;
}
</style>
