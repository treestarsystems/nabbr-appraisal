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

export function SwalConfirmationDeleteHelper(swal: any, message: any): void {
  const swalWithBootstrapButtons = swal.mixin({
    customClass: {
      confirmButton: 'btn btn-success m-3',
      cancelButton: 'btn btn-danger m-3',
    },
    buttonsStyling: false,
  });
  swalWithBootstrapButtons
    .fire({
      title: 'Are you sure?',
      text: message,
      icon: 'warning',
      showCancelButton: true,
      cancelButtonText: 'Delete',
      confirmButtonText: 'Cancel',
      allowOutsideClick: false,
    })
    .then((result: any) => {
      if (result.isConfirmed) {
        result.dismiss === swal.DismissReason.cancel;
        return;
      }
      swalWithBootstrapButtons.fire({
        title: 'Appraisal Deleted!',
        icon: 'success',
      });
      return;
    });
}
