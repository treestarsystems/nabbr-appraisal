import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import VueSweetalert2 from 'vue-sweetalert2';
import router from './router/';
import App from './App.vue';

// CSS
import 'sweetalert2/dist/sweetalert2.min.css';
import './assets/css/bootstrap-icons.min.css';
import './assets/css/main.min.css';
import './assets/css/style.css';
import './assets/css/OverlayScrollbars.min.css';
import './assets/css/daterange.css';

// Create Pinia instance
const pinia = createPinia();
// Use persisted state with Pinia so our store data will persist even after page refresh
pinia.use(piniaPluginPersistedstate);

createApp(App).use(pinia).use(router).use(VueSweetalert2).mount('#app');
