import Vue from 'vue'
import {
    Button,
    FormModel,
    Input,
    Icon,
    message,
    Layout,
    Menu,
    Card,
    Table,
    Row,
    Col,
} from 'ant-design-vue'

message.config({
    duration:2,
    maxCount:4,
})


Vue.prototype.$message = message

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Card)
Vue.use(Table)
Vue.use(Row)
Vue.use(Col)
