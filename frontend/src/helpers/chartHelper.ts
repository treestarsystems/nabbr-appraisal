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
