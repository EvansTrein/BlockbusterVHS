<template>
  <nav class="navbar">
    <header class="container">
      <ul class="navbar-content">
        <li
          v-for="link in links"
          :key="link.alias"
          :class="{ active: activeLink === link.alias }"
          @click="setActiveLink(link.alias)"
        >
          <router-link :to="link.url">{{ link.title }}</router-link>
        </li>
      </ul>
    </header>
  </nav>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { RouterLink, useRoute } from "vue-router";

const route = useRoute();
const activeLink = ref<string | null>(null);

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

const setActiveLink = (alias: string | null) => {
  activeLink.value = alias;
};

const updateActiveLink = () => {
  const currentPath = route.path;
  const activeLinkItem = links.find((link) => link.url === currentPath);
  if (activeLinkItem) {
    setActiveLink(activeLinkItem.alias);
  } else {
    setActiveLink(null);
  }
};

updateActiveLink();

watch(() => route.path, updateActiveLink);
</script>

<style scoped></style>
