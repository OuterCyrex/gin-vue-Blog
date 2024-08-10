import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/LoginWeb.vue'
import Admin from '../views/AdminWeb.vue'

import Index from '../components/admin/Index.vue'
import AddArticle from '../components/article/AddArticle.vue'
import ArticleList from '../components/article/ArticleList.vue'
import CategoryList from '../components/category/CategoryList.vue'
import UserList from '../components/user/UserList.vue'

Vue.use(VueRouter)

const routes = [
    {
        path:'/login',
        name:'login',
        component:Login
    },
    {
        path:'/admin',
        name:'admin',
        component: Admin,
        children:[
            {path:'index',component: Index},
            {path:'addarticle',component: AddArticle},
            {path:'articlelist',component: ArticleList},
            {path:'categorylist',component: CategoryList},
            {path:'userlist',component: UserList},
        ]
    }
]


const router = new VueRouter({
    routes,
})

router.beforeEach((to, from, next) => {
    const token = window.sessionStorage.getItem('token');
    if(to.path === '/login')return next()
    if(!token && to.path === '/admin'){
        next('/login')
    }else{
        next()
    }
})

export default router
