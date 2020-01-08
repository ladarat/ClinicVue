import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/layouts/Layout.vue'
import Menu from '@/pages/menu/Menu.vue'
import Login from '@/pages/login/Login.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/app/menu',
  },
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/app',
    component: Layout,
    children: [
      {
        path: 'menu',
        component: Menu
      }
    ]
  },
]

const router = new VueRouter({
  routes
})

export default router
