<script setup lang="ts">
import { onMounted } from 'vue';
import { loadThirdPartyJSHelper } from '../helpers/utilsHelper';
import { thirdPartyJSFilePathsBase } from '../helpers/thirdPartyFIlesListHelper';
import DashboardBodyComponent from '../components/DashboardBodyComponent.vue';
import NavSideBarComponent from '../components/NavSideBarComponent.vue';
import NavHeaderComponent from '../components/NavHeaderComponent.vue';
import FooterComponent from '../components/FooterComponent.vue';

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
        <NavHeaderComponent />
        <DashboardBodyComponent />
        <FooterComponent />
      </div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <div id="thirdPartyScripts"></div>
</template>
