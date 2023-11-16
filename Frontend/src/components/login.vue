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
          <p class="signup-link">Sem login? <router-link to="/cadastro">Cadastre-se aqui!</router-link></p>
        </form>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  import { useRouter } from 'vue-router';

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
      redirectToCadastro() {
      const router = useRouter();
      router.push('/cadastro');
    },
      
    },
  });
  </script>
  
  <style scoped>
    body {
    margin: 0;
    display: flex;
    place-items: center;
    min-width: 320px;
    min-height: 100vh;
}

.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 75vh;
    font-weight: bold;
}

.login-box {
    display: flex;
    flex-direction: column;
    justify-content: center;
    text-align: center;
    align-items: center;
    background-color: rgb(63, 62, 62);
    padding: 20px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    width: 300px;
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
    background-color: rgb(63, 62, 62);
}

.login-button {
    padding: 0.6em 1.2em;
    background-color: #6c45a0;
    color: #fff;
    border: none;
    font-weight: 500;
    font-size: 1em;
    font-family: inherit;
    border-radius: 8px;
    border: 1px solid transparent;
    cursor: pointer;
}

.login-logo {
    max-width: 100%;
    /* Garante que a imagem não ultrapasse a largura da div pai */
    margin-bottom: 20px;
    width: 80px;
}
.form-group label{
    color: #fff;
    font-weight: 500;
    font-size: 1em;
}
.signup-link {
    color: white;
    font-weight: 500;
    font-size: 1em;
}
  </style>