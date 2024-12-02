<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/auth';
import { SwalToastError } from '../helpers/sweetalert';

const authStore = useAuthStore();
const swal: any = inject('$swal');
onMounted(async () => {
  // Check for non-null user
  // await authStore.protectView;
  // Check for unprivileged user
  const currentRouteUserId = router.currentRoute.value.params.userId as string;
  const isAuthorizedUserId = authStore.checkUserIdAuthorized(currentRouteUserId);
  if (!isAuthorizedUserId) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastError(swal, 'User Unauthorized ');
  }
});
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper">
    <!-- Auth container starts -->
    <div class="auth-container">
      <div class="row justify-content-center">
        <div class="col-9">
          <div class="d-flex align-items-center justify-content-center vh-100">
            <div class="text-center text-white">
              <!-- <h1 class="error-title mb-3">User Profile: {{ userId }}</h1> -->
              <h1 class="error-title mb-3">User Profile:</h1>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Auth container ends -->
  </div>
  <!-- Page wrapper ends -->
</template>
