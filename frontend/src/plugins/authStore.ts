import { useAuthStore } from '../stores/';

// Setup auth store as a plugin so it can be accessed globally in our FE
const authStorePlugin = {
  install(app: any, option: { pinia: any }) {
    app.config.globalProperties.$store = useAuthStore(option.pinia);
  },
};

export default authStorePlugin;
