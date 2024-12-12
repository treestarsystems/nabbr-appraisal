import { Ref } from 'vue';
import axios from 'axios';
import router from '../router';
import { SwalToastSuccessHelper, SwalToastErrorHelper } from './sweetalertHelper';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { Chart } from '../types/chartTypes';

export async function getAppraisalChartHelper(swal: any, token: string, appraisalId?: string): Promise<Chart | void> {
  try {
    const chartRequest = {
      method: 'GET',
      url: appraisalId ? `/api/v1/appraisal/chart/${appraisalId}` : '/api/v1/appraisal/chart/template',
      headers: {
        token: token,
      },
    };
    const response: ResponseObjectDefaultInterface = (await axios(chartRequest))?.data;
    if (response.httpStatus > 299) {
      SwalToastErrorHelper(swal, response?.message);
      if (response.httpStatus === 404) router.push('/appraisal');
      return;
    }
    const chartTemplate: Chart = response.payload[0];
    SwalToastSuccessHelper(swal, 'Appraisal Loaded Successfully');
    return chartTemplate;
  } catch (err) {
    SwalToastErrorHelper(swal, err);
    return;
  }
}

export async function getAppraisalChartAllHelper(swal: any, token: string): Promise<Chart[] | void> {
  try {
    const chartRequest = {
      method: 'GET',
      url: '/api/v1/appraisal/chart',
      headers: {
        token: token,
      },
    };
    const response: ResponseObjectDefaultInterface = (await axios(chartRequest))?.data;
    if (response.httpStatus > 299) {
      SwalToastErrorHelper(swal, response?.message);
      return;
    }
    return response.payload as Chart[];
  } catch (err) {
    SwalToastErrorHelper(swal, err);
    return;
  }
}

export async function postPutAppraisalChartHelper(
  swal: any,
  token: string,
  data: Chart,
  appraisalId?: string,
): Promise<void> {
  try {
    const chartRequest = {
      method: appraisalId ? 'PUT' : 'POST',
      url: appraisalId ? `/api/v1/appraisal/chart/${appraisalId}` : '/api/v1/appraisal/chart',
      headers: {
        token: token,
      },
      data: data,
    };
    const response: ResponseObjectDefaultInterface = (await axios(chartRequest))?.data;
    if (response.httpStatus > 299) {
      SwalToastErrorHelper(swal, response?.message);
      return;
    }
    SwalToastSuccessHelper(swal, `Appraisal Submitted Successfully (${response.payload[0]})`);
    router.go(0);
    return;
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
  return characteristic.every(char => char.score !== undefined && char.score !== null && char.score !== 'nil');
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
  totalScoreRef.value = parseFloat(total.toFixed(3));
}
