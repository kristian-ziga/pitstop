<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <header>
    <div class="image-container">
      <img class="image" src="@/assets/logo_image.png" alt="Website Logo">
    </div>
    <div class="wrapper">
      <nav class="menu" style="margin-left: 5px;">
        <RouterLink to="/">Kto sme</RouterLink>
        <RouterLink to="/about">Hodnotenia</RouterLink>
        <RouterLink to="/test">Kontakt</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.image-container {
  position: absolute;
  top: 40px;
  left: 135px;
  z-index: 10;
  width: 350px;
  height: 200px;
}

.image {
  display: block;
  margin: 0;
  width: 100%;
  height: 100%;
  border-radius: 15px;
}

.menu {
  position: fixed;
  top: 280px;
  left: 130px;
  width: 30%;
  float: left;
  font-size: 60px;
  text-align: left;
  display: flex;
  flex-direction: column;
  align-items: flex-start;

}

.menu a {

  margin: 20px 0;
  text-decoration: none;
}


nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}



</style>
