<template>
  <v-card class="mx-auto" max-width="300">
    <v-card-title>Friend</v-card-title>
    <v-divider></v-divider>
    <v-card-subtitle v-for="item in friendList" :key="item.id">
      <a :href="item.link" class="link">{{ item.name }}</a>
    </v-card-subtitle>
  </v-card>
</template>

<script>
export default {
  data(){
    return{
      queryParam:{
        pagesize:100,
        pagenum:1,
      },
      friendList:[],
    }
  },
  created() {
    this.getFriendList();
  },
  methods:{
    async getFriendList(){
      const {data:res} = await this.$http.get("friend",{
        params:{
          pagesize:this.queryParam.pagesize,
          pagenum:this.queryParam.pagenum
        }
      })
      this.friendList = res.data;
    },
  }
}
</script>

<style scoped>
.link {
  color: grey;
  text-decoration: none;
}

.link:hover {
  color: deepskyblue;
  text-decoration: underline;
}
</style>
