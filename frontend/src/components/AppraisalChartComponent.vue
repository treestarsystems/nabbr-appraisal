<script setup lang="ts">
import { inject } from 'vue';
import { generateRadioIds } from '../helpers/chart';
const chartData = inject('chartData');
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
              <template v-for="(division, dn) in chartData?.appraisalInformation?.mainDivision">
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
                              value="-1"
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
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[0].name, 'plus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[0].name, 'score')}`"
                              value="1"
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
                              value="-1"
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
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[1].name, 'plus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[1].name, 'score')}`"
                              value="1"
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
                              value="-1"
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
                            />
                          </div>
                          <div class="col-sm-4 d-flex justify-content-center">
                            <input
                              class="form-check-input chart-form-check-input"
                              type="radio"
                              :id="`${generateRadioIds(division.name, characteristic[2].name, 'plus')}`"
                              :name="`${generateRadioIds(division.name, characteristic[2].name, 'score')}`"
                              value="1"
                            />
                          </div>
                        </div>
                      </div>
                    </td>
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
