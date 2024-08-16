<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed"  style="user-select: none">
  <div class="logo">
    <span>{{ collapsed?"Blog":"Outer Blog"}}</span>
  </div>

  <a-menu theme="dark" mode="inline" @click="gotoPage">
    <a-menu-item key="index"><a-icon type="dashboard"></a-icon><span>仪表盘</span></a-menu-item>
    <a-sub-menu>
      <span slot="title"><a-icon type="form"></a-icon><span>文章管理</span></span>
      <a-menu-item key="addarticle"><a-icon type="plus"></a-icon><span>添加文章</span></a-menu-item>
      <a-menu-item key="articlelist"><a-icon type="book"></a-icon><span>文章列表</span></a-menu-item>
    </a-sub-menu>
    <a-menu-item key="categorylist"><a-icon type="tag"></a-icon><span>标签列表</span></a-menu-item>
    <a-menu-item key="userlist"><a-icon type="user"></a-icon><span>用户列表</span></a-menu-item>
    <a-menu-item key="profile"><a-icon type="setting"></a-icon><span>个人设置</span></a-menu-item>
    <a-menu-item key="friendlink"><a-icon type="link"></a-icon><span>友链设置</span></a-menu-item>
  </a-menu>
  </a-layout-sider>
</template>
<style scoped>

.logo{
  height:32px;
  margin:16px;
  background-color: #f4f4f4;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size:18px;
}
</style>
<script>
export default {
  data(){
    return{
      collapsed:false,
    }
  },
  methods:{
    gotoPage(item){
      this.$router.push('/'+item.key).catch((err) => err)
    }
  }
}
</script>
