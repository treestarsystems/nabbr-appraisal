<script setup lang="ts">
import { toRaw } from 'vue';
import { RouterLink } from 'vue-router';
import { useAuthStore } from '../stores/authStore';
const authStore = useAuthStore();
const user = toRaw(authStore.getState);
let link401 = `/user/${user?.userId}`;
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper">
    <!-- Auth container starts -->
    <div class="auth-container">
      <div class="row justify-content-center">
        <div class="col-9">
          <div class="d-flex align-items-center justify-content-center vh-100">
            <div class="text-center text-white">
              <h1 class="error-title mb-3">401</h1>
              <h3 class="mb-5 fw-lighter lh-2">
                We're sorry but it looks like you are <br />not authorized to view that page.
              </h3>
              <RouterLink :to="link401" class="btn btn-warning px-5 shadow-lg py-3 fs-5">User Profile</RouterLink>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Auth container ends -->
  </div>
  <!-- Page wrapper ends -->
</template>
