import './assets/main.css'
/* main.js 或 main.ts文件中 */
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'  // 引入全局 CSS 样式

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
// 挂载 Element-Plus
app.use(ElementPlus)

app.mount('#app')
