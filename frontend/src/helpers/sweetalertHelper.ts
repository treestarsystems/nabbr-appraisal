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
