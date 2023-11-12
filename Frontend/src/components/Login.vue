<template>
  <div class="login-container">
    <div class="login-box">
      <img src="../assets/Logo-AR.png" alt="Logo" class="login-logo" />
      <form @submit.prevent="login" class="login-form">
        <div class="form-group">
          <label for="username">Usuário:</label>
          <input type="text" id="username" v-model="username" class="form-input" />
        </div>
        <div class="form-group">
          <label for="password">Senha:</label>
          <input type="password" id="password" v-model="password" class="form-input" />
        </div>
        <button type="submit" class="login-button">Login</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Login',
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    login() {
      // Implemente a lógica de autenticação aqui
      console.log('Username:', this.username);
      console.log('Password:', this.password);
    },
  },
});
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-box {
  background-color: rgb(63, 62, 62);
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 300px;
}

.login-title {
  text-align: center;
  font-size: 1.5rem;
  margin-bottom: 20px;
}

.login-form {
  display: flex;
  flex-direction: column;
}

.form-group {
  margin-bottom: 15px;
}

label {
  margin-bottom: 5px;
  display: block;
}

.form-input {
  padding: 10px;
  width: 100%;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 3px;
}

.login-button {
  padding: 10px;
  background-color: rgb(85, 58, 150);
  color: #fff;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #6c45a0;
}

.login-logo {
  max-width: 100%; /* Garante que a imagem não ultrapasse a largura da div pai */
  margin-bottom: 20px;
  width: 80px;
  height: auto;
}

</style>
