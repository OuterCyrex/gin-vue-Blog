import axios from "axios";
import Vue from "vue";

let Url = 'http://localhost:3000/api/v1'

axios.defaults.baseURL = Url

axios.interceptors.request.use((config) => {
    config.headers.Authorization = 'Bearer ' + window.sessionStorage.getItem('token')
    return config
},(error)=>{
    return Promise.reject(error)
})
Vue.prototype.$http = axios

export {Url}
