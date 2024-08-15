
<template>
  <div>
    <v-app-bar app color="indigo darken-2" >
      <v-container class="py-0 fill-height">
        <v-avatar @click="LoginCouldSee" size="40" class="mx-5" color="grey">
          <v-img v-if="token" src="../assets/defaultAvatar.png"></v-img>
        </v-avatar>
        <v-btn text color="white" @click="$router.push('/')">首页</v-btn>
        <v-btn text color="white" @click="$router.push('/articles')">文章列表</v-btn>
        <v-btn v-for="item in categoryList" :key="item.id" text color="white" @click="$router.push(`/articles/${item.id}`)">{{item.name}}</v-btn>
      </v-container>
        <v-spacer></v-spacer>
        <v-responsive min-width="300" max-width="400" color="white">
          <v-text-field dense flat hide-details solo-inverted rounded dark
                        placeholder="根据标题查询"
                        v-model="queryParam.name"
                        v-on:keyup.enter="sendDataToParent"
                        append-icon="mdi-close"
                        @click:append="clearInput"
          ></v-text-field>
        </v-responsive>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  data(){
    return{
      queryParam:{
        name:'',
        pagesize:5,
        pagenum:1,
      },
      categoryList:[],
      token : window.sessionStorage.token
    }
  },
  created() {
    this.GetCategoryList()
  },
  methods:{
    async GetCategoryList(){
      const {data:res}= await this.$http.get('categories',{
        params:
            {
              name:'',
              pagesize:1000,
              pagenum:1,
            },
      });
      this.categoryList=res.data;
    },
    sendDataToParent() {
      this.$router.push('/articles')
      this.$emit('update:data', this.queryParam.name);
    },
    clearInput(){
      this.queryParam.name=''
      this.sendDataToParent()
    },
    LoginCouldSee(){
      this.$emit('update:Login', true);
    }
  }
}
</script>
