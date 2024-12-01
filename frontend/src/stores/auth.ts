import { defineStore } from 'pinia';
import axios from 'axios';
import { apiResponseDefault, formDataLogin, formDataRegistration, userState } from '../types/auth';
import router from '../router';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
  }),
  getters: {
    getState(): userState | null {
      return this.user;
    },
  },
  actions: {
    async protectView() {
      console.log(this.user);
      // if (this.user === null) await router.back();
    },
    async checkUserPrivilegeLevel(userPrivilegeLevel: string, authorizedPrivilegeLevel: string[]) {
      console.log(userPrivilegeLevel);
      if (!authorizedPrivilegeLevel.includes(userPrivilegeLevel)) {
        // await router.back();
      }
    },
    async register(userRegistrationFormData: formDataRegistration): Promise<apiResponseDefault> {
      const apiResponse: apiResponseDefault = (await axios.post('/api/v1/auth/signup', userRegistrationFormData)).data;
      return apiResponse;
    },
    async login(userLoginFormData: formDataLogin): Promise<apiResponseDefault> {
      const apiResponse: apiResponseDefault = (await axios.post('/api/v1/auth/login', userLoginFormData)).data;
      console.log(apiResponse);
      // console.log(apiResponse.payload[0])
      // this.user = apiResponse.payload[0];
      return apiResponse;
    },
    async logout() {
      this.user = null;
      await router.push('/login');
    },
  },
  persist: true,
});
