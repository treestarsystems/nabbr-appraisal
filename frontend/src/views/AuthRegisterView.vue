<script setup lang="ts">
import LogoAuth from '../components/LogoAuth.vue';
import { reactive, ref, inject } from 'vue';
import { RouterLink } from 'vue-router';
import axios from 'axios';
import router from '../router';
import { apiResponseDefault, formDataRegistration } from '../types/auth';
const swal: any = inject('$swal');
let wasValidated = ref('');

const formRegister = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',
  userPrivilegeLevel: '',
  registrationKey: '',
});

async function submitRegistrationForm() {
  try {
    const userRegistrationForm: formDataRegistration = {
      firstName: formRegister.firstName,
      lastName: formRegister.lastName,
      email: formRegister.email,
      phone: formRegister.phone,
      password: formRegister.password,
      userPrivilegeLevel: formRegister.userPrivilegeLevel,
      registrationKey: formRegister.registrationKey,
    };

    for (const key in formRegister) {
      const value = formRegister[key as keyof typeof formRegister];
      if (!value) {
        console.log(key);
        wasValidated.value = 'was-validated';
        return;
      }
    }

    // Final check for password fields
    if (formRegister.password !== formRegister.confirmPassword) return;

    // const apiResponse = await axios.post(
    //   `${import.meta.env.BACKEND_API_BASE_URL}/api/v1/auth/signup`,
    //   userRegistrationForm,
    // );

    const apiResponse: apiResponseDefault = (await axios.post('/backend/api/v1/auth/signup', userRegistrationForm))
      .data;
    if (apiResponse.httpStatus > 299) {
      throw apiResponse;
    }

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
        timerProgressBar: true,
      })
      .fire({
        icon: 'success',
        title: `User Created!`,
      });
    await router.push('/login');
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
        <form @submit.prevent="submitRegistrationForm" class="need-validation" :class="wasValidated" novalidate>
          <LogoAuth />
          <!-- Authbox starts -->
          <div class="auth-box">
            <h4 class="mb-4">Signup</h4>
            <div class="mb-0">
              <div class="table-responsive">
                <table class="table mb-2">
                  <tbody>
                    <tr>
                      <td class="ps-0">
                        <label class="form-label" for="firstName">First Name <span class="text-danger">*</span></label>
                        <div class="input-group">
                          <span class="input-group-text">
                            <i class="bi bi-person"></i>
                          </span>
                          <input
                            v-model="formRegister.firstName"
                            ref="firstName"
                            type="text"
                            id="firstName"
                            class="form-control"
                            placeholder="First name"
                            required
                          />
                        </div>
                      </td>
                      <td class="pe-0">
                        <label class="form-label" for="lastName">Last Name <span class="text-danger">*</span></label>
                        <div class="input-group">
                          <span class="input-group-text">
                            <i class="bi bi-person"></i>
                          </span>
                          <input
                            v-model="formRegister.lastName"
                            type="text"
                            id="lastName"
                            class="form-control"
                            placeholder="Last name"
                            required
                          />
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
            <div class="mb-3">
              <label class="form-label" for="email">Email <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-envelope"></i>
                </span>
                <input
                  v-model="formRegister.email"
                  type="email"
                  id="email"
                  class="form-control"
                  placeholder="Enter your email"
                  required
                />
              </div>
            </div>
            <div class="mb-3">
              <label class="form-label" for="phone">Phone <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-phone"></i>
                </span>
                <input
                  v-model="formRegister.phone"
                  type="tel"
                  id="phone"
                  class="form-control"
                  placeholder="Enter your phone number"
                  pattern="[0-9]{10}"
                  required
                />
              </div>
            </div>
            <div class="mb-3">
              <label class="form-label" for="password">Password <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-lock"></i>
                </span>
                <input
                  v-model="formRegister.password"
                  type="password"
                  id="password"
                  class="form-control"
                  placeholder="Enter password"
                  minlength="8"
                  maxlength="64"
                  required
                />
              </div>
              <div class="form-text">Your password must be 8-64 characters long.</div>
            </div>
            <div class="mb-3">
              <label class="form-label" for="confirmPassword"
                >Confirm Password <span class="text-danger">*</span></label
              >
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-lock"></i>
                </span>
                <input
                  v-model="formRegister.confirmPassword"
                  type="password"
                  id="confirmPassword"
                  class="form-control"
                  placeholder="Confirm password"
                  minlength="8"
                  maxlength="64"
                  required
                />
              </div>
              <div v-if="formRegister.password !== formRegister.confirmPassword" class="form-text text-danger">
                Please ensure passwords match!
              </div>
              <div v-else class="form-text">Please confirm your password.</div>
            </div>

            <div class="mb-3">
              <label class="form-label" for="userPrivilegeLevel">User Role<span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-safe"></i>
                </span>
                <select class="form-select" id="userPrivilegeLevel" aria-label="Default select example">
                  <option selected="true" value="APPRAISER">Appraiser</option>
                  <option value="PETOWNER">Pet Owner</option>
                  <option value="ADMIN">Admin</option>
                </select>
              </div>
            </div>
            <div class="mb-5">
              <label class="form-label" for="registrationKey"
                >Registration Key <span class="text-danger">*</span></label
              >
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-key"></i>
                </span>
                <input
                  v-model="formRegister.registrationKey"
                  type="password"
                  id="registrationKey"
                  class="form-control"
                  placeholder="Enter registration key"
                  required
                />
              </div>
            </div>
            <div class="d-grid gap-2">
              <button type="submit" class="btn btn-primary">Signup</button>
              <RouterLink to="/login" class="btn btn-outline-secondary">Already have an account? Login</RouterLink>
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
