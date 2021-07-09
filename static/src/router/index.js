import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  {
    path: "/assignments",
    name: "Assignments",
    component: () => 
    import(/* webpackChunkName: "assignments" */ "../views/Assignments.vue"),
  },
  {
    path: "/kid",
    name: "Kid",
    component: () => 
    import(/* webpackChunkName: "kid" */ "../views/Kid.vue"),
  },
  {
    path: "/settings",
    name: "settings",
    component: () => 
    import(/* webpackChunkName: "settings" */ "../views/Settings.vue"),
  },
];

const router = new VueRouter({
  routes,
});

export default router;
