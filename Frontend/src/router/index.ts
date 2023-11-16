import {
  createRouter,
  createWebHistory
} from 'vue-router';
import home from "../views/home.vue";
import singIn from "../views/singIn.vue";
import singUp from "../views/singUp.vue";

const routes = [{
      path: '/',
      component: home
  },
  {
      path: '/login',
      component: singIn
  },
  {
      path: '/cadastro',
      component: singUp
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;