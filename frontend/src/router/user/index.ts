export const userRoutes = [
  {
    path: "/users",
    name: "users",
    component: () => import("@/views/users/UsersView.vue"),
  },
  {
    path: "/users/add",
    name: "addUser",
    component: () => import("@/views/users/AddUser/AddUser.vue"),
  },
  {
    path: "/users/edit/:id",
    name: "editUser",
    component: () => import("@/views/users/EditUser/EditUser.vue"),
  },
];
