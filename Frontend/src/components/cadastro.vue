<template>
    <div class="cadastro-container">
      <div class="cadastro-box">
        <img src="../assets/Logo-AR.png" alt="Logo" class="cadastro-logo" />
        <form @submit.prevent="cadastrar" class="cadastro-form">
          <div class="form-group">
            <label for="username">Usuário:</label>
            <input type="text" id="username" v-model="username" class="form-input" />
          </div>
          <div class="form-group">
            <label for="email">E-mail:</label>
            <input type="email" id="email" v-model="email" class="form-input" />
          </div>
          <div class="form-group">
            <label for="password">Senha:</label>
            <input type="password" id="password" v-model="password" class="form-input" />
          </div>
          <button type="submit" class="cadastro-button">Cadastrar</button>
        </form>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  
  export default defineComponent({
    name: 'Cadastro',
    data() {
      return {
        username: '',
        email: '',
        password: '',
      };
    },
    methods: {
      cadastrar() {
        // Implemente a lógica de cadastro aqui
        console.log('Username:', this.username);
        console.log('Email:', this.email);
        console.log('Password:', this.password);
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
  .cadastro-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 75vh;
    font-weight: bold;
  }
  
  .cadastro-box {
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
  
  .cadastro-form {
    display: flex;
    flex-direction: column;
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  label {
    margin-bottom: 5px;
    display: block;
    align-items: center;
    
  }
  
  .form-input {
    padding: 10px;
    width: 100%;
    box-sizing: border-box;
    border: 1px solid #ccc;
    border-radius: 3px;
    background-color: rgb(63, 62, 62);
    color: #fff;
  }
  
  .cadastro-button {
    padding: 0.6em 1.2em; 
    background-color:#6c45a0;
    color: #fff;
    border: none;
    font-weight: 500;
    font-size: 1em;
    font-family: inherit;
    border-radius: 8px;
    border: 1px solid transparent;
    cursor: pointer;
  }
  
  .cadastro-logo {
    max-width: 100%;
    margin-bottom: 20px;
    width: 80px;
    height: auto;
  }

  .form-group label{
    color: #fff;
    font-weight: 500;
    font-size: 1em;
  }
  </style>