<template>
	<header class="header">
		<Header />
	</header>
  <main class="main-content">
    <router-view />
  </main>
	<footer class="footer">
		<Footer />
	</footer>
</template>

<script setup lang="ts">
import Header from "@/UI/Header.vue";
import Footer from "@/UI/Footer.vue";
</script>

<style scoped></style>
