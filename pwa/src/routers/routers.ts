import type { RoutesInterface } from '../interfaces/RoutesInterfaces';
import defaultLayout from '@/layouts/default.vue';

const routes: RoutesInterface = [
  {
    path: '/',
    name: 'home',
    title: '',
    component: () => import('@/views/home.vue'),
    meta: {
      layout: defaultLayout,
    },
  },
];

export default routes;
