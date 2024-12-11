<script setup lang="ts">
import LogoAuth from '../components/LogoAuthComponent.vue';
import { reactive, ref, inject, toRaw } from 'vue';
import { RouterLink } from 'vue-router';
import { useAuthStore } from '../stores/authStore';
import router from '../router';
import { formValidationAreAllFieldsFilledHelper } from '../helpers/utilsHelper';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { FormDataUserBase } from '../types/formTypes';
import { SwalToastErrorHelper, SwalToastSuccessHelper } from '../helpers/sweetalertHelper';

const authStore = useAuthStore();
const swal: any = inject('$swal');
const user = toRaw(authStore.getState);
const userProfileLink = `/user/${user?.userId}`;
let wasValidated = ref('');

const formLogin: FormDataUserBase = reactive({
  email: '',
  password: '',
});

async function submitLoginForm() {
  try {
    const userLoginFormData: FormDataUserBase = {
      email: formLogin.email.toLowerCase(),
      password: formLogin.password,
    };

    if (!formValidationAreAllFieldsFilledHelper(formLogin, wasValidated)) {
      return;
    }

    // Show a visual cue that the form has been submitted.
    wasValidated.value = 'was-validated';
    const apiResponse: ResponseObjectDefaultInterface = await authStore.login(userLoginFormData);
    if (apiResponse.httpStatus > 299) {
      throw apiResponse.message;
    }

    SwalToastSuccessHelper(swal, 'Login Successful!');

    // Our user state should be defined at this point.
    const user = authStore.getState;
    if (user?.userPrivilegeLevel === 'PETOWNER') {
      await router.push(userProfileLink);
    } else {
      await router.push('/dashboard');
    }
  } catch (err: any) {
    SwalToastErrorHelper(swal, err);
  }
}
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper">
    <!-- Auth container starts -->
    <div v-if="authStore.getState === null" class="auth-container">
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
    <div v-else class="auth-container">
      <div class="row justify-content-center">
        <div class="col-9">
          <div class="d-flex align-items-center justify-content-center vh-100">
            <div class="text-center text-white">
              <h1 class="error-title mb-3">Please Logout First</h1>
              <RouterLink :to="userProfileLink" class="btn btn-warning px-5 shadow-lg py-3 fs-5"
                >Go back to User Profile</RouterLink
              >
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Auth container ends -->
  </div>
  <!-- Page wrapper ends -->
</template>
