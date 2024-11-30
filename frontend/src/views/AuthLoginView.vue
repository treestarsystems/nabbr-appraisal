<script setup lang="ts">
import LogoAuth from '../components/LogoAuth.vue';
import { reactive, ref, inject } from 'vue';
import { RouterLink } from 'vue-router';
import axios from 'axios';
import router from '../router';
import { apiResponseDefault, formDataLogin } from '../types/auth';
const swal: any = inject('$swal');
let wasValidated = ref('');

const formLogin = reactive({
  email: '',
  password: '',
});

async function submitLoginForm() {
  try {
    const userLoginFormData: formDataLogin = {
      email: formLogin.email,
      password: formLogin.password,
    };

    for (const key in formLogin) {
      const value = formLogin[key as keyof typeof formLogin];
      if (!value) {
        console.log(key);
        wasValidated.value = 'was-validated';
        return;
      }
    }

    const apiResponse: apiResponseDefault = (await axios.post('/backend/api/v1/auth/login', userLoginFormData)).data;
    if (apiResponse.httpStatus > 299) {
      throw apiResponse;
    }

    console.log(apiResponse);

    await swal
      .mixin({
        toast: true,
        position: 'top-right',
        iconColor: 'white',
        customClass: {
          popup: 'colored-toast',
        },
        showConfirmButton: false,
        timer: 1500,
      })
      .fire({
        icon: 'success',
        title: `Login Successful!`,
      });

    const { userPrivilegeLevel, userId } = apiResponse.payload[0];
    if (userPrivilegeLevel === 'PETOWNER') {
      await router.push(`/user/${userId}`);
    } else {
      await router.push('/dashboard');
    }
  } catch (err: any) {
    swal
      .mixin({
        toast: true,
        position: 'top-right',
        iconColor: 'white',
        customClass: {
          popup: 'colored-toast',
        },
        showConfirmButton: false,
        timer: 5000,
        timerProgressBar: true,
      })
      .fire({
        icon: 'error',
        title: `${err.response.data.message || err.data.message}`,
      });
    console.log(err);
  }
}
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper">
    <!-- Auth container starts -->
    <div class="auth-container">
      <div class="d-flex justify-content-center">
        <!-- Form starts -->
        <form @submit.prevent="submitLoginForm" class="need-validation" :class="wasValidated" novalidate>
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
                <input
                  v-model="formLogin.email"
                  type="email"
                  id="email"
                  class="form-control"
                  placeholder="Enter your email"
                  required
                />
              </div>
            </div>

            <div class="mb-4">
              <label class="form-label" for="password">Password <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-lock"></i>
                </span>
                <input
                  v-model="formLogin.password"
                  type="password"
                  id="password"
                  class="form-control"
                  placeholder="Enter password"
                  minlength="8"
                  maxlength="64"
                  required
                />
              </div>
            </div>

            <!-- <div class="d-flex justify-content-end mb-3">
              <RouterLink to="/forgot-password" class="text-decoration-underline">Forgot password?</RouterLink>
            </div> -->

            <div class="d-grid gap-2">
              <button type="submit" class="btn btn-primary">Login</button>
              <RouterLink to="/register" class="btn btn-outline-secondary">Not registered? Signup</RouterLink>
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
