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
  position: absolute;         /* Positions the container of the image absolutely within the header */
  top: 50px;                     /* Aligns it to the top */
  left: 135px;                    /* Aligns it to the left */
  z-index: 10;                /* Ensures the image stays above other content */
}

.image {
  display: block;
  margin: 0;
  width: 40%;
  height: auto;
  border-radius: 15px;
}

.menu {
  width: 100%;

  font-size: 60px;
  text-align: left;
  display: flex;
  flex-direction: column;
  align-items: flex-start;

}

.menu a {
  /* Assuming your menu items are anchor tags */
  margin: 20px 0;            /* Increase the space between menu items (up/down direction) */
  text-decoration: none;     /* Remove underline from links */
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

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
