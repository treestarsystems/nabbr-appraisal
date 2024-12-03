<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/auth';
import { SwalToastError } from '../helpers/sweetalert';
import { loadThirdPartyJS } from '../helpers/script';
import { thirdPartyJSFilePathsBase } from '../helpers/thirdPartyFIlesList';
import NavHeader from '../components/NavHeader.vue';
import NavSideBar from '../components/NavSideBar.vue';
import Footer from '../components/Footer.vue';
import UserProfileBody from '../components/UserProfileBody.vue';

const authStore = useAuthStore();
const swal: any = inject('$swal');
const breadCrumbTitle = `User Profile: ${authStore.getState?.firstName} ${authStore.getState?.lastName}`;
onMounted(async () => {
  // Check for unprivileged user
  const currentRouteUserId = router.currentRoute.value.params.userId as string;
  const isAuthorizedUserId = authStore.checkUserIdAuthorized(currentRouteUserId);
  if (!isAuthorizedUserId) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastError(swal, 'User Unauthorized ');
  }
  await loadThirdPartyJS(thirdPartyJSFilePathsBase);
});
</script>

<template>
  <!-- Page wrapper start -->
  <div class="page-wrapper pinned">
    <!-- Main container start -->
    <div class="main-container">
      <NavSideBar />
      <!-- App container starts -->
      <div class="app-container">
        <NavHeader :breadCrumbCurrentPageTitle="breadCrumbTitle" />
        <UserProfileBody />
        <Footer />
      </div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <div id="thirdPartyScripts"></div>
</template>
