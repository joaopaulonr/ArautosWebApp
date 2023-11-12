import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router';

const app = createApp(App);

// Use o roteador aqui
app.use(router);

// Monte a aplicação no elemento com o ID 'app'
app.mount('#app');
