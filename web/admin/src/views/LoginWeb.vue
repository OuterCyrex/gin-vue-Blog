<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model ref="loginFormRef" :rules="rules" :model="formData" class="loginForm">
        <a-form-model-item  prop="username" class="loginFormItem">
          <a-input  v-model="formData.username" placeholder="用户名">
            <template #prefix><a-icon type="user" style="color:rgba(0,0,0,.25)"/></template>
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password" class="loginFormItem">
          <a-input v-model="formData.password" placeholder="密码" type="password" v-on:keyup.enter="login">
            <template #prefix><a-icon type="lock" style="color: rgba(0, 0, 0, 0.25)" /></template>
          </a-input>
        </a-form-model-item>
        <a-form-model-item class="loginFormItem loginButton">
          <a-button type="primary" @click="login" style="margin:0 10px">登录</a-button>
          <a-button type="info" @click="resetForm">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return{
      formData:{
        username: '',
        password: '',
      },
      rules:{
        username:[
          {required:true,message:"请输入用户名",trigger:"blur"},
          {min:4,max:12,message: "用户名应为 4 至 12 个字符",trigger:"blur"}
        ],
        password:[
          {required:true,message:"请输入密码",trigger:"blur"},
          {min:6,max:20,message: "密码应为 6 至 20 个字符",trigger:"blur"}
        ]
      }
    }
  },
  methods:{
    resetForm(){
      this.$refs.loginFormRef.resetFields();
    },
    login(){
      this.$refs.loginFormRef.validate(async valid => {
        if(!valid) return this.$message.error("登录信息无效");
        const { data: res } = await this.$http.post("login",this.formData)
        if (res.status !== 200)return this.$message.error(res.msg)
        window.sessionStorage.setItem("token", res.token)
        this.$router.push("/admin/index")
      })
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color:#f4f4f4;
  display: flex;
  justify-content: center;
  align-items: center;
}
.loginBox{
  padding-top:50px;
  width:450px;
  height:300px;
  background-color: white;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}
.loginFormItem{
  padding:0 60px;
  align-self: center;
  box-sizing: border-box;
}
.loginForm{
  display: grid;
  grid-template-rows:1fr 1fr 1fr;
  align-items: center;
}
.loginButton{
  display: flex;
  justify-content: flex-end;
}
</style>
