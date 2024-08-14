import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/http'
import './plugin/antui'
import './assets/css/style.css'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

Vue.use(mavonEditor)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
