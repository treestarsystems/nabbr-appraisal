import { Ref } from 'vue';
import axios from 'axios';
import { SwalToastErrorHelper } from './sweetalertHelper';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { Chart } from '../types/chartTypes';

export async function getAppraisalChartTemplateHelper(swal: any, token: string): Promise<any> {
  try {
    const chartTemplatePostRequest = {
      method: 'GET',
      url: '/api/v1/appraisal/chart/template',
      headers: {
        token: token,
      },
    };
    const response: ResponseObjectDefaultInterface = (await axios(chartTemplatePostRequest))?.data;
    if (response.httpStatus > 299) {
      SwalToastErrorHelper(swal, response?.message);
      return;
    }
    const chartTemplate: Chart = response.payload[0];
    return chartTemplate;
  } catch (err) {
    SwalToastErrorHelper(swal, err);
    return;
  }
}

export function generateRadioIdsHelper(...string: string[]) {
  const id = string.reduce(function (acc, cur) {
    return acc.split(' ').join('-').toLowerCase() + `-${cur.split(' ').join('-').toLowerCase()}`;
  });
  return id;
}

export function calculateTotalHelper(characteristics: any[]): number {
  return characteristics.reduce((total, characteristic) => {
    const value = parseInt(characteristic?.score) || 0;
    return total + value;
  }, 0);
}

export function allRadiosFilledHelper(characteristic: any[]): boolean {
  return characteristic.every(char => char.score !== undefined && char.score !== null && char.score !== 0);
}

export function updateTotalScoreHelper(document: any, totalScoreRef: Ref<number, number>) {
  const rows = document.querySelectorAll('tbody tr');
  let total = 0;
  for (const row of rows) {
    const lastTd = row.querySelector('td:last-child span');
    if (lastTd) {
      total += parseFloat(lastTd.textContent || '0');
    }
  }
  totalScoreRef.value = parseFloat(total.toFixed(2));
}
