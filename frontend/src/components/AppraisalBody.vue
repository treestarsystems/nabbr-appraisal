<script setup lang="ts">
import { RouterLink } from 'vue-router';
import AppraisalChart from './AppraisalChart.vue';
import { generateCalendarDateString } from '../helpers/script';
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
              <ul class="nav nav-tabs tab-nav-ul" id="customTab2" role="tablist">
                <li class="nav-item" role="presentation">
                  <a
                    class="nav-link active"
                    id="tab-appraisalDetails"
                    data-bs-toggle="tab"
                    href="#appraisalDetails"
                    role="tab"
                    aria-controls="appraisalDetails"
                    aria-selected="true"
                    ><i class="bi bi-person me-2"></i> Details</a
                  >
                </li>
                <li class="nav-item" role="presentation">
                  <a
                    class="nav-link"
                    id="tab-appraisalScore"
                    data-bs-toggle="tab"
                    href="#appraisalScore"
                    role="tab"
                    aria-controls="appraisalScore"
                    aria-selected="false"
                    tabindex="-1"
                    ><i class="bi bi-ui-radios-grid me-2"></i> Score Card</a
                  >
                </li>
                <li class="nav-item" role="presentation">
                  <a
                    class="nav-link"
                    id="tab-appraisalResults"
                    data-bs-toggle="tab"
                    href="#appraisalResults"
                    role="tab"
                    aria-controls="appraisalResults"
                    aria-selected="false"
                    tabindex="-1"
                    ><i class="bi bi-ui-radios me-2"></i> Results</a
                  >
                </li>
                <li class="nav-item px-5">
                  <!-- Form field start -->
                  <div class="input-group">
                    <span class="input-group-text">
                      <i class="bi bi-trophy"></i>
                    </span>
                    <input
                      type="text"
                      class="form-control"
                      id="appraisalTotalScore"
                      disabled="disabled"
                      placeholder="Total Score"
                    />
                  </div>
                  <!-- Form field end -->
                </li>
              </ul>
              <!-- Nav tabs end -->

              <!-- Tab content start -->
              <div class="tab-content">
                <div
                  class="tab-pane fade active show"
                  id="appraisalDetails"
                  role="tabpanel"
                  aria-labelledby="tab-appraisalDetails"
                >
                  <!-- Row starts -->
                  <div class="row gx-3">
                    <div class="col-sm-12 col-12">
                      <div class="card border mb-3">
                        <div class="card-body">
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-6 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="fullName" class="form-label">Member Name:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input type="text" class="form-control" id="fullName" placeholder="Full name" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-6 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="memberNumber" class="form-label">Member #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input type="text" class="form-control" id="memberNumber" placeholder="Member #" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="dogName" class="form-label">Dog Name:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input type="text" class="form-control" id="dogName" placeholder="Dog Name" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="dogRegistrationNumber" class="form-label">Dog Registration #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    type="email"
                                    class="form-control"
                                    id="dogRegistrationNumber"
                                    placeholder="Dog Registration #"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="microChipNumber" class="form-label">Microchip #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    type="text"
                                    class="form-control"
                                    id="microChipNumber"
                                    placeholder="Microchip #"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="dnaNumber" class="form-label">DNA #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input type="text" class="form-control" id="dnaNumber" placeholder="DNA #" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="dogColor" class="form-label">Color:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-palette"></i>
                                  </span>
                                  <input type="text" class="form-control" id="dogColor" placeholder="Brown" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="dogMask" class="form-label">Mask:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-palette"></i>
                                  </span>
                                  <input type="text" class="form-control" id="dogMask" placeholder="Black" />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="contactNumber" class="form-label">Age:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <select class="form-select" id="age" aria-label="Default select example">
                                    <option value="0">&lt;1yr</option>
                                    <option selected="" value="1">1yr</option>
                                    <option value="2">2yr</option>
                                    <option value="3">4yr</option>
                                    <option value="4">4yr</option>
                                    <option value="5">5yr</option>
                                    <option value="6">6yr</option>
                                    <option value="7">7yr</option>
                                    <option value="8">8yr</option>
                                    <option value="9">9yr</option>
                                    <option value="10">10yr</option>
                                    <option value="11">11yr</option>
                                    <option value="12">12yr</option>
                                    <option value="13">13yr</option>
                                    <option value="14">14yr</option>
                                    <option value="15">15yr</option>
                                    <option value="16">16yr</option>
                                    <option value="17">17yr</option>
                                    <option value="18">18yr</option>
                                    <option value="19">19yr</option>
                                    <option value="20">20yr</option>
                                  </select>
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="birthDay" class="form-label">Sex:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-gender-ambiguous"></i>
                                  </span>
                                  <select class="form-select" id="a5" aria-label="Default select example">
                                    <option selected="" value="male">Male</option>
                                    <option value="female">Female</option>
                                  </select>
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
                <div class="tab-pane fade" id="appraisalScore" role="tabpanel" aria-labelledby="tab-appraisalScore">
                  <AppraisalChart />
                </div>
                <div class="tab-pane fade" id="appraisalResults" role="tabpanel" aria-labelledby="tab-appraisalResults">
                  <!-- Row starts -->
                  <div class="row gx-3">
                    <div class="col-sm-12 col-12">
                      <div class="card border mb-3">
                        <div class="card-body">
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-6 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="seniorAppraiserName" class="form-label">Senior Appraiser:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input
                                    type="text"
                                    class="form-control"
                                    id="seniorAppraiserName"
                                    placeholder="Senior Appraiser Name"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-6 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="seniorAppraiserNumber" class="form-label">Senior Appraiser #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    type="text"
                                    class="form-control"
                                    id="seniorAppraiserNumber"
                                    placeholder="Senior Appraiser #"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-6 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="appraiserName" class="form-label">Appraiser:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <input
                                    type="text"
                                    class="form-control"
                                    id="appraiserName"
                                    placeholder="Appraiser Name"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-6 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="appraiserNumber" class="form-label">Appraiser #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    type="text"
                                    class="form-control"
                                    id="appraiserNumber"
                                    placeholder="Appraiser #"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <div class="col-sm-12 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="additionalCommentary" class="form-label">Additonal Commentary:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <textarea class="form-control" id="additionalCommentary" rows="6"></textarea>
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                          </div>
                          <!-- Row ends -->
                          <!-- Row starts -->
                          <div class="row gx-3">
                            <!-- Form placeholder fields start -->
                            <div class="col-sm-2 col-4">
                              <div class="m-0">
                                <label class="form-label" for="abc3">Date:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-calendar4"></i>
                                  </span>
                                  <input
                                    type="date"
                                    id="appraisalDate"
                                    class="form-control datepicker"
                                    name="birthday"
                                    :value="generateCalendarDateString()"
                                  />
                                </div>
                              </div>
                            </div>
                            <div class="col-sm-6"></div>
                            <!-- Form placeholder fields start -->
                            <div class="col-sm-2 col-4">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="appraisalPlace" class="form-label">Place:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-trophy-fill"></i>
                                  </span>
                                  <select class="form-select" id="appraisalPlace" aria-label="Default select example">
                                    <option selected="" value="1">1st</option>
                                    <option value="2">2nd</option>
                                    <option value="3">3rd</option>
                                    <option value="4">4th</option>
                                    <option value="5">5th</option>
                                    <option value="6">6th</option>
                                    <option value="7">7th</option>
                                    <option value="8">8th</option>
                                    <option value="9">9th</option>
                                    <option value="10">10th</option>
                                  </select>
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-2 col-4">
                              <label class="form-label"></label>
                              <!-- Buttons start -->
                              <div class="d-flex gap-2 justify-content-end">
                                <button type="button" class="btn btn-primary">Submit Appraisal</button>
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

  <!-- App body end -->
</template>
