import { defineStore } from 'pinia';

export const useAuthStore = defineStore('storeAuth', {
  state: () => {
    return {
      authenticated: false,
      user: {},
      test: false,
    };
  },
  getters: {},
  actions: {
    testAction() {
      console.log(this.user);
    },
  },
});
