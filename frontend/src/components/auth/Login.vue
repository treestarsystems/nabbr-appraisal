<script setup lang="ts">
import { ref } from 'vue'
import LogoAuth from '../logo/LogoAuth.vue'
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper">
    <!-- Auth container starts -->
    <div class="auth-container">
      <div class="d-flex justify-content-center">
        <!-- Form starts -->
        <form action="/login">
          <LogoAuth />
          <!-- Authbox starts -->
          <div class="auth-box">
            <h4 class="mb-4">Welcome back,</h4>

            <div class="mb-3">
              <label class="form-label" for="email">Email <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-envelope"></i>
                </span>
                <input type="email" id="email" class="form-control" placeholder="Enter your email" />
              </div>
            </div>

            <div class="mb-2">
              <label class="form-label" for="password">Password <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-lock"></i>
                </span>
                <input type="password" id="password" class="form-control" placeholder="Enter password" />
                <button class="btn btn-outline-secondary" type="button">
                  <i class="bi bi-eye"></i>
                </button>
              </div>
            </div>

            <div class="d-flex justify-content-end mb-3">
              <a href="forgot-password.html" class="text-decoration-underline">Forgot password?</a>
            </div>

            <div class="d-grid gap-2">
              <button type="submit" class="btn btn-primary">Login</button>
              <a href="signup.html" class="btn btn-outline-secondary">Not registered? Signup</a>
            </div>
          </div>
          <!-- Authbox ends -->
        </form>
        <!-- Form ends -->
      </div>
    </div>
    <!-- Auth container ends -->
  </div>
  <!-- Page wrapper ends -->
</template>
