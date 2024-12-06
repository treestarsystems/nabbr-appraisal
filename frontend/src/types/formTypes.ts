export interface FormDataRegistrationSubmit {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  password: string;
  userPrivilegeLevel: string;
  registrationKey: string;
}

export interface FormDataRegistration extends FormDataRegistrationSubmit {
  confirmPassword: string;
}

export interface FormDataLoginSubmit {
  email: string;
  password: string;
}
