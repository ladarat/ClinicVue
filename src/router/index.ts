import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../pages/HomePage.vue';
import Login from '../pages/LoginPage.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  }
];

const router = new VueRouter({
  routes,
});

export default router;
