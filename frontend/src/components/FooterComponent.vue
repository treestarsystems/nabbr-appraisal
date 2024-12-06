<script setup lang="ts">
import { currentYearHelper } from '../helpers/utilsHelper';
</script>

<template>
  <!-- App footer start -->
  <div class="app-footer">
    <span>© NABBR Appraisal {{ currentYearHelper() }}</span>
  </div>
  <!-- App footer end -->
</template>
