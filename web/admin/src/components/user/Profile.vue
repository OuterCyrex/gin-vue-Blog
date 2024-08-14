<template>
<div>
  <a-card>
    <h3>个人设置</h3>
    <a-form-model>
      <a-form-model-item label="个人名称">
        <a-input style="width:300px" v-model="profileInfo.name"></a-input>
      </a-form-model-item>
      <a-form-model-item label="自我介绍" >
        <a-input type="textarea" v-model="profileInfo.desc"></a-input>
      </a-form-model-item>
      <a-form-model-item label="QQ号码">
        <a-input style="width:300px" v-model="profileInfo.qqchat"></a-input>
      </a-form-model-item>
      <a-form-model-item label="B站账号">
        <a-input style="width:300px" v-model="profileInfo.bili"></a-input>
      </a-form-model-item>
      <a-form-model-item label="github">
        <a-input style="width:400px" v-model="profileInfo.github"></a-input>
      </a-form-model-item>
      <a-form-model-item label="头像">
        <a-upload name="file" :multiple="false" :action="uploadURL" :headers="headers" @change="uploadAvatar" listType="picture" :defaultFileList="fileList">
          <a-button>
            <a-icon type="upload"/>点击上传文件
          </a-button>
          <br/>
          <template v-if="profileInfo.avatar">
            <img :src="profileInfo.avatar" style="width:200px;height:150px">
          </template>
        </a-upload>
      </a-form-model-item>
      <a-form-model-item>
        <a-button type="danger" @click="updateProfile">更新</a-button>
        <a-button type="primary">取消</a-button>
      </a-form-model-item>
    </a-form-model>
  </a-card>
</div>
</template>

<script>
import {Url} from '@/plugin/http'
export default {
    data(){
      return {
        profileInfo:{
          id:1,
          name:'',
          desc:'',
          qq:'',
          bili:'',
          github:'',
          avatar:'',
        },
        uploadURL : Url + '/upload',
        headers:{},
        fileList:[],
      }
    },
    created(){
      this.getProfileInfo()
      this.headers = {Authorization : 'Bearer ' + window.sessionStorage.getItem('token')}
    },
  methods:{
    async getProfileInfo(){
      const {data:res} = await this.$http.get("profile/1")
      if(res.code !== 200)return this.$message.error(res.message)
      this.profileInfo = res.data
    },
    uploadAvatar(value){
      if(value.file.status === 'done'){
        this.$message.success("图片上传成功")
        const imgURL = value.file.response.url
        this.profileInfo.avatar = imgURL
      }else if(value.file.status === 'error'){
        this.$message.error("图片上传失败")
      }
    },
    async updateProfile(){
      const {data : res} = await this.$http.put(`profile/1`,this.profileInfo)
      if(res.status !== 200)return this.$message.error(res.message)
      this.$message.success("个人信息成功")
    }
  }
}
</script>
