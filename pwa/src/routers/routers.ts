import type { RoutesInterface } from '@/interfaces/RoutesInterfaces';
import defaultLayout from '@/layouts/default.vue';

import Error404 from '@/views/error-404.vue';
import Splash from '@/views/splash.vue';
import Dashboard from '../views/dashboard.vue';
import Login from '@/views/login.vue';
import Profile from '@/views/profile.vue';
import MyDonations from '@/views/my-donations.vue';
import AllDonations from '@/views/all-donations.vue';
import Article from '@/views/article.vue';

const routes: RoutesInterface = [
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: Error404,
    title: '',
  },
  {
    path: '/',
    name: 'splash',
    component: Splash,
    title: '',
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    title: '',
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Dashboard,
    meta: {
      layout: defaultLayout,
    },
    title: '',
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile,
    meta: {
      layout: defaultLayout,
    },
    title: '',
  },
  {
    path: '/my-donations',
    name: 'my-donations',
    component: MyDonations,
    meta: {
      layout: defaultLayout,
    },
    title: '',
  },
  {
    path: '/all-donations',
    name: 'all-donations',
    component: AllDonations,
    meta: {
      layout: defaultLayout,
    },
    title: '',
  },
  {
    path: '/article/:uuid',
    name: 'article',
    component: Article,
    meta: {
      layout: defaultLayout,
    },
    title: '',
  },
];

export default routes;
