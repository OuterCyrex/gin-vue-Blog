<template>
  <v-col>
    <v-card class="ma-3" v-for="item in articleList" :key="item.ID" link @click="$router.push(`/detail/${item.ID}`)">
      <v-row no-gutters  v-if="articleList.length!==0">
        <v-col class="d-flex justify-center align-center mx-3" cols="2">
          <v-img height="120px" width="80px" :src="item.img"></v-img>
        </v-col>
        <v-col cols="7">
          <v-card-title class="ma-3">
            <v-chip color="pink" label class="white--text mr-3">{{item.Category.name}}</v-chip>
            {{item.title}}
          </v-card-title>
          <v-card-subtitle v-text="item.desc"></v-card-subtitle>
          <v-divider></v-divider>
          <v-card-text>
            <v-icon class="mr-2">{{'mdi-calendar-month'}}</v-icon>
            <span>上传时间：{{item.CreatedAt.slice(0,10) + ' ' + item.CreatedAt.slice(11,19)}}</span>
          </v-card-text>
          <v-card-text></v-card-text>
        </v-col>
      </v-row>
    </v-card>
    <v-row v-if="articleList.length===0" rows="12" class="pa-5">
      <v-col cols="12">
      <v-card class="pa-10">
      <v-card-title class="d-flex justify-center align-center"><v-icon>{{'mdi-error'}}</v-icon>未查询到任何数据</v-card-title>
      </v-card>
      </v-col>
    </v-row>
    <div class="text-center">
      <v-pagination total-visible="7" v-model="queryParam.pagenum" :length="Math.ceil(this.total/this.queryParam.pagesize)" @input="getArticleList()"></v-pagination>
    </div>
  </v-col>
</template>

<script>
export default {
  props:{
    id:{
      type:String,
    },
    receivedData:{
      type:String,
      default:'',
    }
  },
  data(){
    return{
      articleList:[],
      total:0,
      queryParam:{
        pagesize:5,
        pagenum:1,
      }
    }
  },
  created(){
    if (this.id) {
      this.getArticleByCategory();
    }else{
      this.getArticleList();
    }
    console.log(this.articleList)
  },
  watch: {
    id(newVal) {
      if (newVal) {
        this.getArticleByCategory(newVal);
      } else {
        this.getArticleList();
      }
    },
    receivedData() {
      this.getArticleList()
    }
  },
  methods:{
    async getArticleList(){
      const {data:res} = await this.$http.get("articles",{
        params:{
          title:this.receivedData,
          pagesize:this.queryParam.pagesize,
          pagenum:this.queryParam.pagenum
        }
      })
        this.articleList = res.data;
        this.total = res.total;
    },
    async getArticleByCategory(){
      const {data:res} = await this.$http.get(`article/category/${this.id}`,{
        params:{
          pagesize:this.queryParam.pagesize,
          pagenum:this.queryParam.pagenum,
        }
      })
      this.articleList = res.data;
      this.total = res.total;
    }
  },
}
</script>
