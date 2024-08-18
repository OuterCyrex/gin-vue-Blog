<template>
  <div>
    <div class="mainTitle">标签列表</div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search v-model="queryParam.name" placeholder="搜索标签" enter-button allowClear @search="getCategoryList"/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addCategoryVisible = true">新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="name" :columns="columns" :pagination="pagination" :dataSource="categoryList" @change="handleTableChange" bordered style="margin-top:40px;">
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" @click="editCategory(data.id)" style="margin:0 10px;">编辑</a-button>
            <a-button type="danger" icon="delete" @click="deleteCategory(data.id)" style="margin:0 10px;">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!--新增标签-->
    <a-modal
        title="新增用户"
        :visible="addCategoryVisible"
        @ok="addCategoryOk"
        @cancel="addCategoryCancel"
        destroyOnClose >
      <a-form-model :model="newCategory" :rules="newCategoryRule" ref="addCategoryRef">
        <a-form-model-item has-feedback label="标签名" prop="name">
          <a-input v-model="newCategory.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!--编辑用户-->
    <a-modal
        title="编辑标签名称"
        :visible="editCategoryVisible"
        @ok="editCategoryOk"
        @cancel="editCategoryCancel"
        destroyOnClose >
      <a-form-model :model="categoryInfo" :rules="categoryRule" ref="editCategoryRef">
        <a-form-model-item has-feedback label="标签名" prop="username">
          <a-input v-model="categoryInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title:'ID',
    dataIndex:'id',
    width:'10%',
    key:'id',
    align:"center",
  },
  {
    title:"标签名",
    dataIndex:"name",
    width:"20%",
    key:"name",
    align:"center",
  },
  {
    title:"操作",
    width:"30%",
    key:"action",
    align:"center",
    scopedSlots:{customRender:'action'},
  },
]

export default {
  data(){
    return{
      pagination:{
        pageSizeOptions: ['5','10','20'],
        pageSize:5,
        total:0,
        showSizeChanger:true,
        showTotal:(total)=>`共${total}条`,
      },
      categoryInfo:{
        id:0,
        name:'',
      },
      newCategory:{
        name:'',
      },
      categoryList:[],
      columns,
      queryParam:{
        name:'',
        pagesize:5,
        pagenum:1,
      },
      visible:false,
      addCategoryVisible:false,
      editCategoryVisible:false,
      categoryRule:{
        name:[{validator:(rule,value,callback)=>{
            if(this.newCategory.name === ""){
              callback(new Error("请输入分类名"))
            }
            callback()
          },trigger:"blur"}],
      },
      newCategoryRule:{
        name:[{validator:(rule,value,callback)=>{
            if(this.newCategory.name === ""){
              callback(new Error("请输入分类名"))
            }
            callback()
          },trigger:"blur"}],
      },
    }
  },
  created(){
    this.getCategoryList()
  },

  methods:{
    async getCategoryList(){
      const { data : res } = await this.$http.get("categories", {
        params: {
          name: this.queryParam.name,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.categoryList = res.data
      this.pagination.total = res.total
    },

    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getCategoryList()
    },

    async deleteCategory(id){
      this.$confirm({
        title:"确定要删除吗？",
        content:"一旦删除将不可恢复，真的会消失很久！",
        onOk:async()=>{
          const res = await this.$http.delete(`category/${id}`)
          if(res.data.status !== 200) {
            this.$message.error("存在文章在当前标签下")
          }else{
            this.$message.success("操作成功")
          }
          this.getCategoryList()
        },
        onCancel:()=>{
          this.$message.info("取消成功")
        },
      })
    },

    addCategoryOk(){
      this.$refs.addCategoryRef.validate(async (valid)=>{
        if(!valid)return this.$message.error("请填写正确信息")
        const { data:res } = await this.$http.post('category/add', {
          name: this.newCategory.name,
        })
        if(res.status !== 200) return this.$message.error(res.message)
        this.$message.success("创建标签成功")
        this.addCategoryCancel()
        this.getCategoryList()
      })
    },

    addCategoryCancel(){
      this.$refs.addCategoryRef.resetFields()
      this.addCategoryVisible = false
    },

    async editCategory(id){
      this.editCategoryVisible = true
      const {data:res} = await this.$http.get(`category/${id}`)
      this.categoryInfo = res.data
      this.categoryInfo.id = id
    },

    editCategoryOk(){
      this.$refs.editCategoryRef.validate(async (valid)=>{
        if(!valid)return this.$message.error("请填写正确信息")
        const { data:res } = await this.$http.put('category/'+ this.categoryInfo.id, {
          id: this.categoryInfo.id,
          name: this.categoryInfo.name,
        })
        if(res.status !== 200) return this.$message.error(res.message)
        this.$message.success("更新标签名成功")
        this.editCategoryCancel()
        this.getCategoryList()
      })
    },
    editCategoryCancel(){
      this.$refs.editCategoryRef.resetFields()
      this.editCategoryVisible = false
    },
  }
}
</script>
<style scoped>
.mainTitle{
  font-size: calc(20px + 2vw);
  text-align: center;
  align-content: center;
  background-color: white;
  font-family:"方正粗黑宋简体",sans-serif;
  user-select: none;
}
</style>
