<template>
  <div class="wapper">
    <Header />
    <main class="wrapper-content">
      <div class="container">
        <router-view />
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Header from "@/UI/Header.vue";
import Footer from "@/UI/Footer.vue";
</script>

<style scoped></style>
