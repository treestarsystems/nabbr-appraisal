<script setup lang="ts">
import { toRaw } from 'vue';
import { RouterLink, useRoute } from 'vue-router';
import { useAuthStore } from '../stores/auth';
const authStore = useAuthStore();
const user = toRaw(authStore.getState);
const userProfileLink = `/user/${user?.userId}`;

const isActiveLink = (routePath: string) => {
  const route = useRoute();
  return route.path === routePath;
};
</script>

<template>
  <!-- Sidebar wrapper start -->
  <nav id="sidebar" class="sidebar-wrapper">
    <!-- App brand starts -->
    <div class="app-brand px-3 py-2 d-flex align-items-center">
      <a href="index.html">
        <img src="/dog.svg" class="logo" alt="Bootstrap Gallery" />
      </a>
    </div>
    <!-- App brand ends -->

    <!-- Sidebar profile starts -->
    <div class="sidebar-profile">
      <img src="/dog.svg" class="img-3x me-3 rounded-3" alt="Admin Dashboard" />
      <div class="m-0">
        <p class="m-0">Hello &#128075;</p>
        <h6 class="m-0 text-nowrap">{{ user.firstName }} {{ user.lastName }}</h6>
      </div>
    </div>
    <!-- Sidebar profile ends -->

    <!-- Sidebar menu starts -->
    <div class="sidebarMenuScroll">
      <ul class="sidebar-menu side-bar-ul">
        <li
          v-if="user?.userPrivilegeLevel === 'ADMIN'"
          class=""
          :class="[isActiveLink('/dashboard') ? 'active current-page' : '']"
        >
          <RouterLink to="/dashboard">
            <i class="bi bi-justify"></i>
            <span class="menu-text">Dashboard</span>
          </RouterLink>
        </li>
        <li
          v-if="['ADMIN', 'APPRAISER'].includes(user?.userPrivilegeLevel)"
          :class="[isActiveLink('/appraisal') ? 'active current-page' : '']"
        >
          <RouterLink to="/appraisal">
            <i class="bi bi-calculator"></i>
            <span class="menu-text">New Appraisal</span>
          </RouterLink>
        </li>
        <li :class="[isActiveLink(userProfileLink) ? 'active current-page' : '']">
          <RouterLink :to="userProfileLink">
            <i class="bi bi-person-square"></i>
            <span class="menu-text">Profile</span>
          </RouterLink>
        </li>
        <li>
          <a href="#" @click="authStore.logout">
            <i class="bi bi-escape"></i>
            <span class="menu-text">Logout</span>
          </a>
        </li>
      </ul>
      <!-- <div>
        <ul class="sidebar-menu">
          <li>
            <a href="#" @click="authStore.logout">
              <i class="bi bi-escape"></i>
              <span class="menu-text">Logout</span>
            </a>
          </li>
        </ul>
      </div> -->
    </div>
    <!-- Sidebar menu ends -->
  </nav>
  <!-- Sidebar wrapper start -->
</template>
