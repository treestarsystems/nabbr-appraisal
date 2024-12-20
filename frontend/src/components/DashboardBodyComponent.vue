<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { ref, inject, toRaw, onMounted, Ref } from 'vue';
import { useAuthStore } from '../stores/authStore';
import { getAppraisalChartHelper, getAppraisalChartAllHelper } from '../helpers/chartHelper';
import { SwalToastErrorHelper, SwalConfirmationDeleteAppraisalHelper } from '../helpers/sweetalertHelper';
import { Chart, Division } from '../types/chartTypes';
import { UserState } from '../types/authTypes';
import * as ExcelJS from 'exceljs';

const swal: any = inject('$swal');
const authStore = useAuthStore();
const chartDataAll = ref<Chart[]>();
const chartData = ref<Chart>();
const token = toRaw(authStore.getState as UserState)?.token as string;

async function deleteAppraisal(appraisalId: string, message: string) {
  await SwalConfirmationDeleteAppraisalHelper(swal, message, token, appraisalId, chartDataAll as Ref<Chart[]>);
}

async function generateSpreadsheet(appraisalId: string) {
  try {
    chartData.value = (await getAppraisalChartHelper(swal, token, appraisalId)) as Chart;
    const workbook = new ExcelJS.Workbook();
    workbook.creator = 'NABBR Appraisal';
    workbook.lastModifiedBy = 'NABBR Appraisal Tool';
    workbook.description = `Appraisal for ${chartData.value.memberInformation.name}'s dog ${chartData.value.petInformation.name} (${chartData.value.appraisalId})`;
    workbook.company = 'North American Boerboel Registry';
    workbook.created = new Date();
    workbook.modified = new Date();
    workbook.lastPrinted = new Date();
    const sheet = workbook.addWorksheet(`Appraisal ${chartData.value.petInformation.name}`, {
      pageSetup: { fitToPage: true, fitToHeight: 5, fitToWidth: 7, paperSize: 9, orientation: 'landscape' },
    });
    sheet.columns = [
      { key: 'mainDivision', width: 20 },
      { key: 'mainDivisionPercentage', width: 5, style: { alignment: { horizontal: 'center' } } },
      { key: 'subDivision', width: 12 },
      { key: 'characteristic1', width: 22 },
      { key: 'characteristic1Score', width: 5, style: { alignment: { horizontal: 'center' } } },
      { key: 'characteristic2', width: 22 },
      { key: 'characteristic2Score', width: 5, style: { alignment: { horizontal: 'center' } } },
      { key: 'characteristic3', width: 27 },
      { key: 'characteristic3Score', width: 5, style: { alignment: { horizontal: 'center' } } },
      { key: 'nettoDeviation', width: 15, style: { alignment: { horizontal: 'center' } } },
      { key: 'factor', width: 7, style: { alignment: { horizontal: 'center' } } },
      { key: 'total', width: 6, style: { alignment: { horizontal: 'center' } } },
      { key: 'value', width: 6, style: { alignment: { horizontal: 'center' } } },
      { key: 'percent', width: 5, style: { alignment: { horizontal: 'center' } } },
    ];
    // Add title row and merge cells
    const titleRow = sheet.getRow(1);
    titleRow.values = ['Appraisal Chart for North American Boerboel Registry'];
    sheet.mergeCells('A1:N1');
    titleRow.getCell(1).alignment = { horizontal: 'center' };
    titleRow.font = { bold: true, size: 16 };

    // Pet Info row
    const petInfoRow = sheet.getRow(2);
    petInfoRow.values = [
      "Dog's Full Name:",
      chartData.value.petInformation.name,
      '',
      '',
      '',
      "Dog's Registration Number:",
      chartData.value.petInformation.registrationNumber,
      '',
      '',
      '',
      '',
      '',
      '',
      '',
    ];
    sheet.mergeCells('B2:D2');
    sheet.mergeCells('G2:I2');
    petInfoRow.getCell(2).alignment = { horizontal: 'left', vertical: 'middle' };
    petInfoRow.getCell(7).alignment = { horizontal: 'left', vertical: 'middle' };
    petInfoRow.eachCell(cell => {
      cell.font = { bold: true };
      cell.alignment = { horizontal: 'left', vertical: 'middle' };
    });

    // Member Info row
    const memberInfoRow = sheet.getRow(3);
    memberInfoRow.values = [
      'Member Name:',
      chartData.value.memberInformation.name,
      '',
      '',
      '',
      'Member Number:',
      chartData.value.memberInformation.memberNumber,
      '',
      '',
      '',
      '',
      '',
      '',
      '',
    ];
    sheet.mergeCells('B3:D3');
    sheet.mergeCells('G3:I3');
    memberInfoRow.getCell(2).alignment = { horizontal: 'left', vertical: 'middle' };
    memberInfoRow.getCell(7).alignment = { horizontal: 'left', vertical: 'middle' };
    memberInfoRow.eachCell(cell => {
      cell.font = { bold: true };
      cell.alignment = { horizontal: 'left', vertical: 'middle' };
    });

    // Microchip, DNA Info row
    const petInfoRowIds = sheet.getRow(4);
    petInfoRowIds.values = [
      'Microchip Number:',
      chartData.value.petInformation.microchip,
      '',
      '',
      '',
      'DNA Number:',
      chartData.value.petInformation.dnaNumber,
      '',
      '',
      '',
      '',
      '',
      '',
      '',
    ];
    sheet.mergeCells('B4:D4');
    sheet.mergeCells('G4:I4');
    petInfoRowIds.getCell(2).alignment = { horizontal: 'left', vertical: 'middle' };
    petInfoRowIds.getCell(7).alignment = { horizontal: 'left', vertical: 'middle' };
    petInfoRowIds.eachCell(cell => {
      cell.font = { bold: true };
      cell.alignment = { horizontal: 'left', vertical: 'middle' };
    });

    // Other Pet Info row
    const petInfoRowDetail = sheet.getRow(5);
    petInfoRowDetail.values = [
      'Colour:',
      chartData.value.petInformation.color,
      '',
      '',
      '',
      'Mask:',
      chartData.value.petInformation.markings,
      '',
      '',
      'Age:',
      chartData.value.petInformation.age,
      'Sex:',
      chartData.value.petInformation.sex.toUpperCase(),
      '',
    ];

    sheet.mergeCells('B5:D5');
    sheet.mergeCells('G5:H5');
    petInfoRowDetail.eachCell(cell => {
      cell.font = { bold: true };
      cell.alignment = { horizontal: 'left', vertical: 'middle' };
    });
    petInfoRowDetail.getCell(13).font = { bold: true, size: 8 };

    // Separator row
    sheet.mergeCells('A6:N6');
    // Add header row
    const headerRow = sheet.getRow(7);
    headerRow.values = [
      'Main Division',
      '',
      'Sub-Division',
      'Characteristic 1',
      '+',
      'Characteristic 2',
      '+',
      'Characteristic 3',
      '+',
      'Netto Deviation',
      'Factor',
      'Total',
      'Value',
      '%',
    ];
    headerRow.eachCell(cell => {
      cell.font = { bold: true };
      cell.alignment = { horizontal: 'center', vertical: 'middle' };
      cell.border = {
        top: { style: 'thin' },
        left: { style: 'thin' },
        bottom: { style: 'thin' },
        right: { style: 'thin' },
      };
    });
    let chartStartingRow = 8;
    for (const divisionKey in chartData.value.appraisalInformation.mainDivision) {
      const division = chartData.value.appraisalInformation.mainDivision[
        divisionKey as keyof typeof chartData.value.appraisalInformation.mainDivision
      ] as Division;
      // for (const characteristic of division.characteristics) {
      for (let i = 0; i < division.characteristics.length; i++) {
        const characteristic = division.characteristics[i];
        const nettoDeviation =
          Number(characteristic[0]?.score) + Number(characteristic[1]?.score) + Number(characteristic[2]?.score);
        const factor = characteristic[0].factor;
        const rowValue = characteristic[0].value;
        const rowTotal = nettoDeviation + factor;
        const rowPercent = (rowValue * rowTotal) / 10;
        const characteristicRow = sheet.getRow(chartStartingRow);
        characteristicRow.values = [
          i == 0 ? division.name : '',
          i == 0 ? division.percentageWeight : '',
          characteristic[0]?.subDivision?.name || '',
          characteristic[0]?.name,
          characteristic[0]?.score,
          characteristic[1]?.name,
          characteristic[1]?.score,
          characteristic[2]?.name,
          characteristic[2]?.score,
          nettoDeviation,
          factor,
          rowTotal,
          rowValue,
          rowPercent,
        ];
        characteristicRow.eachCell(cell => {
          cell.border = {
            top: { style: 'thin' },
            left: { style: 'thin' },
            bottom: { style: 'thin' },
            right: { style: 'thin' },
          };
        });
        chartStartingRow++;
      }
    }
    sheet.addRow({ value: 'Total:', percent: `${chartData.value.appraisalInformation.appraisalScore}%` });
    const characteristicRow = sheet.getRow(chartStartingRow);
    characteristicRow.eachCell(cell => {
      cell.font = { bold: true };
      cell.alignment = { horizontal: 'center', vertical: 'middle' };
      cell.border = {
        top: { style: 'thin' },
        left: { style: 'thin' },
        bottom: { style: 'thin' },
        right: { style: 'thin' },
      };
    });
    // Separator row
    sheet.mergeCells(`A${chartStartingRow + 1}:N${chartStartingRow + 1}`);
    // sheet.mergeCells(`L${chartStartingRow}:M${chartStartingRow}`);

    sheet.addRow(['Additional Commentary:', chartData.value.appraisalInformation.additionalComments]);
    sheet.addRow([
      'Senior Appraiser:',
      chartData.value.appraisalInformation.seniorAppraiserName,
      'Senior Appraiser Number:',
      chartData.value.appraisalInformation.seniorAppraiserNumber,
      'Appraiser:',
      chartData.value.appraisalInformation.appraiserName,
      'Appraiser Number:',
      chartData.value.appraisalInformation.appraiserNumber,
      'Place:',
      chartData.value.appraisalInformation.place,
      'Date:',
      chartData.value.appraisalInformation.date,
    ]);

    // Generate buffer and create a Blob for download
    const buffer = await workbook.xlsx.writeBuffer();
    const blob = new Blob([buffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `Appraisal_${chartData.value.memberInformation.name}_${chartData.value.petInformation.name}.xlsx`;
    a.click();
    URL.revokeObjectURL(url);
  } catch (err: any) {
    SwalToastErrorHelper(swal, err);
  }
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
                      <td class="align-middle">
                        <div class="d-flex justify-content-center">
                          <RouterLink :to="`/appraisal/${chart.appraisalId}`"
                            ><i class="bi bi-calculator fs-4 text-info"></i
                          ></RouterLink>
                        </div>
                      </td>
                      <td class="align-middle">
                        <div class="d-flex justify-content-center">
                          <i
                            @click="generateSpreadsheet(chart.appraisalId)"
                            class="bi bi-file-earmark-excel fs-4"
                            style="color: greenyellow; cursor: pointer"
                          ></i>
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
                            style="cursor: not-allowed"
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
