<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/authStore';
import { SwalToastErrorHelper } from '../helpers/sweetalertHelper';
import { loadThirdPartyJSHelper } from '../helpers/utilsHelper';
import { thirdPartyJSFilePathsBase } from '../helpers/thirdPartyFIlesListHelper';
import NavHeaderComponent from '../components/NavHeaderComponent.vue';
import NavSideBarComponent from '../components/NavSideBarComponent.vue';
import FooterComponent from '../components/FooterComponent.vue';
import UserProfileBodyComponent from '../components/UserProfileBodyComponent.vue';

const authStore = useAuthStore();
const swal: any = inject('$swal');
const breadCrumbTitle = `User Profile: ${authStore.getState?.firstName} ${authStore.getState?.lastName}`;
onMounted(async () => {
  // Check for unprivileged user
  const currentRouteUserId = router.currentRoute.value.params.userId as string;
  const isAuthorizedUserId = authStore.checkUserIdAuthorized(currentRouteUserId);
  if (!isAuthorizedUserId) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastErrorHelper(swal, 'User Unauthorized ');
  }
  await loadThirdPartyJSHelper(thirdPartyJSFilePathsBase);
});
</script>

<template>
  <!-- Page wrapper start -->
  <div class="page-wrapper">
    <!-- Main container start -->
    <div class="main-container">
      <NavSideBarComponent />
      <!-- App container starts -->
      <div class="app-container">
        <NavHeaderComponent :breadCrumbPageTitleCurrent="breadCrumbTitle" />
        <UserProfileBodyComponent />
        <FooterComponent />
      </div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <div id="thirdPartyScripts"></div>
</template>
