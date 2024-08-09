import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/LoginWeb.vue'
import Admin from '../views/AdminWeb.vue'

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
        component: Admin
    }
]


const router = new VueRouter({
    routes,
})

export default router
