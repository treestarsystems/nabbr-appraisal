import { createRouter, createWebHistory } from 'vue-router'
import AuthRegister from './views/AuthRegister.vue'
import AuthLogin from './views/AuthLogin.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/register', component: AuthRegister },
    { path: '/login', component: AuthLogin },
  ],
})

export default router
