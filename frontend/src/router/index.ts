import { createRouter, createWebHistory } from "vue-router";

import { userRoutes } from "@/router/user";

import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../views/Auth/LoginView.vue"),
    },
    ...userRoutes,
    {
      path: "/:pathMatch(.*)*",
      name: "NotFound",
      component: () => import("../views/NotFoundView/NotFoundView.vue"),
    },
  ],
});

export default router;
