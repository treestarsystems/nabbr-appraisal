<script setup lang="ts"></script>

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
                    id="tab-twoA"
                    data-bs-toggle="tab"
                    href="#twoA"
                    role="tab"
                    aria-controls="twoA"
                    aria-selected="false"
                    ><i class="bi bi-info-circle me-2"></i>Notifications</a
                  >
                </li>
                <li class="nav-item" role="presentation">
                  <a
                    class="nav-link"
                    id="tab-threeA"
                    data-bs-toggle="tab"
                    href="#threeA"
                    role="tab"
                    aria-controls="threeA"
                    aria-selected="false"
                    ><i class="bi bi-credit-card-2-front me-2"></i>Credit Cards</a
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
                    ><i class="bi bi-eye-slash me-2"></i>Reset Password</a
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
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="fullName" class="form-label">Full Name</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input type="text" class="form-control" id="fullName" placeholder="Full name" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="yourEmail" class="form-label">Email</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-envelope"></i>
                                  </span>
                                  <input
                                    type="email"
                                    class="form-control"
                                    id="yourEmail"
                                    placeholder="Email ID"
                                    value="info@email.com"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="contactNumber" class="form-label">Contact</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-phone"></i>
                                  </span>
                                  <input type="text" class="form-control" id="contactNumber" placeholder="Contact" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="birthDay" class="form-label">DOB</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-calendar4"></i>
                                  </span>
                                  <input type="text" class="form-control" id="birthDay" placeholder="" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-12">
                              <!-- Form field start -->
                              <div class="m-0">
                                <label class="form-label" for="abt">About</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-filter-circle"></i>
                                  </span>
                                  <textarea class="form-control" id="abt" rows="4"></textarea>
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                        </div>
                      </div>
                    </div>
                  </div>
                  <!-- Row ends -->
                </div>
                <div class="tab-pane fade" id="twoA" role="tabpanel">
                  <!-- Row starts -->
                  <div class="row gx-5 align-items-center">
                    <div class="col-sm-4 col-12">
                      <div class="p-3">
                        <img src="/assets/images/notifications.svg" alt="Notifications" class="img-fluid" />
                      </div>
                    </div>
                    <div class="col-sm-4 col-12">
                      <!-- List group start -->
                      <ul class="list-group mb-3">
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          Desktop Notifications
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchOne" checked />
                          </div>
                        </li>
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          Email Notifications
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchTwo" checked />
                          </div>
                        </li>
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          Chat Notifications
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchThree" checked />
                          </div>
                        </li>
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          New Message
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchFour" />
                          </div>
                        </li>
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          New Follower
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchFive" />
                          </div>
                        </li>
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          New Review
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchSix" />
                          </div>
                        </li>
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                          New Order
                          <div class="form-check form-switch m-0">
                            <input class="form-check-input" type="checkbox" role="switch" id="switchSeven" />
                          </div>
                        </li>
                      </ul>
                      <!-- List group end -->
                    </div>
                  </div>
                  <!-- Row ends -->
                </div>
                <div class="tab-pane fade" id="threeA" role="tabpanel">
                  <!-- Row starts -->
                  <div class="row gx-3">
                    <div class="col-12">
                      <div class="table-outer mb-3">
                        <div class="table-responsive">
                          <table class="table truncate align-middle m-0">
                            <thead>
                              <tr>
                                <th>Bank Name</th>
                                <th>Card Number</th>
                                <th>Card type</th>
                                <th>Expiry Date</th>
                                <th>Credit Balance</th>
                                <th>Actions</th>
                              </tr>
                            </thead>
                            <tbody>
                              <tr>
                                <td>Discover It Cash Back</td>
                                <td>
                                  <span class="badge bg-primary text-white">0000 0000 0000 0000</span>
                                </td>
                                <td>Visa</td>
                                <td>20/10/2026</td>
                                <td>$100000</td>
                                <td>
                                  <div class="form-check form-switch m-0">
                                    <input
                                      class="form-check-input"
                                      type="checkbox"
                                      role="switch"
                                      id="cardActive"
                                      checked
                                    />
                                  </div>
                                </td>
                              </tr>
                              <tr>
                                <td>Capital One Quicksilver Card</td>
                                <td>
                                  <span class="badge bg-success text-white">0000 0000 0000 0000</span>
                                </td>
                                <td>Master</td>
                                <td>02/24/2028</td>
                                <td>$150000</td>
                                <td>
                                  <div class="form-check form-switch m-0">
                                    <input
                                      class="form-check-input"
                                      type="checkbox"
                                      role="switch"
                                      id="cardActive2"
                                      checked
                                    />
                                  </div>
                                </td>
                              </tr>
                              <tr>
                                <td>Bank of America Rewards Card</td>
                                <td>
                                  <span class="badge bg-warning text-white">0000 0000 0000 0000</span>
                                </td>
                                <td>Visa</td>
                                <td>05/09/2028</td>
                                <td>$50000</td>
                                <td>
                                  <div class="form-check form-switch m-0">
                                    <input
                                      class="form-check-input"
                                      type="checkbox"
                                      role="switch"
                                      id="cardActive3"
                                      checked
                                    />
                                  </div>
                                </td>
                              </tr>
                              <tr>
                                <td>Chase Freedom Unlimited</td>
                                <td>
                                  <span class="badge bg-danger text-white">0000 0000 0000 0000</span>
                                </td>
                                <td>Discover</td>
                                <td>08/20/2027</td>
                                <td>$100000</td>
                                <td>
                                  <div class="form-check form-switch m-0">
                                    <input
                                      class="form-check-input"
                                      type="checkbox"
                                      role="switch"
                                      id="cardActive4"
                                      checked
                                    />
                                  </div>
                                </td>
                              </tr>
                              <tr>
                                <td>Blue Cash Everyday Card</td>
                                <td>
                                  <span class="badge bg-secondary text-white">0000 0000 0000 0000</span>
                                </td>
                                <td>Visa</td>
                                <td>24/10/2029</td>
                                <td>$90000</td>
                                <td>
                                  <div class="form-check form-switch m-0">
                                    <input class="form-check-input" type="checkbox" role="switch" id="cardActive5" />
                                  </div>
                                </td>
                              </tr>
                              <tr>
                                <td>Wells Fargo Active Card</td>
                                <td>
                                  <span class="badge bg-info text-white">0000 0000 0000 0000</span>
                                </td>
                                <td>Visa</td>
                                <td>25/05/2030</td>
                                <td>$75000</td>
                                <td>
                                  <div class="form-check form-switch m-0">
                                    <input class="form-check-input" type="checkbox" role="switch" id="cardActive6" />
                                  </div>
                                </td>
                              </tr>
                            </tbody>
                          </table>
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
                                id="currentPwd"
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
                                placeholder="Your password must be 8-20 characters long."
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
                                placeholder="Confirm new password"
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
                </div>
              </div>
              <!-- Tab content end -->
            </div>
            <!-- Custom tabs end -->

            <!-- Buttons start -->
            <div class="d-flex gap-2 justify-content-end">
              <button type="button" class="btn btn-outline-secondary">Cancel</button>
              <button type="button" class="btn btn-primary">Update</button>
            </div>
            <!-- Buttons end -->
          </div>
        </div>
      </div>
    </div>
    <!-- Row end -->
  </div>
  <!-- App body ends -->
</template>
