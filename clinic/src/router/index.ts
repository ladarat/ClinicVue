import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginPage from '@/pages/login/LoginPage.vue'
import MenuPage from '@/pages/menu/MenuPage.vue'

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
  {
    path: '/menu',
    name: 'menu',
    component: MenuPage
    // component: () => import('@/pages/menu/MenuPage.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
