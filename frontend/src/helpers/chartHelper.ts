import { Ref } from 'vue';
import axios from 'axios';
import router from '../router';
import { SwalToastSuccessHelper, SwalToastErrorHelper } from './sweetalertHelper';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { Chart, Characteristic } from '../types/chartTypes';

/**
 * Retrieves the appraisal chart template or a specific appraisal chart by id.
 * @param {any} swal injected sweetalert2 instance
 * @param {string} token user token
 * @param {string} appraisalId appraisal id that is provided by the backend when creating a new appraisal
 * @returns {Promise<Chart | void>}
 */

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

/**
 * Retrieves all appraisal charts.
 * @param {any} swal injected sweetalert2 instance
 * @param {string} token user token
 * @returns {Promise<Chart[] | void>}
 */

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

/**
 * Sends a POST or PUT request to create or update an appraisal chart.
 * @param {any} swal injected sweetalert2 instance
 * @param {string} token user token
 * @param {Chart} chartData api request data
 * @param {string} appraisalId appraisal id that is provided by the backend when creating a new appraisal
 * @returns {Promise<void>}
 */

export async function postPutAppraisalChartHelper(
  swal: any,
  token: string,
  chartData: Chart,
  appraisalId?: string,
): Promise<void> {
  try {
    const chartRequest = {
      method: appraisalId ? 'PUT' : 'POST',
      url: appraisalId ? `/api/v1/appraisal/chart/${appraisalId}` : '/api/v1/appraisal/chart',
      headers: {
        token: token,
      },
      data: chartData,
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

/**
 * Sends a DELETE request to remove an appraisal chart.
 * @param {string} token user token
 * @param {string} appraisalId appraisal id that is provided by the backend when creating a new appraisal
 * @returns {Promise<ResponseObjectDefaultInterface>}
 */

export async function deleteAppraisalChartByIdHelper(
  token: string,
  appraisalId: string,
): Promise<ResponseObjectDefaultInterface> {
  try {
    const chartRequest = {
      method: 'DELETE',
      url: `/api/v1/appraisal/chart/${appraisalId}`,
      headers: {
        token: token,
      },
    };
    const response: ResponseObjectDefaultInterface = (await axios(chartRequest))?.data;
    return response;
  } catch (err: any) {
    return {
      status: 'failure',
      httpStatus: 500,
      message: `${err?.response?.data?.message || err?.data?.message || err?.message || err}`,
      payload: [],
    };
  }
}

/**
 * Generate a unique id for radio inputs.
 * @param {string[]} string strings to concatenate
 * @returns {string} concatenated string with hyphens
 */
export function generateRadioIdsHelper(...string: string[]): string {
  const id = string.reduce(function (acc, cur) {
    return acc.split(' ').join('-').toLowerCase() + `-${cur.split(' ').join('-').toLowerCase()}`;
  });
  return id;
}

/**
 * Calculate the total score of a characteristic.
 * @param {Characteristic[]} characteristics array of characteristics
 * @returns {number} total score
 */
export function calculateTotalHelper(characteristics: Characteristic[]): number {
  return characteristics.reduce((total, characteristic) => {
    const value = parseInt(characteristic?.score) || 0;
    return total + value;
  }, 0);
}

/**
 * Check if all radio inputs are filled.
 * @param {Characteristic[]} characteristic array of characteristics
 * @returns {boolean} true if all radios are filled
 */

export function allRadiosFilledHelper(characteristic: Characteristic[]): boolean {
  return characteristic.every(char => char.score !== undefined && char.score !== null && char.score !== 'nil');
}

/**
 * Update the total score of a characteristic.
 * @param {any} document document object from the browser
 * @param {Ref<number, number>} totalScoreRef total score reference
 * @returns {void}
 */
export function updateTotalScoreHelper(document: any, totalScoreRef: Ref<number, number>): void {
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
