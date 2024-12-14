<script setup lang="ts">
// import { onMounted } from 'vue';
// import NavSideBarComponent from '../components/NavSideBarComponent.vue';
// import NavHeaderComponent from '../components/NavHeaderComponent.vue';
// import AppraisalBodyComponent from '../components/AppraisalBodyComponent.vue';
// import FooterComponent from '../components/FooterComponent.vue';

// onMounted(async () => {
//   await loadThirdPartyJSHelper(thirdPartyJSFilePathsBase);
// });
</script>

<template>
  <!-- Page wrapper start -->
  <div class="page-wrapper pinned">
    <!-- Main container start -->
    <div class="main-container">
      <!-- App container starts -->
      <div class="app-container"></div>
      <!-- App container ends -->
    </div>
    <!-- Main container end -->
  </div>
  <!-- Page wrapper end -->
  <div id="thirdPartyScripts"></div>
</template>

<style lang="css" scoped>
*,
*::before,
*::after {
  box-sizing: content-box;
}
</style>
