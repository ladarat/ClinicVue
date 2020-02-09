import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginPage from '@/pages/login/LoginPage.vue'
import MenuPage from '@/pages/menu/MenuPage.vue'
import store from '@/store'
import { IS_AUTH } from '@/store/modules/auth/getters.type'

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
    component: () => import('@/pages/menu/MenuPage.vue')
  },
  {
    path: '/patient/search',
    name: 'patient_search',
    component: () => import('@/pages/patient/search/PatientSearchPage.vue')
  },
  {
    path: '/patient/create',
    name: 'patient_create',
    component: () => import('@/pages/patient/create/CreatePatientPage.vue')
  }
]

const router = new VueRouter({
  routes
})


router.beforeEach((to, from, next) => {
  let isAuthen = store.getters[IS_AUTH]
  if(to.name == 'login' || isAuthen) {
    next()
  } else {
    next({ name: 'login'})
  }
})

export default router
