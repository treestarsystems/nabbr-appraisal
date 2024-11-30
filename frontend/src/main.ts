import { createApp } from 'vue'
import App from './App.vue'
import router from './router/'
import store from './stores'
import VueSweetalert2 from 'vue-sweetalert2'
import 'sweetalert2/dist/sweetalert2.min.css'
import './assets/css/bootstrap-icons.min.css'
import './assets/css/main.min.css'
import './assets/css/style.css'
import './assets/css/OverlayScrollbars.min.css'

createApp(App).use(router).use(store).use(VueSweetalert2).mount('#app')
