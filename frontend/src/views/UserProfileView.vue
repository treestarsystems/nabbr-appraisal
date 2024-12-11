<script setup lang="ts">
import { onMounted } from 'vue';
import { useAuthStore } from '../stores/authStore';
import { loadThirdPartyJSHelper } from '../helpers/utilsHelper';
import { thirdPartyJSFilePathsBase } from '../helpers/thirdPartyFIlesListHelper';
import NavHeaderComponent from '../components/NavHeaderComponent.vue';
import NavSideBarComponent from '../components/NavSideBarComponent.vue';
import FooterComponent from '../components/FooterComponent.vue';
import UserProfileBodyComponent from '../components/UserProfileBodyComponent.vue';

const authStore = useAuthStore();
const breadCrumbTitle = `User Profile: ${authStore.getState?.firstName} ${authStore.getState?.lastName}`;
onMounted(async () => {
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
