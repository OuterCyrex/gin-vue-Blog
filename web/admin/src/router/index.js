import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/LoginWeb.vue'
import Admin from '../views/AdminWeb.vue'

import Index from '../components/admin/Index.vue'
import AddArticle from '../components/article/AddArticle.vue'
import ArticleList from '../components/article/ArticleList.vue'
import CategoryList from '../components/category/CategoryList.vue'
import UserList from '../components/user/UserList.vue'
import Profile from '../components/user/Profile.vue'
import FriendLink from "../components/user/FriendLink.vue";
import CommentView from "@/components/user/CommentView.vue";

Vue.use(VueRouter)

const routes = [
    {
        path:'/login',
        name:'login',
        component:Login
    },
    {
        path:'/',
        name:'admin',
        component: Admin,
        children:[
            {path:'index',component: Index},
            {path:'addarticle',component: AddArticle},
            {path:'addarticle/:id',component: AddArticle,props:true},
            {path:'articlelist',component: ArticleList},
            {path:'categorylist',component: CategoryList},
            {path:'userlist',component: UserList},
            {path:'profile',component:Profile},
            {path:'friendlink',component:FriendLink},
            {path:'comment',component:CommentView}
        ]
    }
]


const router = new VueRouter({
    mode: 'history',
    routes,
})

router.beforeEach((to, from, next) => {
    const token = window.sessionStorage.getItem('token');
    if(to.path === '/login')return next()
    if(!token){
        next('/login')
    }else{
        next()
    }
})

export default router
