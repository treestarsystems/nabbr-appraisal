<script setup lang="ts">
import { onMounted } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/auth';
import { userState } from '@/types/auth';
onMounted(async () => {
  const authStore = useAuthStore();
  // Check for non-null user
  authStore.protectView();
  // Ensure user has the correct userPrivilegeLevel to access this page.
  const userState: userState | null = authStore.getState;
  console.log(userState?.email);
  // authStore.checkUserPrivilegeLevel(userState?.userPrivilegeLevel ?? '', ['ADMIN', 'APPRAISER']);
});
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper"></div>
  <!-- Page wrapper ends -->
</template>
