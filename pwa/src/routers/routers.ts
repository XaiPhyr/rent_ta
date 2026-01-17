import type { RoutesInterface } from '@/interfaces/RoutesInterfaces';
import defaultLayout from '@/layouts/default.vue';

import Error404 from '@/views/error-404.vue';
import Splash from '@/views/splash.vue';
import Dashboard from '../views/dashboard.vue';
import Login from '@/views/login.vue';
import Profile from '@/views/profile.vue';
import Article from '@/views/article.vue';
import AddCourt from '@/views/court-wizard/add-court.vue';
import CourtDashboard from '@/views/court/dashboard.vue';

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
    path: '/article/:uuid',
    name: 'article',
    component: Article,
    meta: {
      layout: defaultLayout,
    },
    title: '',
  },
  {
    path: '/court-wizard/add',
    name: 'add-court',
    component: AddCourt,
    meta: {
      layout: defaultLayout, // Or potentially a different layout if needed
    },
    title: 'Add New Court',
  },
  {
    path: '/court/dashboard',
    name: 'court-dashboard',
    component: CourtDashboard,
    meta: {
      layout: defaultLayout,
    },
    title: 'Court Dashboard',
  },
];

export default routes;
