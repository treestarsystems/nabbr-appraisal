import { defineStore } from 'pinia';
import axios from 'axios';
import { jwtDecode, JwtPayload } from 'jwt-decode';
import { UserState } from '../types/authTypes';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { FormDataUserBase, FormDataUserRegistrationSubmit, FormDataUserRegistrationBase } from '../types/formTypes';
import router from '../router';

export const useAuthStore = defineStore('auth', {
  state: (): { user: UserState | null } => ({
    user: null,
  }),
  getters: {
    getState(): UserState | null {
      return this.user;
    },
  },
  actions: {
    checkUserPrivilegeLevel(authorizedPrivilegeLevel: string[]): boolean {
      if (!authorizedPrivilegeLevel.includes(this.user?.userPrivilegeLevel || '')) return false;
      return true;
    },
    checkUserIdAuthorized(userId: string): boolean {
      if (this.user?.userId !== userId) return false;
      return true;
    },
    async checkTokenExpired() {
      const errMsg = 'Invalid/Expired Token. Please Login.';
      try {
        const now = Date.now() / 1000;
        // Decode token for the exp property
        const decodedToken = jwtDecode<JwtPayload>(this.user?.token || '');
        // Guards to return early.
        if (!decodedToken?.exp) throw '';
        if (now > decodedToken.exp) throw '';
        // Should return so it can execute next()
        return;
      } catch (err: any) {
        // Ternary: I just wanted a nicely formatted error incase the jwtDecode returns an error.
        console.error(`${errMsg} ${err?.message ? '- ' + err.message : ''}`);
        this.user = null;
        await router.push('/login');
        //TODO: Do I need this return?
        // return;
      }
    },
    async register(userRegistrationFormData: FormDataUserRegistrationBase): Promise<ResponseObjectDefaultInterface> {
      const apiResponse: ResponseObjectDefaultInterface = (
        await axios.post('/api/v1/auth/signup', userRegistrationFormData)
      ).data;
      return apiResponse;
    },
    async login(userLoginFormData: FormDataUserBase): Promise<ResponseObjectDefaultInterface> {
      try {
        const apiResponse: ResponseObjectDefaultInterface = (await axios.post('/api/v1/auth/login', userLoginFormData))
          .data;
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
