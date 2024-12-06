<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/authStore';
import { SwalToastErrorHelper } from '../helpers/sweetalertHelper';
import { loadThirdPartyJSHelper } from '../helpers/utilsHelper';
import { thirdPartyJSFilePathsBase } from '../helpers/thirdPartyFIlesListHelper';
import NavSideBarComponent from '../components/NavSideBarComponent.vue';
import NavHeaderComponent from '../components/NavHeaderComponent.vue';
import AppraisalBodyComponent from '../components/AppraisalBodyComponent.vue';
import FooterComponent from '../components/FooterComponent.vue';

const authStore = useAuthStore();
onMounted(async () => {
  const swal: any = inject('$swal');
  // Check for unprivileged user
  const authorizedPrivilegeLevels: string[] = ['ADMIN', 'APPRAISER'];
  const isAuthorized = authStore.checkUserPrivilegeLevel(authorizedPrivilegeLevels);
  if (!isAuthorized) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastErrorHelper(swal, 'User Unauthorized ');
  }
  await loadThirdPartyJSHelper(thirdPartyJSFilePathsBase);
});
</script>

<template>
  <!-- Page wrapper start -->
  <div class="page-wrapper pinned">
    <!-- Main container start -->
    <div class="main-container">
      <NavSideBarComponent />
      <!-- App container starts -->
      <div class="app-container">
        <NavHeaderComponent breadCrumbPageTitleCurrent="Appraisal" />
        <AppraisalBodyComponent />
        <FooterComponent />
      </div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <div id="thirdPartyScripts"></div>
</template>
