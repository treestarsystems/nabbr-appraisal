<script setup lang="ts">
import { onMounted, inject } from 'vue';
import router from '../router';
import { useAuthStore } from '../stores/auth';
import { SwalToastError } from '../helpers/sweetalert';
import { loadThirdPartyJS } from '../helpers/script';
import { thirdPartyJSFilePathsBase, thirdPartyJSFilePathsDate } from '../helpers/thirdPartyFIlesList';
import NavSideBar from '../components/NavSideBar.vue';
import NavHeader from '../components/NavHeader.vue';
import AppraisalBody from '../components/AppraisalBody.vue';
import Footer from '../components/Footer.vue';

const authStore = useAuthStore();
onMounted(async () => {
  const swal: any = inject('$swal');
  // Check for unprivileged user
  const authorizedPrivilegeLevels: string[] = ['ADMIN', 'APPRAISER'];
  const isAuthorized = authStore.checkUserPrivilegeLevel(authorizedPrivilegeLevels);
  if (!isAuthorized) {
    router.push(`/user/${authStore.getState?.userId}`);
    SwalToastError(swal, 'User Unauthorized ');
  }
  await loadThirdPartyJS([...thirdPartyJSFilePathsBase]);
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
        <NavHeader breadCrumbCurrentPageTitle="Appraisal" />
        <AppraisalBody />
        <Footer />
      </div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <div id="thirdPartyScripts"></div>
</template>
