export interface FormDataUserBase {
  email: string;
  password: string;
}

export interface FormDataUserContact {
  firstName: string;
  lastName: string;
  phone: string;
}

export interface FormDataUserRegistrationBase extends FormDataUserBase, FormDataUserContact {
  userPrivilegeLevel: string;
  registrationKey: string;
}

export interface FormDataUserPasswordConfirm {
  confirmPassword: string;
}

export interface FormDataUserPasswordReset extends FormDataUserPasswordConfirm {
  password: string;
  oldPassword: string;
}

export interface FormDataUserProfile extends FormDataUserContact, FormDataUserPasswordConfirm {}
export interface FormDataUserRegistrationSubmit extends FormDataUserRegistrationBase, FormDataUserPasswordConfirm {}
