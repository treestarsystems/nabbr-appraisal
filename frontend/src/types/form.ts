export interface formDataRegistrationSubmit {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  password: string;
  userPrivilegeLevel: string;
  registrationKey: string;
}

export interface formDataRegistration extends formDataRegistrationSubmit {
  confirmPassword: string;
}

export interface formDataLoginSubmit {
  email: string;
  password: string;
}
