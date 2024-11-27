import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";

import "./assets/css/bootstrap-icons.min.css";
import "./assets/css/main.min.css";
import "./assets/css/style.css";
import "./assets/css/OverlayScrollbars.min.css";
import App from "./App.vue";
import Register from "./components/auth/Register.vue";
import Login from "./components/auth/Login.vue";

const router = createRouter({
 history: createWebHistory(),
 routes: [
  { path: "/register", component: Register },
  { path: "/login", component: Login },
 ],
});

createApp(App).use(router).mount("#app");
