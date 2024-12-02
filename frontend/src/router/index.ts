import { createRouter, createWebHistory } from 'vue-router';
import AuthRegisterView from '../views/AuthRegisterView.vue';
import AuthLoginView from '../views/AuthLoginView.vue';
import AppraisalView from '../views/AppraisalView.vue';
import DashboardView from '../views/Dashboard.vue';
import UserProfileView from '../views/UserProfileView.vue';
import NotFoundView from '../views/NotFoundView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/register', component: AuthRegisterView },
    { path: '/login', component: AuthLoginView },
    { path: '/appraisal', component: AppraisalView },
    { path: '/dashboard', component: DashboardView },
    { path: '/user/:userId', component: UserProfileView },
    { path: '/:catchAll(.*)', component: NotFoundView },
  ],
});

export default router;
