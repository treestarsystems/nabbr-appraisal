<script setup lang="ts">
import { onMounted } from 'vue';
import { loadThirdPartyJSHelper } from '../helpers/utilsHelper';
import { thirdPartyJSFilePathsBase } from '../helpers/thirdPartyFIlesListHelper';
import NavSideBarComponent from '../components/NavSideBarComponent.vue';
import NavHeaderComponent from '../components/NavHeaderComponent.vue';
import AppraisalBodyComponent from '../components/AppraisalBodyComponent.vue';
import FooterComponent from '../components/FooterComponent.vue';

onMounted(async () => {
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
        <NavHeaderComponent />
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
