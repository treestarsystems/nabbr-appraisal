export function SwalToastSuccess(swal: any, message: any): void {
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
      icon: 'success',
      title: message,
    });
}

export function SwalToastError(swal: any, message: any): void {
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
