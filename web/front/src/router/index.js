import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ArticleList from '../components/ArticleList.vue'
import DetailsWeb from "@/components/DetailsWeb.vue";
import IndexWeb from "@/components/IndexWeb.vue";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    children:[
      {path:"/",component:IndexWeb,meta:{title:'欢迎访问Outer Blog'}},
      {path:'/articles',component:ArticleList,meta:{title:'文章列表 - Outer Blog'}},
      {path:'/articles/:id',component:ArticleList,meta:{title:'文章列表 - Outer Blog'},props:true},
      {path:'/detail/:id',component: DetailsWeb,meta:{title:'文章详情 - Outer Blog'},props:true},
    ]
  },
]

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
})

document.title = '欢迎访问Outer Blog'

router.beforeEach((to, from, next) => {
  if(to.meta.title){
    document.title = to.meta.title
  }
  next()
})

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

export default router
