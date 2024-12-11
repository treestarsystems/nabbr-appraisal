export interface FormDataUserBase {
  email: string;
  password: string;
}

export interface FormDataUserInfo extends FormDataUserBase {
  firstName: string;
  lastName: string;
  phone: string;
}

export interface FormDataUserRegistrationBase extends FormDataUserInfo {
  userPrivilegeLevel: string;
  registrationKey: string;
}

export interface FormDataUserPasswordConfirm {
  confirmPassword: string;
}

export interface FormDataPasswordReset extends FormDataUserPasswordConfirm {
  password: string;
  oldPassword: string;
}

export interface FormDataUserRegistrationSubmit
  extends FormDataUserInfo,
    FormDataUserRegistrationBase,
    FormDataUserPasswordConfirm {}
