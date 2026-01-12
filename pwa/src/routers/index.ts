import { createRouter, createWebHistory } from 'vue-router';
import routers from './routers';

const routes: any = routers;
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  console.log('TO: ', to);
  console.log('FROM: ', from);

  next();
});

export default router;
