<template>
  <div>
    <h3>用户列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search v-model="queryParam.username" placeholder="搜索用户" enter-button allowClear @search="getUserList"/>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addUserVisible = true">新增</a-button>
        </a-col>
      </a-row>
      <a-table rowKey="username" :columns="columns" :pagination="pagination" :dataSource="userList" @change="handleTableChange" bordered style="margin-top:40px;">
        <span slot="role" slot-scope="role">{{ role === 1 ? "管理员" : "用户"}}</span>
        <template slot="action" slot-scope="data">
        <div class="actionSlot">
          <a-button type="primary" icon="edit" @click="editUser(data.ID)" style="margin-right:20px;">编辑</a-button>
          <a-button type="danger" icon="delete" @click="deleteUser(data.ID)">删除</a-button>
        </div>
        </template>
      </a-table>
    </a-card>

    <!--新增用户-->
    <a-modal
      title="新增用户"
      :visible="addUserVisible"
      @ok="addUserOk"
      @cancel="addUserCancel"
    destroyOnClose >
      <a-form-model :model="newUser" :rules="newUserRule" ref="addUserRef">
        <a-form-model-item has-feedback label="用户名" prop="username">
          <a-input v-model="newUser.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="newUser.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpsw">
          <a-input-password v-model="newUser.checkpsw"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!--编辑用户-->
    <a-modal
        title="编辑用户"
        :visible="editUserVisible"
        @ok="editUserOk"
        @cancel="editUserCancel"
        destroyOnClose >
      <a-form-model :model="userInfo" :rules="userRule" ref="editUserRef">
        <a-form-model-item has-feedback label="用户名" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="权限" prop="role">
          <a-select placeholder="请选择权限" @change="adminChange">
            <a-select-option key="1" value="1">管理员</a-select-option>
            <a-select-option key="2" value="2">用户</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title:'ID',
    dataIndex:'ID',
    width:'10%',
    key:'id',
    align:"center",
  },
  {
    title:"用户名",
    dataIndex:"username",
    width:"20%",
    key:"username",
    align:"center",
  },
  {
    title:"权限",
    dataIndex:"role",
    width:"20%",
    key:"role",
    align:"center",
    scopedSlots:{customRender:'role'},
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
      userInfo:{
        id:0,
        username:'',
        role:2,
      },
      newUser:{
        id:0,
        username:'',
        password:'',
        checkpsw:'',
      },
      userList:[],
      columns,
      queryParam:{
        username:'',
        pagesize:5,
        pagenum:1,
      },
      visible:false,
      addUserVisible:false,
      editUserVisible:false,
      userRule:{
        username:[{validator:(rule,value,callback)=>{
          if(this.userInfo.username === ""){
            callback(new Error("请输入用户名"))
          }
            if([...this.userInfo.username].length < 4 || [...this.userInfo.username].length > 12){
              callback(new Error(" 用户名长度应在 4 到 12 之间"))
            }else{
              callback()
            }
          },trigger:"blur"}],
      },
      newUserRule:{
        username:[{validator:(rule,value,callback)=>{
            if(this.newUser.username === ""){
              callback(new Error("请输入用户名"))
            }
            if([...this.newUser.username].length < 4 || [...this.newUser.username].length > 12){
              callback(new Error(" 用户名长度应在 4 到 12 之间"))
            }else{
              callback()
            }
          },trigger:"blur"}],
        password:[{validator:(rule,value,callback)=>{
            if(this.newUser.password === ""){
              callback(new Error('请输入密码'))
            }
            if([...this.newUser.password].length < 6 || [...this.newUser.password].length > 20){
              callback(new Error("密码长度应在 6 到 20 之间"))
            }else{
              callback()
            }
          },trigger:'blur',
        }],
        checkpsw:[{validator:(rule,value,callback)=>{
            if(this.newUser.checkpsw === ""){
              callback(new Error('请输入密码'))
            }
            if(this.newUser.checkpsw !== this.newUser.password){
              callback(new Error('两次输入密码不一致'))
            }else{
              callback()
            }
          },trigger:'blur'}]
      },
    }
  },
  created(){
    this.getUserList()
  },
  methods:{
    async getUserList(){
      const { data : res } = await this.$http.get("users", {
        params: {
          username:this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        }
      })
      if(res.status !== 200) return this.$message.error(res.message)
      this.userList = res.data
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
      this.getUserList()
    },

    async deleteUser(id){
      this.$confirm({
        title:"确定要删除吗？",
        content:"一旦删除将不可恢复，真的会消失很久！",
        onOk:async()=>{
          const res = await this.$http.delete(`user/${id}`)
          if(res.status !== 200) return this.$message.error(res.message)
          this.getUserList()
          this.$message.success("操作成功")
        },
        onCancel:()=>{
          this.$message.info("取消成功")
        },
      })
    },

    addUserOk(){
      this.$refs.addUserRef.validate(async (valid)=>{
        if(!valid)return this.$message.error("请填写正确信息")
        const { data:res } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role:2,
        })
        if(res.status !== 200) return this.$message.error(res.message)
        this.$message.success("创建用户成功")
        this.addUserCancel()
        this.getUserList()
      })
    },

    addUserCancel(){
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
    },
    adminChange(value){
      this.userInfo.role = Number(value)
    },
    async editUser(id){
      this.editUserVisible = true
      const {data:res} = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserOk(){
      this.$refs.editUserRef.validate(async (valid)=>{
        if(!valid)return this.$message.error("请填写正确信息")
        const { data:res } = await this.$http.put('user/'+ this.userInfo.id, {
          username: this.userInfo.username,
          role: this.userInfo.role,
        })
        if(res.status !== 200) return this.$message.error(res.message)
        this.$message.success("更新用户信息成功")
        this.editUserCancel()
        this.getUserList()
      })
    },
    editUserCancel(){
        this.$refs.editUserRef.resetFields()
        this.editUserVisible = false
    },
  }
}
</script>
