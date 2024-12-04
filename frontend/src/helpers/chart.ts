import axios from 'axios';
import { SwalToastError } from '../helpers/sweetalert';
import { ResponseObjectDefaultInterface } from '../types/model';
import { Chart } from '../types/chart';
export async function getAppraisalChartTemplate(swal: any, token: string): Promise<void | Chart> {
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
      SwalToastError(swal, response?.message);
      return;
    }
    const chartTemplate: Chart = response.payload[0];
    return chartTemplate;
  } catch (err) {
    SwalToastError(swal, err);
    return;
  }
}
