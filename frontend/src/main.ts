import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import './assets/css/bootstrap-icons.min.css'
import './assets/css/main.min.css'
import './assets/css/style.css'
import './assets/css/OverlayScrollbars.min.css'

createApp(App).use(router).mount('#app')
