<template>
  <v-col>
    <v-card class="ml-5 mr-5" v-for="item in commentList" :key="item.ID" outlined>
      <v-row no-gutters  v-if="commentList.length!==0">
        <v-col cols="7">
          <div class="mt-2">
          <v-avatar size="40" class="mx-5" color="grey">
            <v-img src="../assets/defaultAvatar.png"></v-img>
          </v-avatar>
            <v-chip :color="item.User.ID === 1 ? 'pink':'blue'" label class="white--text mr-3">{{item.User.ID === 1? '管理员' : '用户'}}</v-chip>
            {{item.User.username}}
          </div>
          <v-card-subtitle v-text="item.comments"></v-card-subtitle>
          <v-card-text>
            <v-icon class="mr-2">{{'mdi-calendar-month'}}</v-icon>
            <span>上传时间：{{item.CreatedAt.slice(0,10) + ' ' + item.CreatedAt.slice(11,19)}}</span>
          </v-card-text>
        </v-col>
      </v-row>
    </v-card>
    <v-row v-if="commentList.length===0" rows="12" class="pa-5">
      <v-col cols="12">
        <v-card class="pa-10">
          <v-card-title class="d-flex justify-center align-center"><v-icon>{{'mdi-error'}}</v-icon>暂时还没有人评论哦</v-card-title>
        </v-card>
      </v-col>
    </v-row>
    <div class="text-center mb-4 bt-4">
      <div class="mt-2 mb-2" style="font-size: 20px">共{{this.total}}条</div>
      <v-pagination total-visible="7" v-model="queryParam.pagenum" :length="Math.ceil(this.total/this.queryParam.pagesize)" @input="getCommentList"></v-pagination>
    </div>
    <v-divider></v-divider>
    <v-card class="ma-5">
      <div class="pt-4 pb-4" style="color:Black;font-size:calc(20px + 0.5vw);padding-left:40px">
        <v-icon class="mr-2" color="blue">mdi-message</v-icon>
        发表评论
      </div>
      <v-divider></v-divider>
      <v-textarea placeholder="请正常发言，不然会被管理员删除哦" required v-on:keyon.enter="pushComment" v-model="CommentInfo.comments" height="100px" class="pa-5 ma-5 pb-0">
      </v-textarea>
      <div class="d-flex justify-end pa-10 pt-0" >
        <v-btn @click="clear" color="error" style="margin-right:20px;color:black">取消</v-btn>
        <v-btn color="primary" @click="handleCommentSubmit" :disabled="this.submitTimeout !== null">{{this.submitTimeout ? '上传中':'提交'}}</v-btn>
      </div>
    </v-card>
    <v-snackbar
        v-model="snackbar.value"
        :timeout="2000"
        top
        style="color:white"
    ><v-icon style="margin-right:20px">mdi-cancel</v-icon>{{snackbar.text}}</v-snackbar>
  </v-col>
</template>
<script>
export default {
  props:['id'],
  data(){
   return{
     queryParam:{
       pagesize:5,
       pagenum:1,
     },
     total:0,
     commentList:[],
     CommentInfo:{
       aid:Number(this.id),
       comments:'',
     },
     snackbar:{
       value:false,
       text:'',
     },
     submitTimeout: null,
   }
  },
  created(){
    this.getCommentList();
  },
  methods:{
    handleCommentSubmit() {
      if (this.submitTimeout) {
        clearTimeout(this.submitTimeout);
      }
      this.submitTimeout = setTimeout(() => {
        this.pushComment();
        this.submitTimeout = null;
      }, 2000);
    },
    async getCommentList(){
      const {data:res} = await this.$http.get("comment/"+this.id,{
        params:{
          pagesize:this.queryParam.pagesize,
          pagenum:this.queryParam.pagenum
        }
      })
      this.commentList = res.data;
      this.total = res.total;
    },
    async pushComment(){

      const token = window.sessionStorage.getItem("token");

      if(!token){
        this.snackbar.value = true
        this.snackbar.text = '请先登录！';
        return;
      }

      const config = {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      };

      const {data:res} = await this.$http.post("comment",this.CommentInfo,config)
      this.commentList = res.data;
      if(res.status!==200){
        this.snackbar.value = true
        this.snackbar.text = '请先登录！';
        return;
      }
      this.getCommentList();
      this.total = res.total;
      this.clear()
    },
    clear(){
      this.CommentInfo.comments = ''
    },
  },
}
</script>
