import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import VueSweetalert2 from 'vue-sweetalert2';
import AuthStorePlugin from './plugins/authStore';
import router from './router/';
import App from './App.vue';
import 'sweetalert2/dist/sweetalert2.min.css';
import './assets/css/bootstrap-icons.min.css';
import './assets/css/main.min.css';
import './assets/css/style.css';
import './assets/css/OverlayScrollbars.min.css';

// Create Pinia instance
const pinia = createPinia();

// Use persisted state with Pinia so our store data will persist even after page refresh
pinia.use(piniaPluginPersistedstate);

// createApp(App).use(router).use(userAuthStore).use(VueSweetalert2).mount('#app');

createApp(App).use(AuthStorePlugin, { pinia }).use(pinia).use(router).use(VueSweetalert2).mount('#app');
