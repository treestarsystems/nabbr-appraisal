import { defineStore } from 'pinia';
import axios from 'axios';
import { jwtDecode, JwtPayload } from 'jwt-decode';
import { UserState } from '../types/authTypes';
import { ResponseObjectDefaultInterface } from '../types/generalTypes';
import { FormDataUserBase, FormDataUserRegistrationBase } from '../types/formTypes';
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
    /**
     * Check if user has the required privilege level.
     * @param {string[]} authorizedPrivilegeLevel array of authorized privilege levels. Ex: ['ADMIN', 'APPRAISER']
     * @returns {boolean}
     */

    checkUserPrivilegeLevel(authorizedPrivilegeLevel: string[]): boolean {
      if (!authorizedPrivilegeLevel.includes(this.user?.userPrivilegeLevel || '')) return false;
      return true;
    },
    /**
     * Check if user has the required privilege level.
     * @param {string} path path of the current route.
     * @returns {boolean}
     */

    checkUserPrivilegeLevelAuthorized(path: string): boolean {
      let authorizedPrivilegeLevels: string[];
      const basePath = path.split('/')[1];
      // If not apart of privileged users then push to unauth page.
      switch (basePath) {
        case 'dashboard':
        case 'appraisal':
          authorizedPrivilegeLevels = ['ADMIN', 'APPRAISER'];
          break;
        case 'user':
          authorizedPrivilegeLevels = ['ADMIN', 'APPRAISER', 'PETOWNER'];
          break;
        default:
          authorizedPrivilegeLevels = [''];
      }
      return this.checkUserPrivilegeLevel(authorizedPrivilegeLevels);
    },
    /**
     * Check if user has the required privilege level and redirect to the /unauthorized view if unauthorized.
     * @param {string} path path of the current route.
     * @returns {Promise<void>}
     */

    async checkUserPrivilegeLevelAuthorizedThenRedirect(path: string): Promise<void> {
      try {
        const isAuthorized = this.checkUserPrivilegeLevelAuthorized(path);
        if (!isAuthorized) {
          await router.push('/unauthorized');
        }
      } catch (err: any) {
        console.error(err);
      }
    },
    /**
     * Check if the userId is authorized to view the page if not redirect to /unauthorized. Usually used in the user profile page.
     * @param path path of the current route.
     * @returns {Promise<void>}
     */

    async checkUserIdAuthorized(path: string): Promise<void> {
      try {
        const userIdFromPath = path.split('/')[2];
        if (this.user?.userId !== userIdFromPath) {
          await router.push('/unauthorized');
        }
      } catch (err: any) {
        console.error(err);
      }
    },
    /**
     * Check if the token is expired or invalid and redirect to /login if expired or invalid.
     * @returns {Promise<void>}
     */

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
    /**
     * Register a new user.
     * @param {FormDataUserRegistrationBase} userRegistrationFormData
     * @returns {Promise<ResponseObjectDefaultInterface>}
     */

    async register(userRegistrationFormData: FormDataUserRegistrationBase): Promise<ResponseObjectDefaultInterface> {
      try {
        const apiResponse: ResponseObjectDefaultInterface = (
          await axios.post('/api/v1/auth/signup', userRegistrationFormData)
        ).data;
        return apiResponse;
      } catch (err: any) {
        console.error(err);
        const errResponse: ResponseObjectDefaultInterface = {
          status: 'failure',
          httpStatus: 400,
          message: err,
          payload: [],
        };
        return errResponse;
      }
    },
    /**
     * Login a user.
     * @param {FormDataUserBase} userLoginFormData
     * @returns {Promise<ResponseObjectDefaultInterface>}
     */

    async login(userLoginFormData: FormDataUserBase): Promise<ResponseObjectDefaultInterface> {
      try {
        const apiResponse: ResponseObjectDefaultInterface = (await axios.post('/api/v1/auth/login', userLoginFormData))
          .data;
        if (apiResponse.httpStatus < 300) {
          this.user = apiResponse.payload[0];
        }
        return apiResponse;
      } catch (err: any) {
        console.error(err);
        const errResponse: ResponseObjectDefaultInterface = {
          status: 'failure',
          httpStatus: 400,
          message: err,
          payload: [],
        };
        return errResponse;
      }
    },
    /**
     * Logout a user.
     * @returns {Promise<void>}
     */

    async logout(): Promise<void> {
      try {
        this.user = null;
        await router.push('/login');
      } catch (err: any) {
        console.error(err);
        // This seems redundant but safe.
        await router.push('/login');
      }
    },
  },
  persist: true,
});
