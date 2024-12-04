<script setup lang="ts">
import { ref, inject, toRaw, onMounted } from 'vue';
import { useAuthStore } from '../stores/auth';
import { getAppraisalChartTemplate } from '../helpers/chart';

let chartData = ref('');
onMounted(async () => {
  const authStore = useAuthStore();
  const swal: any = inject('$swal');
  const token = toRaw(authStore.getState)?.token as string;
  chartData.value = (await getAppraisalChartTemplate(swal, token)) as any;
});
</script>

<template>
  <!-- Row start -->
  <div class="col-xxl-12">
    <div class="card mb-3">
      <div class="card-body">
        <div class="table-responsive">
          <table class="table table-bordered m-0">
            <thead>
              <tr>
                <th>Main Division:</th>
                <th></th>
                <th>Subdivision:</th>
                <th>Characteristic 1:</th>
                <th><div class="d-flex justify-content-center">+</div></th>
                <th>Characteristic 2:</th>
                <th><div class="d-flex justify-content-center">+</div></th>
                <th>Characteristic 3:</th>
                <th><div class="d-flex justify-content-center">+</div></th>
                <th>Netto Deviation:</th>
                <th>Factor:</th>
                <th>Total:</th>
                <th>Value:</th>
                <th><div class="d-flex justify-content-center">%</div></th>
              </tr>
            </thead>
            <!-- <tbody v-html="rows"></tbody> -->
            <tbody>
              <template v-for="division in chartData?.appraisalInformation?.mainDivision">
                <template v-for="(characteristic, ci) in division.characteristics">
                  <tr>
                    <td class="align-middle">{{ ci == 0 ? division.name : '' }}</td>
                    <td class="align-middle">{{ ci == 0 ? division.percentageWeight : '' }}</td>
                    <td class="align-middle">{{ characteristic[0]?.subDivision?.name || '' }}</td>
                    <td class="align-middle">{{ characteristic[0].name }}</td>
                    <td class="align-middle"></td>
                    <td class="align-middle">{{ characteristic[1].name }}</td>
                    <td class="align-middle"></td>
                    <td class="align-middle">{{ characteristic[2].name }}</td>
                    <td class="align-middle"></td>
                    <td class="align-middle"></td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span class="badge border border-info text-info">{{ characteristic[0].factor }}</span>
                      </div>
                    </td>
                    <td class="align-middle"></td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span class="badge border border-info text-info">{{ characteristic[0].value }}</span>
                      </div>
                    </td>
                    <td class="align-middle"></td>
                  </tr>
                </template>
              </template>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <!-- Row end -->
</template>
