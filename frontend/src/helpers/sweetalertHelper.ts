import { Ref } from 'vue';
import { getAppraisalChartAllHelper, deleteAppraisalChartByIdHelper } from '../helpers/chartHelper';
import { Chart } from '../types/chartTypes';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';

/**
 * @param {any} swal sweetalert2 instance
 * @param {any} message can be err object or string
 * @returns {void}
 */

export function SwalToastSuccessHelper(swal: any, message: any): void {
  swal
    .mixin({
      toast: true,
      position: 'top-right',
      iconColor: 'white',
      customClass: {
        popup: 'colored-toast',
      },
      showConfirmButton: false,
      timer: 2000,
    })
    .fire({
      icon: 'success',
      title: message,
    });
}

/**
 * @param {any} swal sweetalert2 instance
 * @param {any} message can be err object or string
 * @returns {void}
 */

export function SwalToastWarnHelper(swal: any, message: any): void {
  swal
    .mixin({
      toast: true,
      position: 'top-right',
      iconColor: 'white',
      customClass: {
        popup: 'colored-toast',
      },
      showConfirmButton: false,
      timer: 1500,
    })
    .fire({
      icon: 'warning',
      title: message,
    });
}

/**
 * @param {any} swal sweetalert2 instance
 * @param {any} message can be err object or string
 * @returns {void}
 */

export function SwalToastErrorHelper(swal: any, message: any): void {
  swal
    .mixin({
      toast: true,
      position: 'top-right',
      iconColor: 'white',
      customClass: {
        popup: 'colored-toast',
      },
      showConfirmButton: false,
      timer: 5000,
      timerProgressBar: true,
    })
    .fire({
      icon: 'error',
      // TODO: create a better way to catch error obj from different sources like axios and throw
      title: `${message?.response?.data?.message || message?.data?.message || message?.message || message}`,
    });
}

/**
 * @param {any} swal sweetalert2 instance
 * @param {any} message can be err object or string
 * @returns {void}
 */

export async function SwalConfirmationDeleteAppraisalHelper(
  swal: any,
  message: string,
  token: string,
  appraisalId: string,
  chartDataAllRef: Ref<Chart[]>,
): Promise<void> {
  const swalWithBootstrapButtons = swal.mixin({
    customClass: {
      confirmButton: 'btn btn-outline-success m-3',
      cancelButton: 'btn btn-outline-danger m-3',
    },
    buttonsStyling: false,
    iconColor: 'orange',
  });
  swalWithBootstrapButtons
    .fire({
      title: '<i class="" style="color:#a5adbc;">Are you sure?</i>',
      html: `<i class="" style="color:#a5adbc;">${message}</i>`,
      icon: 'warning',
      showCancelButton: true,
      cancelButtonText: 'Delete',
      confirmButtonText: 'Cancel',
      allowOutsideClick: false,
      background: '#353c48',
    })
    .then(async (result: any) => {
      try {
        if (result.isConfirmed) {
          result.dismiss === swal.DismissReason.cancel;
          return;
        }
        const response: ResponseObjectDefaultInterface = await deleteAppraisalChartByIdHelper(token, appraisalId);
        if (response.httpStatus > 299) {
          throw response?.message;
        }
        // Retrieve all appraisal charts and update the chartDataAllRef
        chartDataAllRef.value = (await getAppraisalChartAllHelper(swal, token)) as Chart[];
        swalWithBootstrapButtons.fire({
          title: '<i class="" style="color:#a5adbc;">Appraisal Deleted!</i>',
          icon: 'success',
          iconColor: 'success',
          background: '#353c48',
        });
        return;
      } catch (error) {
        SwalToastErrorHelper(swal, error);
      }
    });
}
