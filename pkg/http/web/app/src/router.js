import Vue from "vue";
import Router from "vue-router";
import Meta from "vue-meta";
import Home from "./views/Home.vue";
import Lost from "./views/404.vue";

Vue.use(Router);
Vue.use(Meta);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/About.vue")
    },
    {
      path: "/papers",
      name: "papers",
      // route level code-splitting
      // this generates a separate chunk (papers.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "papers" */ "./views/Papers.vue")
    },
    {
      path: "/calendar",
      name: "calendar",
      // route level code-splitting
      // this generates a separate chunk (calendar.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "calendar" */ "./views/Calendar.vue")
    },
    {
      path: "*",
      name: "404",
      component: Lost
    }
  ]
});
