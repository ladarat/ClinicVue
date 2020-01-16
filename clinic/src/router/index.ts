import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginPage from '@/pages/login/LoginPage.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'login',
    component: LoginPage
  },
]

const router = new VueRouter({
  routes
})

export default router
