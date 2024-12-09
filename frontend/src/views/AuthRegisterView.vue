<script setup lang="ts">
import LogoAuthComponent from '../components/LogoAuthComponent.vue';
import { reactive, ref, inject } from 'vue';
import { RouterLink } from 'vue-router';
import { useAuthStore } from '../stores/authStore';
import router from '../router';
import { formValidationAreAllFieldsFilledHelper } from '../helpers/utilsHelper';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { FormDataRegistration, FormDataRegistrationSubmit } from '../types/formTypes';
import { SwalToastErrorHelper, SwalToastSuccessHelper } from '../helpers/sweetalertHelper';

const authStore = useAuthStore();
const swal: any = inject('$swal');
let wasValidated = ref('');

const formRegister: FormDataRegistration = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  password: '',
  confirmPassword: '',
  userPrivilegeLevel: 'PETOWNER',
  registrationKey: '',
});

async function submitRegistrationForm() {
  try {
    const userRegistrationFormData: FormDataRegistrationSubmit = {
      firstName: formRegister.firstName,
      lastName: formRegister.lastName,
      email: formRegister.email.toLowerCase(),
      phone: formRegister.phone,
      password: formRegister.password,
      userPrivilegeLevel: formRegister.userPrivilegeLevel,
      registrationKey: formRegister.registrationKey,
    };

    if (!formValidationAreAllFieldsFilledHelper(formRegister, wasValidated)) {
      return;
    }

    // Final check for password fields
    if (formRegister.password !== formRegister.confirmPassword) return;

    // Show a visual cue that the form has been submitted.
    wasValidated.value = 'was-validated';
    const apiResponse: ResponseObjectDefaultInterface = await authStore.register(userRegistrationFormData);
    if (apiResponse.httpStatus > 299) {
      throw apiResponse;
    }

    SwalToastSuccessHelper(swal, 'User Created!');
    await router.push('/login');
  } catch (err: any) {
    SwalToastErrorHelper(swal, err);
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
          <LogoAuthComponent />
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
                <select
                  v-model="formRegister.userPrivilegeLevel"
                  class="form-select"
                  id="userPrivilegeLevel"
                  aria-label="Default select example"
                >
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
