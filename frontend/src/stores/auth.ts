import { defineStore } from 'pinia';
import axios from 'axios';
import { apiResponseDefault, userState } from '../types/auth';
import { formDataLoginSubmit, formDataRegistrationSubmit } from '../types/form';
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
      console.log('protectedView', this.getState);
      // if (this.user === null) await router.back();
    },
    async checkUserPrivilegeLevel(userPrivilegeLevel: string, authorizedPrivilegeLevel: string[]) {
      console.log(userPrivilegeLevel);
      if (!authorizedPrivilegeLevel.includes(userPrivilegeLevel)) {
        // await router.back();
      }
    },
    async register(userRegistrationFormData: formDataRegistrationSubmit): Promise<apiResponseDefault> {
      const apiResponse: apiResponseDefault = (await axios.post('/api/v1/auth/signup', userRegistrationFormData)).data;
      return apiResponse;
    },
    async login(userLoginFormData: formDataLoginSubmit): Promise<apiResponseDefault> {
      try {
        const apiResponse: apiResponseDefault = (await axios.post('/api/v1/auth/login', userLoginFormData)).data;
        if (apiResponse.httpStatus < 300) {
          this.user = apiResponse.payload[0];
        }
        return apiResponse;
      } catch (err: any) {
        return {
          status: 'failure',
          httpStatus: 400,
          message: err,
          payload: [],
        };
      }
    },
    async logout() {
      this.user = null;
      await router.push('/login');
    },
  },
  persist: true,
});
