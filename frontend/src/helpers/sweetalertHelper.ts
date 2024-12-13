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
      confirmButton: 'btn btn-success',
      cancelButton: 'btn btn-danger',
    },
    buttonsStyling: false,
  });
  swalWithBootstrapButtons
    .fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'error',
      showCancelButton: true,
      confirmButtonText: 'Cancel!',
      cancelButtonText: 'Delete',
      reverseButtons: true,
    })
    .then((result: any) => {
      if (result.isConfirmed) {
        swalWithBootstrapButtons.fire({
          title: 'Deleted!',
          text: 'Your file has been deleted.',
          icon: 'success',
        });
      } else if (
        /* Read more about handling dismissals below */
        result.dismiss === swal.DismissReason.cancel
      ) {
        swalWithBootstrapButtons.fire({
          title: 'Cancelled',
          text: 'Your imaginary file is safe :)',
          icon: 'error',
        });
      }
    });
}
