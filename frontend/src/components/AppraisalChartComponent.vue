<script setup lang="ts">
import { ref, inject } from 'vue';
import { generateRadioIds } from '../helpers/chart';
import { Chart } from '../types/chart';
const chartData = inject<Chart>('chartData');
// TODO: Ensure this works properly.
// let totalScore = inject<number>('totalScore') || 0;
const totalScore = inject<number>('totalScore');
const totalScoreRef = ref(totalScore || 0);

function calculateTotal(characteristics: any[]): number {
  return characteristics.reduce((total, characteristic) => {
    const value = parseInt(characteristic?.score) || 0;
    return total + value;
  }, 0);
}

function allRadiosFilled(characteristic: any[]): boolean {
  return characteristic.every(char => char.score !== undefined && char.score !== null && char.score !== 0);
}

function updateTotalScore() {
  const rows = document.querySelectorAll('tbody tr');
  let total = 0;
  rows.forEach(row => {
    const lastTd = row.querySelector('td:last-child span');
    if (lastTd) {
      total += parseFloat(lastTd.textContent || '0');
    }
  });
  totalScoreRef.value = total;
}
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
            <tbody>
              <template v-for="division in chartData?.appraisalInformation?.mainDivision">
                <template v-for="(characteristic, ci) in division.characteristics">
                  <tr>
                    <td class="align-middle">{{ ci == 0 ? division.name : '' }}</td>
                    <td class="align-middle">{{ ci == 0 ? division.percentageWeight : '' }}</td>
                    <td class="align-middle">{{ characteristic[0]?.subDivision?.name || '' }}</td>
                    <td class="align-middle">{{ characteristic[0].name }}</td>
                    <td class="align-middle p-10">
                      <div class="col-sm-12 col-12">
                        <div class="row">
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label
                              class="form-label"
                              :for="`${generateRadioIds(division.name, characteristic[0].name, 'minus')}`"
                              >-</label
                            >
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label :for="`${generateRadioIds(division.name, characteristic[0].name, 'zero')}`">0</label>
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label :for="`${generateRadioIds(division.name, characteristic[0].name, 'plus')}`">+</label>
                          </div>
                        </div>
                        <div class="row was-validated">
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input was-validate"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[0].name, 'minus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[0].name, 'score')}`"
                              value="0"
                              @input="characteristic[0].score = $event.target.value - 1"
                              @change="updateTotalScore"
                              required
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[0].name, 'zero')}`"
                              :name="`${generateRadioIds(division.name, characteristic[0].name, 'score')}`"
                              value="0"
                              @input="characteristic[0].score = $event.target.value"
                              @change="updateTotalScore"
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[0].name, 'plus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[0].name, 'score')}`"
                              value="0"
                              @input="characteristic[0].score = $event.target.value + 1"
                              @change="updateTotalScore"
                            />
                          </div>
                        </div>
                      </div>
                    </td>
                    <td class="align-middle">{{ characteristic[1].name }}</td>
                    <td class="align-middle">
                      <div class="col-sm-12 col-12">
                        <div class="row">
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label
                              class="form-label"
                              :for="`${generateRadioIds(division.name, characteristic[1].name, 'minus')}`"
                              >-</label
                            >
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label :for="`${generateRadioIds(division.name, characteristic[1].name, 'zero')}`">0</label>
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label :for="`${generateRadioIds(division.name, characteristic[1].name, 'plus')}`">+</label>
                          </div>
                        </div>
                        <div class="row was-validated">
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input was-validate"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[1].name, 'minus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[1].name, 'score')}`"
                              value="0"
                              @input="characteristic[1].score = $event.target.value - 1"
                              @change="updateTotalScore"
                              required
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[1].name, 'zero')}`"
                              :name="`${generateRadioIds(division.name, characteristic[1].name, 'score')}`"
                              value="0"
                              @input="characteristic[1].score = $event.target.value"
                              @change="updateTotalScore"
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[1].name, 'plus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[1].name, 'score')}`"
                              value="0"
                              @input="characteristic[1].score = $event.target.value + 1"
                              @change="updateTotalScore"
                            />
                          </div>
                        </div>
                      </div>
                    </td>
                    <td class="align-middle">{{ characteristic[2].name }}</td>
                    <td class="align-middle">
                      <div class="col-sm-12 col-12">
                        <div class="row">
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label
                              class="form-label"
                              :for="`${generateRadioIds(division.name, characteristic[2].name, 'minus')}`"
                              >-</label
                            >
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label :for="`${generateRadioIds(division.name, characteristic[2].name, 'zero')}`">0</label>
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <label :for="`${generateRadioIds(division.name, characteristic[2].name, 'plus')}`">+</label>
                          </div>
                        </div>
                        <div class="row was-validated">
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input was-validate"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[2].name, 'minus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[2].name, 'score')}`"
                              value="0"
                              @input="characteristic[2].score = $event.target.value - 1"
                              @change="updateTotalScore"
                              required
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[2].name, 'zero')}`"
                              :name="`${generateRadioIds(division.name, characteristic[2].name, 'score')}`"
                              value="0"
                              @input="characteristic[2].score = $event.target.value"
                              @change="updateTotalScore"
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[2].name, 'plus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[2].name, 'score')}`"
                              value="0"
                              @input="characteristic[2].score = $event.target.value + 1"
                              @change="updateTotalScore"
                            />
                          </div>
                        </div>
                      </div>
                    </td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span v-if="allRadiosFilled(characteristic)" class="badge border border-white text-white">{{
                          calculateTotal(division.characteristics[ci])
                        }}</span>
                        <span v-else></span>
                      </div>
                    </td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span class="badge border border-muted text-muted">{{ characteristic[0].factor }}</span>
                      </div>
                    </td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span v-if="allRadiosFilled(characteristic)" class="badge border border-info text-info">{{
                          calculateTotal(division.characteristics[ci]) + characteristic[0].factor
                        }}</span>
                        <span v-else></span>
                      </div>
                    </td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span class="badge border border-muted text-muted">{{ characteristic[0].value }}</span>
                      </div>
                    </td>
                    <td class="align-middle">
                      <div class="d-flex justify-content-center">
                        <span v-if="allRadiosFilled(characteristic)" class="badge border border-success text-success">{{
                          ((calculateTotal(division.characteristics[ci]) + characteristic[0].factor) *
                            characteristic[0].value) /
                          10
                        }}</span>
                        <span v-else></span>
                      </div>
                    </td>
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
