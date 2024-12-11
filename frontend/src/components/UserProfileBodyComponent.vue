<script setup lang="ts">
import { reactive, inject, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { UserState } from '../types/authTypes';
import { FormDataUserProfile } from '../types/formTypes';
import { useAuthStore } from '../stores/authStore';
import { SwalToastErrorHelper } from '../helpers/sweetalertHelper';

const swal: any = inject('$swal');
const authStore = useAuthStore();
const user = authStore.getState as UserState;
const formProfile: FormDataUserProfile = reactive({
  firstName: user?.firstName,
  lastName: user?.lastName,
  phone: user?.phone,
  password: '',
  confirmPassword: '',
  oldPassword: '',
});

function updateUserInfo() {
  try {
    throw 'test';
  } catch (err) {
    SwalToastErrorHelper(swal, err);
  }
}

function updateUserPassword() {
  try {
    throw 'updateUserPassword';
  } catch (err) {
    SwalToastErrorHelper(swal, err);
  }
}
</script>

<template>
  <!-- App body starts -->
  <div class="app-body">
    <!-- Row start -->
    <div class="row gx-3">
      <div class="col-xxl-12">
        <div class="card">
          <div class="card-body">
            <!-- Custom tabs start -->
            <div class="custom-tabs-container">
              <!-- Nav tabs start -->
              <ul class="nav nav-tabs" id="customTab2" role="tablist">
                <li class="nav-item" role="presentation">
                  <a
                    class="nav-link active"
                    id="tab-oneA"
                    data-bs-toggle="tab"
                    href="#oneA"
                    role="tab"
                    aria-controls="oneA"
                    aria-selected="true"
                    ><i class="bi bi-person me-2"></i> Personal Details</a
                  >
                </li>
                <li class="nav-item" role="presentation">
                  <a
                    class="nav-link"
                    id="tab-fourA"
                    data-bs-toggle="tab"
                    href="#fourA"
                    role="tab"
                    aria-controls="fourA"
                    aria-selected="false"
                    ><i class="bi bi-eye-slash me-2"></i>Update Password</a
                  >
                </li>
              </ul>
              <!-- Nav tabs end -->

              <!-- Tab content start -->
              <div class="tab-content">
                <div class="tab-pane fade show active" id="oneA" role="tabpanel">
                  <!-- Row starts -->
                  <div class="row gx-3">
                    <div class="col-sm-12 col-12">
                      <div class="card border mb-3">
                        <div class="card-body">
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-4 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="firstName" class="form-label">First Name</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input
                                    v-model="formProfile.firstName"
                                    type="text"
                                    class="form-control"
                                    id="firstName"
                                    placeholder="First Name"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-4 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="lastName" class="form-label">Last Name</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input
                                    v-model="formProfile.lastName"
                                    type="text"
                                    class="form-control"
                                    id="lastName"
                                    placeholder="Last Name"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-4 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="phone" class="form-label">Phone Number</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    v-model="formProfile.phone"
                                    required
                                    type="tel"
                                    id="phone"
                                    class="form-control"
                                    placeholder="Enter your phone number"
                                    pattern="[0-9]{10}"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-10 col-12"></div>
                            <div class="col-sm-2 col-12 mt-5">
                              <!-- Buttons start -->
                              <div class="d-flex gap-2 justify-content-end">
                                <button @click="updateUserInfo" type="button" class="btn btn-primary">
                                  Update User Details
                                </button>
                              </div>
                              <!-- Buttons end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                        </div>
                      </div>
                    </div>
                  </div>
                  <!-- Row ends -->
                </div>
                <div class="tab-pane fade" id="fourA" role="tabpanel">
                  <!-- Row starts -->
                  <div class="row align-items-end">
                    <div class="col-xl-4 col-sm-6 col-12">
                      <div class="p-3">
                        <img src="/assets/images/login.svg" alt="Contact Us" class="img-fluid" />
                      </div>
                    </div>
                    <div class="col-sm-4 col-12">
                      <div class="card border mb-3">
                        <div class="card-body">
                          <div class="mb-3">
                            <label class="form-label" for="currentPwd"
                              >Current password <span class="text-danger">*</span></label
                            >
                            <div class="input-group">
                              <input
                                type="password"
                                id="password"
                                placeholder="Enter Current password"
                                class="form-control"
                              />
                              <button class="btn btn-outline-secondary" type="button">
                                <i class="bi bi-eye"></i>
                              </button>
                            </div>
                          </div>

                          <div class="mb-3">
                            <label class="form-label" for="newPwd"
                              >New password <span class="text-danger">*</span></label
                            >
                            <div class="input-group">
                              <input
                                type="password"
                                id="newPwd"
                                class="form-control"
                                placeholder="Password must be 8-64 characters."
                              />
                              <button class="btn btn-outline-secondary" type="button">
                                <i class="bi bi-eye"></i>
                              </button>
                            </div>
                          </div>

                          <div class="mb-3">
                            <label class="form-label" for="confNewPwd"
                              >Confirm new password <span class="text-danger">*</span></label
                            >
                            <div class="input-group">
                              <input
                                type="password"
                                id="confNewPwd"
                                placeholder="Password must be 8-64 characters."
                                class="form-control"
                              />
                              <button class="btn btn-outline-secondary" type="button">
                                <i class="bi bi-eye"></i>
                              </button>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <!-- Row ends -->
                  <!-- Row starts -->
                  <div class="row gx-3">
                    <div class="col-sm-10 col-12"></div>
                    <div class="col-sm-2 col-12 mt-1">
                      <!-- Buttons start -->
                      <div class="d-flex gap-2 justify-content-end">
                        <button @click="updateUserPassword" type="button" class="btn btn-primary">
                          Update Password
                        </button>
                      </div>
                      <!-- Buttons end -->
                    </div>
                  </div>
                  <!-- Row ends -->
                </div>
              </div>
              <!-- Tab content end -->
            </div>
            <!-- Custom tabs end -->
          </div>
        </div>
      </div>
    </div>
    <!-- Row end -->
  </div>
  <!-- App body ends -->
</template>
