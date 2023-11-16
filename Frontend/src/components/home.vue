<template>
    <header>
        <nav class="navbar">
          <div class="navbar-left">
            <!-- Espaço para o Logo -->
            <img src="../assets/Logo-AR.png" alt="Logo" class="logo" />
            <!-- Link para Publicações -->
            <router-link to="/publicacoes" class="nav-link">Publicações</router-link>
          </div>
          <div class="navbar-right">
            <!-- Barra de Pesquisa -->
            <input type="text" v-model="searchQuery" placeholder="Pesquisar" class="search-bar" />
            <!-- Espaço para o Ícone de Perfil -->
            <img src="../assets/user.jpeg" alt="Perfil" class="profile-icon" />
          </div>
        </nav>
    </header>
    <body>
        <div id="content">

        </div>
    </body>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  
  export default defineComponent({
    name: 'Home',
    data() {
      return {
        searchQuery: '',
      };
    },
  });
  </script>
  
  <style scoped>
  /* Estilos para a Página Inicial */
  .home-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 0;
  }
  
  /* Estilos para o Header e Barra de Navegação */
  header {
    width: 100%;
  }
  
  .navbar {
    display: flex;
    justify-content: space-between;
    background-color: #6c45a0; /* Cor de fundo do TabNews */
    color: white;
    padding: 10px;
  }
  
  .navbar-left {
    display: flex;
    align-items: center;
  }
  
  .logo {
    height: 35px;
    width: auto;
    margin-right: 10px;
  }
  
  .nav-link {
    text-decoration: none;
    color: white;
    margin-right: 20px;
    font-weight: bold;
    font-size: 16px;
  }
  
  .nav-link:hover {
    text-decoration: underline;
  }
  
  .navbar-right {
    display: flex;
    align-items: center;
  }
  
  .search-bar {
    padding: 8px;
    margin-right: 10px;
    border: none;
    border-radius: 5px;
    font-size: 14px;
  }
  
  .profile-icon {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    cursor: pointer;
  }
  
  /* Estilos para o Conteúdo */
  .content {
    margin-top: 20px;
    /* Adicione estilos adicionais para o conteúdo da página */
  }
  </style>
  