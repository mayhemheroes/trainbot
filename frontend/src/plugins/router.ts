import { createRouter, createWebHashHistory } from 'vue-router'
import TrainsView from '@/views/TrainsView.vue'
import TrainDetailView from '@/views/TrainDetailView.vue'
import TrainStatsView from '@/views/TrainStatsView.vue'
import TrainsDBProvider from '@/views/TrainsDBProvider.vue'
import NotFound from '@/views/NotFound.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'root',
      redirect: { name: 'trainsList' }
    },
    {
      path: '/trains',
      component: TrainsDBProvider,
      redirect: { name: 'trainsList' },
      children: [
        {
          path: 'list',
          name: 'trainsList',
          component: TrainsView
        },
        {
          path: 'stats',
          name: 'trainStats',
          component: TrainStatsView
        },
        {
          path: ':id',
          name: 'trainDetail',
          component: TrainDetailView
        }
      ]
    },
    { path: '/:pathMatch(.*)*', name: 'notFound', component: NotFound }
  ]
})

export default router
