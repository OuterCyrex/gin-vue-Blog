<template>
  <v-main style="background-color: rgba(0, 0, 0, 0.5);height:100vh">
    <v-container style="height:100vh;margin-top:20vh">
    <v-row justify="center">
      <v-col cols="4">
        <v-card v-if="!isLoggedIn" style="border-radius:10px">
        <div style="padding: 60px 50px;">
          <v-form ref="formRef" v-model="valid" lazy-validation>
        <v-text-field label="用户名" type="username" v-model="userInfo.username" :rules="nameRules"></v-text-field>
        <v-text-field label="密码" v-on:keyup.enter="login" type="password" v-model="userInfo.password" :rules="pwdRules"></v-text-field>
          <div style="display:flex;justify-content: flex-end">
          <v-btn @click="clear" color="error" style="margin-right:20px;color:black">取消</v-btn>
          <v-btn style="background-color: deepskyblue;color:white" @click="login">登录</v-btn>
            </div>
              </v-form>
          </div>
        </v-card>
        <v-card v-if="isLoggedIn" style="border-radius:10px">
          <v-card-title class="ma-0">登录信息</v-card-title>
          <div style="padding: 0 50px;">
            <div style="font-size:20px;" class="mb-2">用户名：{{this.user.username}}</div>
            <span  style="font-size:20px;">权限：</span>
            <v-chip width="100%" :color="this.user.role === 1 ? 'pink':'blue'" label class="white--text">{{this.user.role === 1? '管理员' : '用户'}}</v-chip>
          </div>
          <v-btn class="ma-5" @click="quitLogin" color="error" style="margin-right:20px;color:black">退出登录</v-btn>
          <v-btn class="ml-2" @click="quitInfo" color="primary" style="margin-right:20px;color:black">确认</v-btn>
        </v-card>
      </v-col>
    </v-row>
    <v-snackbar
      v-model="snackbar.value"
      :timeout="2000"
      top
      style="color:white"
    ><v-icon style="margin-right:20px">mdi-cancel</v-icon>{{snackbar.text}}</v-snackbar>
  </v-container>
  </v-main>
</template>

<script>

export default {
  data(){
    return{
      valid: true,
      userInfo:{
        username:'',
        password:'',
      },
      snackbar: {
        value: false,
        text: '',
      },
      nameRules: [
        v => !!v || '请输入用户名',
        v => (v && v.length <= 12 && v.length>=4) || '用户名应为 4 至 12 个字符',
      ],
      pwdRules:[
        v => !!v || '请输入密码',
        v => (v && v.length <= 12 && v.length>=4) || '密码应为 6 至 20 个字符',
      ],
      loginValid:false,
      user:{},
      isLoggedIn:false,
    }
  },
  created(){
    this.getUserInfo()
  },
  methods:{
    clear(){
      this.$refs.formRef.reset()
      this.$emit('update:data', this.loginValid);
    },
    quitInfo(){
      this.$emit('update:data', this.loginValid);
    },
    quitLogin(){
      window.sessionStorage.setItem("token",'');
      this.isLoggedIn = false;
      this.$emit('update:isLoggedIn', this.isLoggedIn);
      this.$emit('update:data', this.loginValid);
    },
    async login(){
      try {
        const isValid = await this.$refs.formRef.validate();
        if (!isValid) return;

        const response = await this.$http.post("userlogin", this.userInfo);
        if (response.data.status !== 200) {
          this.snackbar.value = true
          this.snackbar.text = response.data.msg;
          return;
        }
        window.sessionStorage.setItem("token", response.data.token);
        this.isLoggedIn = true;
        this.$emit('update:isLoggedIn', this.isLoggedIn);
        this.$emit('update:data', this.loginValid);
      } catch (error) {
        this.snackbar.value = true
        this.snackbar.text = error;
      }
    },
    async getUserInfo(){
      const token = window.sessionStorage.getItem("token");

      if(!token)return;

      const config = {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      };

      const {data:res} = await this.$http.get("userinfo",config)

      if(res.status!==200){
        this.isLoggedIn = false;
        return
      }
      this.isLoggedIn = true;
      this.user = res.data;
    },
  }
}
</script>

