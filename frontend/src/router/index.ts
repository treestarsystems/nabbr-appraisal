import { createRouter, createWebHistory } from 'vue-router';
import AuthRegisterView from '../views/AuthRegisterView.vue';
import AuthLoginView from '../views/AuthLoginView.vue';
import AppraisalView from '../views/AppraisalView.vue';
import DashboardView from '../views/Dashboard.vue';
import UserProfileView from '../views/UserProfileView.vue';
import NotFoundView from '../views/NotFoundView.vue';
import { useAuthStore } from '../stores/auth';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/register', component: AuthRegisterView },
    { path: '/login', component: AuthLoginView },
    {
      path: '/appraisal',
      component: AppraisalView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/dashboard',
      component: DashboardView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/user/:userId',
      component: UserProfileView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/appraisal/:chartId',
      component: AppraisalView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/:catchAll(.*)',
      component: NotFoundView,
      meta: {
        requiresAuth: true,
      },
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  if (to.meta.requiresAuth) {
    const authStore = useAuthStore();
    const authenticated = authStore.getState;
    if (authenticated !== null) {
      // User is authenticated, proceed to the route if token is not expired or invalid.
      await authStore.checkTokenExpired();
      next();
    } else {
      // User is not authenticated, redirect to login
      next('/login');
    }
  } else {
    // Non-protected route, allow access
    next();
  }
});

export default router;
