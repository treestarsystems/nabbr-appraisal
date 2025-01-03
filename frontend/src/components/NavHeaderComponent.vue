<script setup lang="ts">
import { toRaw, ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../stores/authStore';
import { RouterLink } from 'vue-router';
import {
  hideNavSidebarHelper,
  breadCrumbPageLinkPreviousHelper,
  breadCrumbPageLinkCurrentHelper,
} from '../helpers/utilsHelper';
import { UserState } from '../types/authTypes';

const route = useRoute();
const hideNavSidebarRef = ref(false);
const authStore = useAuthStore();
const user = toRaw(authStore.getState as UserState);
const userProfileLink = `/user/${user?.userId}`;
const breadCrumbPageLinkPrevious = () => {
  return breadCrumbPageLinkPreviousHelper(route, authStore);
};
const breadCrumbPageLinkCurrent = () => {
  return breadCrumbPageLinkCurrentHelper(route, authStore);
};

onMounted(async () => {
  hideNavSidebarHelper(route, hideNavSidebarRef);
});
</script>

<template>
  <!-- App header starts -->
  <div class="app-header">
    <!-- Toggle buttons start -->
    <div :class="['d-flex', { 'd-none d-xs-block d-xxl-block': hideNavSidebarRef }]">
      <button class="btn btn-outline-info btn-sm me-3 toggle-sidebar" id="toggle-sidebar">
        <i class="bi bi-list fs-5"></i>
      </button>
      <button class="btn btn-outline-info btn-sm me-3 pin-sidebar" id="pin-sidebar">
        <i class="bi bi-list fs-5"></i>
      </button>
    </div>
    <!-- Toggle buttons end -->

    <!-- App brand start -->
    <div class="app-brand-sm">
      <RouterLink :to="userProfileLink" class="d-lg-none d-md-block">
        <img src="/assets/images/dog.svg" class="logo" alt="" />
      </RouterLink>
    </div>
    <!-- App brand end -->

    <!-- App header actions start -->
    <div class="header-actions gap-3">
      <!-- Header settings starts -->
      <div class="dropdown">
        <a
          id="userSettings"
          class="dropdown-toggle d-flex py-2 align-items-center text-decoration-none"
          href="#!"
          role="button"
          data-bs-toggle="dropdown"
          aria-expanded="false"
        >
          <div class="icon-box md bg-info text-white rounded-5">
            {{ user?.firstName[0].toUpperCase() }}{{ user?.lastName[0].toUpperCase() }}
          </div>
        </a>
        <div class="dropdown-menu dropdown-menu-end">
          <RouterLink class="dropdown-item d-flex align-items-center" :to="userProfileLink"
            ><i class="bi bi-person fs-4 me-2"></i>Profile</RouterLink
          >
          <a href="#" class="dropdown-item d-flex align-items-center" @click="authStore.logout">
            <i class="bi bi-escape fs-4 me-2"></i>Logout
          </a>
        </div>
      </div>
      <!-- Header settings starts -->
    </div>
    <!-- App header actions end -->
  </div>
  <!-- App header ends -->

  <!-- App hero header starts -->
  <div class="app-hero-header d-flex align-items-center">
    <!-- Breadcrumb start -->
    <ol class="breadcrumb d-lg-flex">
      <li class="breadcrumb-item">
        <i class="bi bi-house lh-1"></i>
        <RouterLink :to="breadCrumbPageLinkPrevious().link" class="text-decoration-none">{{
          breadCrumbPageLinkPrevious().linkName || 'Profile'
        }}</RouterLink>
      </li>
      <li class="breadcrumb-item" aria-current="page">
        {{ breadCrumbPageLinkCurrent() }}
      </li>
    </ol>
    <!-- Breadcrumb end -->
  </div>
  <!-- App Hero header ends -->
</template>
