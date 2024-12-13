<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { ref, inject, toRaw, onMounted } from 'vue';
import { useAuthStore } from '../stores/authStore';
import { getAppraisalChartAllHelper } from '../helpers/chartHelper';
import { SwalToastErrorHelper, SwalConfirmationDeleteHelper } from '../helpers/sweetalertHelper';
import { Chart } from '../types/chartTypes';
import { UserState } from '../types/authTypes';

const swal: any = inject('$swal');
const authStore = useAuthStore();
const chartDataAll = ref<Chart[]>();
const token = toRaw(authStore.getState as UserState)?.token as string;

function deleteAppraisal(appraisalId: string, message: string) {
  SwalConfirmationDeleteHelper(swal, message);
}
onMounted(async () => {
  try {
    chartDataAll.value = (await getAppraisalChartAllHelper(swal, token)) as Chart[];
  } catch (err: any) {
    SwalToastErrorHelper(swal, err);
  }
});
</script>
<template>
  <!-- App body starts -->
  <div class="app-body">
    <div class="app-body">
      <!-- Row start -->
      <div class="col-xxl-12">
        <div class="card mb-3">
          <div class="card-body">
            <div class="table-responsive">
              <table class="table table-bordered m-0">
                <thead>
                  <tr>
                    <th>
                      <div class="d-flex justify-content-center">
                        <i class="bi bi-list fs-5"></i>
                      </div>
                    </th>
                    <th>Appraiser Name:</th>
                    <th>Member Name:</th>
                    <th>Dog Name:</th>
                    <th>Sex:</th>
                    <th>Age:</th>
                    <th>Weight:</th>
                    <th>Score:</th>
                    <th>MC#:</th>
                    <th>DNA#:</th>
                    <th>
                      <div class="d-flex justify-content-center">
                        <i class="bi bi-list fs-5"></i>
                      </div>
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <template v-for="chart in chartDataAll">
                    <tr>
                      <td class="align-items-center">
                        <div class="d-flex justify-content-center">
                          <RouterLink :to="`/appraisal/${chart.appraisalId}`"
                            ><i class="bi bi-calculator fs-4 text-info"></i
                          ></RouterLink>
                        </div>
                      </td>
                      <td class="align-middle">{{ chart.appraisalInformation.appraiserName }}</td>
                      <td class="align-middle">{{ chart.memberInformation.name }}</td>
                      <td class="align-middle">{{ chart.petInformation.name }}</td>
                      <td class="align-middle">{{ chart.petInformation.sex.toUpperCase() }}</td>
                      <td class="align-middle">{{ chart.petInformation.age }}</td>
                      <td class="align-middle">{{ chart.petInformation.weight }}</td>
                      <td class="align-middle">
                        <div class="d-flex justify-content-center">
                          <span class="not-allowed badge border border-success text-success"
                            >{{ chart.appraisalInformation.appraisalScore }}%</span
                          >
                        </div>
                      </td>
                      <td class="align-middle">{{ chart.petInformation.microchip }}</td>
                      <td class="align-middle">{{ chart.petInformation.dnaNumber }}</td>
                      <td class="align-middle">
                        <div class="d-flex justify-content-center">
                          <i
                            @click="
                              deleteAppraisal(
                                chart.appraisalId,
                                `Delete Appraisal for ${chart.memberInformation.name}'s dog ${chart.petInformation.name}`,
                              )
                            "
                            class="bi bi-trash fs-4 text-danger"
                          ></i>
                        </div>
                      </td>
                    </tr>
                  </template>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
      <!-- Row end -->
    </div>
  </div>
  <!-- App body end -->
</template>
