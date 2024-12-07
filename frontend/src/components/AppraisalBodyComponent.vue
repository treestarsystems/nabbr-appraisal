<script setup lang="ts">
import { ref, inject, toRaw, provide, onMounted } from 'vue';
import { useAuthStore } from '../stores/authStore';
import { getAppraisalChartTemplateHelper } from '../helpers/chartHelper';
import { generateCalendarDateStringHelper } from '../helpers/utilsHelper';
import { SwalToastErrorHelper, SwalToastSuccessHelper } from '../helpers/sweetalertHelper';
import AppraisalBodyChartComponent from './AppraisalBodyChartComponent.vue';
import { Chart } from '../types/chartTypes';

const totalScore = ref(0);
const swal: any = inject('$swal');
const chartData: any = ref('');
provide('chartData', chartData);
provide('totalScore', totalScore);
// let wasValidated = ref('');

onMounted(async () => {
  const authStore = useAuthStore();
  const token = toRaw(authStore.getState)?.token as string;
  const chartTemplate = await getAppraisalChartTemplateHelper(swal, token);
  chartData.value = chartTemplate as any;
});

async function submitChart() {
  try {
    // We have to get these because they are
    const appraisalTotalScore: any = document.getElementById('appraisalTotalScore');
    // const appraisalDate: any = document.getElementById('appraisalDate');
    // const age: any = document.getElementById('age');
    // const sex: any = document.getElementById('sex');
    chartData.value.appraisalInformation.appraisalScore = parseFloat(appraisalTotalScore.value.replace('%', ''));
    // chartData.value.appraisalInformation.date = appraisalDate.value;
    // chartData.value.petInformation.age = parseInt(age.value);
    // chartData.value.petInformation.sex = sex.value;
    // Show a visual cue that the form has been submitted.
    // wasValidated.value = 'was-validated';
    // chartData.value.appraisalInformation.appraisalScore = parseFloat(
    //   chartData.value.appraisalInformation.appraisalScore.replace('%', ''),
    // );

    // API call here
    console.log(toRaw(chartData.value));
    SwalToastSuccessHelper(swal, 'Appraisal Submitted Successfully!');
  } catch (err: any) {
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
                      :value="`${totalScore}%`"
                      @change="chartData.appraisalInformation.appraisalScore = $event.target.value"
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
                  <!-- <div class="row gx-3 was-validated"> -->
                  <div class="row gx-3 was-validated">
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
                                  <input
                                    required
                                    @input="chartData.memberInformation.name = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="fullName"
                                    placeholder="Full Name"
                                  />
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
                                  <input
                                    required
                                    @input="chartData.memberInformation.memberNumber = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="memberNumber"
                                    placeholder="Member #"
                                  />
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
                                  <input
                                    required
                                    @input="chartData.petInformation.name = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="dogName"
                                    placeholder="Dog Name"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="registrationNumber" class="form-label">Dog Registration #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    required
                                    @input="chartData.petInformation.registrationNumber = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="registrationNumber"
                                    placeholder="Dog Registration #"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="microchip" class="form-label">Microchip #:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    required
                                    @input="chartData.petInformation.microchip = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="microchip"
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
                                  <input
                                    required
                                    @input="chartData.petInformation.dnaNumber = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="dnaNumber"
                                    placeholder="DNA #"
                                  />
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
                                  <input
                                    required
                                    @input="chartData.petInformation.color = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="dogColor"
                                    placeholder="Brown"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-3 col-12">
                              <!-- Form field start -->
                              <div class="mb-3">
                                <label for="markings" class="form-label">Mask/Markings:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-palette"></i>
                                  </span>
                                  <input
                                    required
                                    @input="chartData.petInformation.markings = $event.target.value"
                                    type="text"
                                    class="form-control"
                                    id="markings"
                                    placeholder="Black"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-2 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="weight" class="form-label">Weight:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <input
                                    required
                                    @input="chartData.petInformation.weight = parseInt($event.target.value)"
                                    type="number"
                                    name="weight"
                                    class="form-control"
                                    id="weight"
                                    placeholder="Weight"
                                  />
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>

                            <div class="col-sm-2 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="age" class="form-label">Age:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-123"></i>
                                  </span>
                                  <select
                                    @change="chartData.petInformation.age = parseInt($event.target.value)"
                                    class="form-select"
                                    id="age"
                                    aria-label="Default select example"
                                    required
                                  >
                                    <option value="0">&lt;1yr</option>
                                    <option value="1">1yr</option>
                                    <option value="2">2yr</option>
                                    <option value="3">4yr</option>
                                    <option value="4">4yr</option>
                                    <option value="5">5yr</option>
                                    <option value="6">6yr</option>
                                    <option value="7">7yr</option>
                                    <option value="8">8yr</option>
                                    <option value="9">9yr</option>
                                    <option value="10">10yr</option>
                                    <!-- <option disabled selected value>Select Age</option> -->
                                    <option disabled selected value="0">Select Age</option>
                                  </select>
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-2 col-12">
                              <!-- Form field start -->
                              <div class="mb-2">
                                <label for="sex" class="form-label">Sex:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-gender-ambiguous"></i>
                                  </span>
                                  <select
                                    @change="chartData.petInformation.sex = $event.target.value"
                                    class="form-select"
                                    id="sex"
                                    aria-label="Default select example"
                                    required
                                  >
                                    <option selected="" value="male">Male</option>
                                    <option value="female">Female</option>
                                    <option disabled selected value="">Select Sex</option>
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
                  <!-- Appraisal Chart start -->
                  <AppraisalBodyChartComponent />
                  <!-- Appraisal Chart end -->
                </div>
                <div class="tab-pane fade" id="appraisalResults" role="tabpanel" aria-labelledby="tab-appraisalResults">
                  <!-- Row starts -->
                  <div class="row gx-3 was-validated">
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
                                    required
                                    @input="chartData.appraisalInformation.seniorAppraiserName = $event.target.value"
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
                                    required
                                    @input="chartData.appraisalInformation.seniorAppraiserNumber = $event.target.value"
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
                                    required
                                    @input="chartData.appraisalInformation.appraiserName = $event.target.value"
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
                                    required
                                    @input="chartData.appraisalInformation.appraiserNumber = $event.target.value"
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
                                <label for="additionalComments" class="form-label">Additonal Commentary:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-person"></i>
                                  </span>
                                  <textarea
                                    @input="chartData.appraisalInformation.additionalComments = $event.target.value"
                                    class="form-control"
                                    id="additionalComments"
                                    rows="6"
                                  ></textarea>
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
                                <label class="form-label" for="appraisalDate">Date:</label>
                                <div class="input-group">
                                  <span class="input-group-text">
                                    <i class="bi bi-calendar4"></i>
                                  </span>
                                  <input
                                    @change="chartData.appraisalInformation.date = $event.target.value"
                                    type="date"
                                    id="appraisalDate"
                                    class="form-control datepicker"
                                    name="date"
                                    :value="generateCalendarDateStringHelper()"
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
                                  <select
                                    @change="chartData.appraisalInformation.place = parseInt($event.target.value)"
                                    class="form-select"
                                    id="appraisalPlace"
                                    aria-label="Default select example"
                                  >
                                    <option value="1">1st</option>
                                    <option value="2">2nd</option>
                                    <option value="3">3rd</option>
                                    <option value="4">4th</option>
                                    <option value="5">5th</option>
                                    <option value="6">6th</option>
                                    <option value="7">7th</option>
                                    <option value="8">8th</option>
                                    <option value="9">9th</option>
                                    <option value="10">10th</option>
                                    <option selected value="0">No Place</option>
                                  </select>
                                </div>
                              </div>
                              <!-- Form field end -->
                            </div>
                            <div class="col-sm-2 col-4">
                              <label class="form-label"></label>
                              <!-- Buttons start -->
                              <div class="d-flex gap-2 justify-content-end">
                                <button type="button" @click="submitChart" class="btn btn-primary">
                                  Submit Appraisal
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
