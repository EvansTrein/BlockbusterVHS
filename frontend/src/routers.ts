import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";
import Home from "./pages/home.vue";
import NotFound from "./pages/notFound.vue";
import About from "./pages/about.vue";
import Films from "./pages/films.vue";
import Clients from "./pages/clients.vue";
import Rentals from "./pages/rentals.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    component: About,
  },
  {
    path: "/films",
    name: "Films",
    component: Films,
  },
	{
    path: "/clients",
    name: "Clients",
    component: Clients,
  },
	{
    path: "/rentals",
    name: "Rentals",
    component: Rentals,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: NotFound,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
