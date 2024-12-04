<script setup lang="ts">
import { ref, inject, toRaw, provide, reactive, onMounted, onBeforeMount } from 'vue';
import { useAuthStore } from '../stores/auth';
import { getAppraisalChartTemplate } from '../helpers/chart';
import { generateCalendarDateString } from '../helpers/utils';
import { Chart } from '../types/chart';
import AppraisalChartComponent from './AppraisalChartComponent.vue';
import { formatDate } from '@vueuse/core';

let chartData = ref('');
provide('chartData', chartData);
// const formChartData = reactive({
//   appraisalId: '',
//   memberInformation: {
//     name: '',
//     email: '',
//     memberNumber: '',
//   },
//   petInformation: {
//     name: '',
//     type: '',
//     breed: '',
//     age: '',
//     dnaNumber: '',
//     weight: '',
//     color: '',
//     markings: '',
//     microchip: '',
//     sex: '',
//     registrationNumber: '',
//   },
//   appraisalInformation: {
//     mainDivision: {
//       appearance: {
//         characteristics: [
//           [
//             {
//               name: 'Big And Strong',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Gender Authenticity',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Balance',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Impressive',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Musculature',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Bearing',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//         ],
//         name: 'General Appearance',
//         percentageWeight: 8,
//       },
//       head: {
//         characteristics: [
//           [
//             {
//               name: 'Short',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Square',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Large in circumference',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Skull',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Broad',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Filled',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Flat',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Skull',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Deep',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Characteristics Typically Boerboel',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Muscular',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Skull',
//               },
//             },
//           ],
//         ],
//         name: 'Head',
//         percentageWeight: 7,
//       },
//       face: {
//         characteristics: [
//           [
//             {
//               name: 'Fusion (Skull And Mouth)',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Lips',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Straight And Parallel',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Nasal Bone',
//               },
//             },
//             {
//               name: 'Eyes-Setting',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Eyes',
//               },
//             },
//             {
//               name: 'Earflaps-Setting',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Earflaps',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Stop',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Teeth',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Deep And Broad',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Nasal Bone',
//               },
//             },
//             {
//               name: 'Eyelids',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Eyes',
//               },
//             },
//             {
//               name: 'Shape',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Earflaps',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Well Filled Between Eyes',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Jaws',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Length 8-10 Cm',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Nasal Bone',
//               },
//             },
//             {
//               name: 'Colour And Pigmentation',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Eyes',
//               },
//             },
//             {
//               name: 'Propotion',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Earflaps',
//               },
//             },
//           ],
//         ],
//         name: 'Face',
//         percentageWeight: 15,
//       },
//       neck: {
//         characteristics: [
//           [
//             {
//               name: 'Shape',
//               score: 0,
//               value: 5,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Length',
//               score: 0,
//               value: 5,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Dewlap',
//               score: 0,
//               value: 5,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//         ],
//         name: 'Neck',
//         percentageWeight: 5,
//       },
//       forequarter: {
//         characteristics: [
//           [
//             {
//               name: 'Attachment',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Shoulders',
//               },
//             },
//             {
//               name: 'Thick And Strong',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Forelegs',
//               },
//             },
//             {
//               name: 'Length',
//               score: 0,
//               value: 2,
//               factor: 7,
//               subDivision: {
//                 name: 'Pasterns',
//               },
//             },
//             {
//               name: 'Size',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Forepaws',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Angulation',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Shoulders',
//               },
//             },
//             {
//               name: 'Musculature',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Forelegs',
//               },
//             },
//             {
//               name: 'Thickness',
//               score: 0,
//               value: 2,
//               factor: 7,
//               subDivision: {
//                 name: 'Pasterns',
//               },
//             },
//             {
//               name: 'Shape',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Forepaws',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Elbow',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Shoulders',
//               },
//             },
//             {
//               name: 'Vertical',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: 'Forelegs',
//               },
//             },
//             {
//               name: 'Position',
//               score: 0,
//               value: 2,
//               factor: 7,
//               subDivision: {
//                 name: 'Pasterns',
//               },
//             },
//             {
//               name: 'Tread',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Forepaws',
//               },
//             },
//           ],
//         ],
//         name: 'Forequarter',
//         percentageWeight: 12,
//       },
//       centerPiece: {
//         characteristics: [
//           [
//             {
//               name: 'Topline',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Back',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Broad',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Chest',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Loin',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Crop',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Deep',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Chest',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Tail',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Musculature',
//               score: 0,
//               value: 3,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Ribcage',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Chest',
//               },
//             },
//           ],
//         ],
//         name: 'Center Piece',
//         percentageWeight: 10,
//       },
//       hindquarter: {
//         characteristics: [
//           [
//             {
//               name: 'Strong And Sturdy',
//               score: 0,
//               value: 11,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Shape And Size',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Hind Paws',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Angulation',
//               score: 0,
//               value: 11,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Hind Pasterns, Hocks',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Hind Paws',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Stance',
//               score: 0,
//               value: 11,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//             {
//               name: 'Tread',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: 'Hind Paws',
//               },
//             },
//           ],
//         ],
//         name: 'Hindquarter',
//         percentageWeight: 15,
//       },
//       skinCoat: {
//         characteristics: [
//           [
//             {
//               name: 'Thick And Loose Skin',
//               score: 0,
//               value: 5,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Short And Thick Hair',
//               score: 0,
//               value: 5,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Pigmentation',
//               score: 0,
//               value: 5,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//         ],
//         name: 'Skin/Coat',
//         percentageWeight: 5,
//       },
//       health: {
//         characteristics: [
//           [
//             {
//               name: 'Condition Versus Age',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Genitals',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Condition',
//               score: 0,
//               value: 4,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//         ],
//         name: 'Health',
//         percentageWeight: 4,
//       },
//       temperament: {
//         characteristics: [
//           [
//             {
//               name: 'Obedient And Manageable',
//               score: 0,
//               value: 8,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Reliable',
//               score: 0,
//               value: 8,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Self-Confident',
//               score: 0,
//               value: 8,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//         ],
//         name: 'Temperament',
//         percentageWeight: 8,
//       },
//       movement: {
//         characteristics: [
//           [
//             {
//               name: 'Buoyant',
//               score: 0,
//               value: 8,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Parallel',
//               score: 0,
//               value: 8,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//           [
//             {
//               name: 'Topline',
//               score: 0,
//               value: 8,
//               factor: 7,
//               subDivision: {
//                 name: '',
//               },
//             },
//           ],
//         ],
//         name: 'Movement',
//         percentageWeight: 8,
//       },
//     },
//     appraiser: '',
//     appraiserNumber: '',
//     seniorAppraiser: '',
//     seniorAppraiserNumber: '',
//     date: '',
//     value: '',
//     notes: '',
//     additionalComments: '',
//     place: '',
//   },
// });

// const formChartData = reactive({
//   template: null,
// });

onMounted(async () => {
  const authStore = useAuthStore();
  const swal: any = inject('$swal');
  const token = toRaw(authStore.getState)?.token as string;
  const chartTemplate = await getAppraisalChartTemplate(swal, token);
  chartData.value = chartTemplate as any;
  // formChartData.template = chartTemplate;
});

async function submitChart() {
  // const t = formChartData.template as any;
  // console.log(t?.memberInformation?.name);
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
                                  <input
                                    v-model="formChartData.memberInformation.name"
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
                  <AppraisalChartComponent />
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
