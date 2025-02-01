<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <header>
    <div class="image-container">
      <img class="image" src="@/assets/new_logo.png" alt="Website Logo">
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
  position: fixed;

}

.image-container {
  position: fixed;
  top: 20px;
  left: 50px;
  z-index: 10;
  width: 400px;
  height: 130px;
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
  top: 0;
  left: -5px;
  width: 100%;
  height: 170px;
  background-color: #E5D0AC;
  font-size: 40px;

  display: flex;
  align-items: center;
  justify-content: center;
}

.menu a {
  position: relative;
  left: 70px;
  margin: 30px 25px;
  text-decoration: none;
}


nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
}

nav a.router-link-exact-active {
  color: #85144B;
  font-weight: bold;
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
  border-radius: 20px !important;

}



</style>
