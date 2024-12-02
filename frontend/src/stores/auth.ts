import { defineStore } from 'pinia';
import axios from 'axios';
import { apiResponseDefault, userState } from '../types/auth';
import { formDataLoginSubmit, formDataRegistrationSubmit } from '../types/form';
import router from '../router';

export const useAuthStore = defineStore('auth', {
  state: (): { user: userState | null } => ({
    user: null,
  }),
  getters: {
    getState(): userState | null {
      return this.user;
    },
    async protectView(): Promise<void> {
      if (this.user === null) await router.back();
    },
  },
  actions: {
    // checkUserIfLoggedIn(): boolean {
    //   if (!this.user) return false;
    //   return true;
    // },
    checkUserPrivilegeLevel(authorizedPrivilegeLevel: string[]): boolean {
      if (!authorizedPrivilegeLevel.includes(this.user?.userPrivilegeLevel || '')) return false;
      return true;
    },
    checkUserIdAuthorized(userId: string): boolean {
      if (this.user?.userId !== userId) return false;
      return true;
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
