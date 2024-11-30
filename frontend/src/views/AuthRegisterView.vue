<script setup lang="ts">
import LogoAuth from '../components/LogoAuth.vue'

const submitForm = async () => {
  console.log('Submit Form')
}
</script>

<template>
  <!-- Page wrapper starts -->
  <div class="page-wrapper">
    <!-- Auth container starts -->
    <div class="auth-container">
      <div class="d-flex justify-content-center">
        <!-- Form starts -->
        <form @submit.prevent="submitForm">
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
                          <input type="text" id="firstName" class="form-control" placeholder="First name" />
                        </div>
                      </td>
                      <td class="pe-0">
                        <label class="form-label" for="lastName">Last Name <span class="text-danger">*</span></label>
                        <div class="input-group">
                          <span class="input-group-text">
                            <i class="bi bi-person"></i>
                          </span>
                          <input type="text" id="firstName" class="form-control" placeholder="First name" />
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
                <input type="email" id="email" class="form-control" placeholder="Enter your email" />
              </div>
            </div>
            <div class="mb-3">
              <label class="form-label" for="email">Phone <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text">
                  <i class="bi bi-phone"></i>
                </span>
                <input type="email" id="email" class="form-control" placeholder="Enter your phone number" />
              </div>
            </div>
            <div class="mb-3">
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
                <input type="password" id="confirmPassword" class="form-control" placeholder="Confirm password" />
                <button class="btn btn-outline-secondary" type="button">
                  <i class="bi bi-eye"></i>
                </button>
              </div>
              <div class="form-text">Please confirm your password.</div>
            </div>

            <div class="d-grid gap-2">
              <button type="submit" class="btn btn-primary">Signup</button>
              <a href="/login" class="btn btn-outline-secondary">Already have an account? Login</a>
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
