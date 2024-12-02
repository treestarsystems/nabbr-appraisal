<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/auth';
import { userState } from '../types/auth';
import { SwalToastError } from '../helpers/sweetalert';
onMounted(async () => {
  const authStore = useAuthStore();
  const swal: any = inject('$swal');
  // Check for non-null user
  await authStore.protectView;
  // Check for unprivileged user
  const isAuthorized = authStore.checkUserPrivilegeLevel(['APPRAISER']);
  if (!isAuthorized) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastError(swal, 'User Unauthorized ');
  }
});
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper"></div>
  <!-- Page wrapper ends -->
</template>
