export interface formDataRegistration {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  password: string;
  userPrivilegeLevel: string;
  registrationKey: string;
}

export interface apiResponseDefault {
  status: string;
  httpStatus: number;
  message: string;
  payload: any[any];
}
