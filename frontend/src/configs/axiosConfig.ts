import axios from "axios";

import { ACCESS_TOKEN } from "@/constants";

axios.defaults.withCredentials = true;
axios.defaults.baseURL = import.meta.env.VITE_API_URL;

axios.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${localStorage.getItem(ACCESS_TOKEN)}`;

  return config;
});

export { axios };
