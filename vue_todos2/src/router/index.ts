import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import StatsPage from '@/views/StatsPage.vue'
import NotFoundPage from '@/views/NotFoundPage.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/:catchAll(.*)',
    component: NotFoundPage
  },
  {
    path: '/',
    name: 'Home',
    component: HomePage
  },
  {
    path: '/stats',
    name: 'Stats',
    component: StatsPage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
