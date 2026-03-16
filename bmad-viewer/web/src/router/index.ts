import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/flow/:role',
      name: 'flow',
      component: () => import('@/views/FlowView.vue'),
    },
    {
      path: '/doc/:path(.*)',
      name: 'doc',
      component: () => import('@/views/DocView.vue'),
    },
  ],
})

export default router
