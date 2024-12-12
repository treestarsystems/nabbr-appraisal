<script setup lang="ts">
import { ref, inject, toRaw, provide, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../stores/authStore';
import { getAppraisalChartHelper, postPutAppraisalChartHelper } from '../helpers/chartHelper';
import { generateCalendarDateStringHelper } from '../helpers/utilsHelper';
import { SwalToastErrorHelper } from '../helpers/sweetalertHelper';
import AppraisalBodyChartComponent from './AppraisalBodyChartComponent.vue';
import { Chart } from '../types/chartTypes';

const totalScore = ref(0);
const swal: any = inject('$swal');
const route = useRoute();
const chartData = ref<Chart>();
const authStore = useAuthStore();
const token = toRaw(authStore.getState)?.token as string;
provide('chartData', chartData);
provide('totalScore', totalScore);

onMounted(async () => {
  try {
    if (!route.params?.appraisalId) {
      const chart = (await getAppraisalChartHelper(swal, token)) as Chart;
      chartData.value = chart;
    } else {
      const appraisalId = route.params.appraisalId as string;
      chartData.value = (await getAppraisalChartHelper(swal, token, appraisalId)) as Chart;
    }
  } catch (err: any) {
    SwalToastErrorHelper(swal, err);
  }
});

async function submitChart() {
  try {
    // TODO: Fire Swal then return early if all fields are not checked.

    // We have to get these because they are
    const appraisalTotalScore: any = document.getElementById('appraisalTotalScore');
    if (chartData.value && chartData.value.appraisalInformation) {
      chartData.value.appraisalInformation.appraisalScore = parseFloat(appraisalTotalScore.value.replace('%', ''));
    }
    if (!route.params?.appraisalId) {
      await postPutAppraisalChartHelper(swal, token, chartData.value as Chart);
    } else {
      const appraisalId = route.params.appraisalId as string;
      await postPutAppraisalChartHelper(swal, token, chartData.value as Chart, appraisalId);
    }
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
                      disabled="true"
                      placeholder="Total Score"
                      :value="`${
                        totalScore === 0 ? chartData?.appraisalInformation?.appraisalScore ?? 0 : totalScore
                      }%`"
                      @change="
                        chartData?.appraisalInformation &&
                          $event.target &&
                          (chartData.appraisalInformation.appraisalScore = parseFloat(
                            ($event.target as HTMLInputElement).value,
                          ))
                      "
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
                                    :value="
                                      chartData?.memberInformation?.name ? chartData?.memberInformation?.name : ''
                                    "
                                    @input="
                                      chartData?.memberInformation &&
                                        $event.target &&
                                        (chartData.memberInformation.name = ($event.target as HTMLInputElement).value)
                                    "
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
                                    :value="
                                      chartData?.memberInformation?.memberNumber
                                        ? chartData?.memberInformation?.memberNumber
                                        : ''
                                    "
                                    @input="
                                      chartData?.memberInformation &&
                                        $event.target &&
                                        (chartData.memberInformation.memberNumber = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
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
                                    :value="chartData?.petInformation?.name ? chartData?.petInformation?.name : ''"
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.name = ($event.target as HTMLInputElement).value)
                                    "
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
                                    :value="
                                      chartData?.petInformation?.registrationNumber
                                        ? chartData?.petInformation?.registrationNumber
                                        : ''
                                    "
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.registrationNumber = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
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
                                    :value="
                                      chartData?.petInformation?.microchip ? chartData?.petInformation?.microchip : ''
                                    "
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.microchip = ($event.target as HTMLInputElement).value)
                                    "
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
                                    :value="
                                      chartData?.petInformation?.dnaNumber ? chartData?.petInformation?.dnaNumber : ''
                                    "
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.dnaNumber = ($event.target as HTMLInputElement).value)
                                    "
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
                                    :value="chartData?.petInformation?.color ? chartData?.petInformation?.color : ''"
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.color = ($event.target as HTMLInputElement).value)
                                    "
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
                                    :value="
                                      chartData?.petInformation?.markings ? chartData?.petInformation?.markings : ''
                                    "
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.markings = ($event.target as HTMLInputElement).value)
                                    "
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
                                    :value="chartData?.petInformation?.weight ? chartData?.petInformation?.weight : ''"
                                    @input="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.weight = parseInt(
                                          ($event.target as HTMLInputElement).value,
                                        ))
                                    "
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
                                    @change="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.age = parseInt(
                                          ($event.target as HTMLSelectElement).value,
                                        ))
                                    "
                                    class="form-select"
                                    id="age"
                                    aria-label="Default select example"
                                    required
                                  >
                                    <!-- <option :selected="chartData?.petInformation?.age === 0" value="0">&lt;1yr</option> -->
                                    <option :selected="chartData?.petInformation?.age === 1" value="1">1yr</option>
                                    <option :selected="chartData?.petInformation?.age === 2" value="2">2yr</option>
                                    <option :selected="chartData?.petInformation?.age === 3" value="3">4yr</option>
                                    <option :selected="chartData?.petInformation?.age === 4" value="4">4yr</option>
                                    <option :selected="chartData?.petInformation?.age === 5" value="5">5yr</option>
                                    <option :selected="chartData?.petInformation?.age === 6" value="6">6yr</option>
                                    <option :selected="chartData?.petInformation?.age === 7" value="7">7yr</option>
                                    <option :selected="chartData?.petInformation?.age === 8" value="8">8yr</option>
                                    <option :selected="chartData?.petInformation?.age === 9" value="9">9yr</option>
                                    <option :selected="chartData?.petInformation?.age === 10" value="10">10yr</option>
                                    <option disabled :selected="!chartData?.petInformation?.age">Select Age</option>
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
                                    @change="
                                      chartData?.petInformation &&
                                        $event.target &&
                                        (chartData.petInformation.sex = ($event.target as HTMLSelectElement).value)
                                    "
                                    class="form-select"
                                    id="sex"
                                    aria-label="Default select example"
                                    required
                                  >
                                    <option :selected="chartData?.petInformation?.sex === 'male'" value="male">
                                      Male
                                    </option>
                                    <option :selected="chartData?.petInformation?.sex === 'female'" value="female">
                                      Female
                                    </option>
                                    <option disabled :selected="!chartData?.petInformation?.sex" value="">
                                      Select Sex
                                    </option>
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
                                    :value="
                                      chartData?.appraisalInformation?.seniorAppraiserName
                                        ? chartData?.appraisalInformation?.seniorAppraiserName
                                        : ''
                                    "
                                    @input="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.seniorAppraiserName = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
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
                                    :value="
                                      chartData?.appraisalInformation?.seniorAppraiserNumber
                                        ? chartData?.appraisalInformation?.seniorAppraiserNumber
                                        : ''
                                    "
                                    @input="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.seniorAppraiserNumber = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
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
                                    :value="
                                      chartData?.appraisalInformation?.appraiserName
                                        ? chartData?.appraisalInformation?.appraiserName
                                        : ''
                                    "
                                    @input="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.appraiserName = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
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
                                    :value="
                                      chartData?.appraisalInformation?.appraiserNumber
                                        ? chartData?.appraisalInformation?.appraiserNumber
                                        : ''
                                    "
                                    @input="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.appraiserNumber = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
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
                                    :value="
                                      chartData?.appraisalInformation?.additionalComments
                                        ? chartData?.appraisalInformation?.additionalComments
                                        : ''
                                    "
                                    @input="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.additionalComments = (
                                          $event.target as HTMLTextAreaElement
                                        ).value)
                                    "
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
                                    :value="
                                      chartData?.appraisalInformation?.date
                                        ? chartData?.appraisalInformation?.date
                                        : generateCalendarDateStringHelper()
                                    "
                                    @change="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.date = (
                                          $event.target as HTMLInputElement
                                        ).value)
                                    "
                                    type="date"
                                    id="appraisalDate"
                                    class="form-control datepicker"
                                    name="date"
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
                                    @change="
                                      chartData?.appraisalInformation &&
                                        $event.target &&
                                        (chartData.appraisalInformation.place = parseInt(
                                          ($event.target as HTMLSelectElement).value,
                                        ))
                                    "
                                    class="form-select"
                                    id="appraisalPlace"
                                    aria-label="Default select example"
                                  >
                                    <option :selected="chartData?.appraisalInformation?.place === 1" value="1">
                                      1st
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 2" value="2">
                                      2nd
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 3" value="3">
                                      3rd
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 4" value="4">
                                      4th
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 5" value="5">
                                      5th
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 6" value="6">
                                      6th
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 7" value="7">
                                      7th
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 8" value="8">
                                      8th
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 9" value="9">
                                      9th
                                    </option>
                                    <option :selected="chartData?.appraisalInformation?.place === 10" value="10">
                                      10th
                                    </option>
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
