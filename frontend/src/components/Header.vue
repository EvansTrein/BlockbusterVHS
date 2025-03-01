<template>
  <nav class="navbar">
    <div class="container">
      <ul class="navbar-content">
        <li v-for="link in links" :key="link.alias">
          <router-link :to="link.url">{{ link.title }}</router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";

const links = [
  {
    title: "Home",
    alias: "home",
    url: "/",
  },
  {
    title: "Films",
    alias: "films",
    url: "/films",
  },
  {
    title: "Clients",
    alias: "clients",
    url: "/clients",
  },
  {
    title: "Rentals",
    alias: "rentals",
    url: "/rentals",
  },
  {
    title: "About",
    alias: "about",
    url: "/about",
  },
];
</script>

<style scoped></style>
