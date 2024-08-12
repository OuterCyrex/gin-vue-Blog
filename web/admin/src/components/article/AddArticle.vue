<template>
  <div>
    <a-card>
      <h3>{{id?"编辑用户":"新增用户"}}</h3>
      <a-form-model :model="articleInfo" ref="articleInfoRef" :rules="articleInfoRules" :hideRequiredMark="true">
        <a-form-model-item label="文章标题" prop="title">
          <a-input v-model="articleInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章分类" prop="cid">
          <a-select v-model="articleInfo.cid" allowClear placeholder="请选择标签" style="width: 200px">
            <a-select-option v-for="item in categoryList" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="文章描述" prop="desc">
          <a-input v-model="articleInfo.desc"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章缩略图" prop="img">
          <a-upload name="file" :multiple="false" :action="uploadURL" :headers="headers" @change="uploadChange" listType="picture" :defaultFileList="fileList">
            <a-button>
              <a-icon type="upload"/>点击上传文件
            </a-button>
            <br/>
            <template v-if="id">
              <img :src="articleInfo.img" style="width:200px;height:150px">
            </template>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item label="文章内容" prop="content">
          <a-input type="textarea" v-model="articleInfo.content"></a-input>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="danger" @click="articleOk(articleInfo.id)"> {{ articleInfo.id ? '更新':'提交' }} </a-button>
          <a-button type="primary" @click="addArticleCancel">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import {Url} from '@/plugin/http'
export default {
  props:['id'],
  data(){
    return{
      uploadURL:Url + '/upload',
      headers:{},
      fileList:[],
      articleInfo:{
        id:0,
        title:'',
        cid:undefined,
        desc:'',
        content:'',
        img:'',
      },
      categoryList:[],
      articleInfoRules:{
        title:[{required:true,message:"标题不能为空",trigger:"blur"}],
        cid:[{required:true,message:"请选择标签",trigger:"change"}],
        desc:[{required:true,message:"文章描述不能为空",trigger:"blur"},{max:120,message:"字数限制在120以内",trigger:"change"}],
        img:[{required:true,message:"请选择封面图片",trigger:"blur"}],
        content:[{required:true,message:"内容不能为空",trigger:"blur"}],
      }
    }
  },
  created() {
    this.getCategoryList()
    this.headers = {Authorization : 'Bearer ' + window.sessionStorage.getItem('token')}
    if(this.id !== undefined){
      this.getArticleInfo(this.id)
    }
  },
  methods:{
    async getCategoryList(){
      const { data : res } = await this.$http.get("categories",{
        params:{
          name:'',
          pagesize: 1000,
          pagenum: 1,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.categoryList = res.data
    },
    async getArticleInfo(id){
      const {data:res} = await this.$http.get("article/"+id)
      if(res.status !== 200)return this.$message.error(res.message)
      this.articleInfo = res.data
      this.articleInfo.id = res.data.ID
    },
    uploadChange(value){
      if(value.file.status !== 'uploading'){

      }
      if(value.file.status === 'done'){
        this.$message.success("图片上传成功")
        const imgURL = value.file.response.url
        this.articleInfo.img = imgURL
      }else if(value.file.status === 'error'){
        this.$message.error("图片上传失败")
      }
    },
    articleOk(id){
      this.$refs.articleInfoRef.validate(async (valid) => {
        if(!valid)return this.$message.error("请填写完整内容")
        if(id === 0){
          const {data : res} = await this.$http.post('article/add',this.articleInfo)
          if(res.status !== 200)return this.$message.error(res.message)
          this.$message.success("添加成功")
          this.$router.push('/admin/articleList')
        }else{
          const {data : res} = await this.$http.put(`article/${id}`,this.articleInfo)
          if(res.status !== 200)return this.$message.error(res.message)
          this.$message.success("添加成功")
          this.$router.push('/admin/articleList')
        }
      })
    },
    addArticleCancel(){
      this.$refs.articleInfoRef.resetFields()
    }
  },
}
</script>
