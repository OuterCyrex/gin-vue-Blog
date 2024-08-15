<template>
  <v-main style="background-color: rgba(0, 0, 0, 0.5);height:100vh">
<v-container  style="height:100vh;margin-top:20vh">
  <v-row justify="center">
    <v-col cols="4">
      <v-card style="border-radius:10px">
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
    }
  },
  methods:{
    clear(){
      this.$refs.formRef.reset()
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
        this.$emit('update:data', this.loginValid);
      } catch (error) {
        this.snackbar.value = true
        this.snackbar.text = error;
      }
    },
  }
}
</script>

