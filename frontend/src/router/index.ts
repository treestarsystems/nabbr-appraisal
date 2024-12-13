import { createRouter, createWebHistory } from 'vue-router';
import AuthRegisterView from '../views/AuthRegisterView.vue';
import AuthLoginView from '../views/AuthLoginView.vue';
import AppraisalView from '../views/AppraisalView.vue';
import DashboardView from '../views/DashboardView.vue';
import UserProfileView from '../views/UserProfileView.vue';
import NotFoundView from '../views/NotFoundView.vue';
import { useAuthStore } from '../stores/authStore';
import UnAuthorizedView from '../views/UnAuthorizedView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/register', component: AuthRegisterView },
    { path: '/login', component: AuthLoginView },
    {
      path: '/dashboard',
      component: DashboardView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/appraisal',
      component: AppraisalView,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/appraisal/:appraisalId',
      component: AppraisalView,
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
      path: '/unauthorized',
      component: UnAuthorizedView,
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

/**
 * Check if route:
 * 1. Requires authentication
 * 2. If user is authenticated
 * 3. If user is authorized
 * 4. then proceed to the route else redirect.
 */
router.beforeEach(async (to, from, next) => {
  try {
    if (to.meta.requiresAuth) {
      const authStore = useAuthStore();
      const authenticated = authStore.getState;
      if (authenticated !== null) {
        // User is authenticated, proceed to the route if token is not expired or invalid.
        await authStore.checkTokenExpired();
        await authStore.checkUserPrivilegeLevelAuthorizedThenRedirect(to?.path);
        if (to?.path.includes('user')) await authStore.checkUserIdAuthorized(to?.path);
        next();
      } else {
        // User is not authenticated, redirect to login
        next('/login');
      }
    } else {
      // Non-protected route, allow access
      next();
    }
  } catch (error) {
    console.error(error);
  }
});

export default router;
