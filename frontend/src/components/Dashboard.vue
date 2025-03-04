<template>
  <div class="dashboard">
    <h2 class="dashboard__title">Dashboard panel</h2>
    <div class="dashboard__parametrs">
      <div v-for="param in parameters" :key="param.title" class="dashboard__parametrs-item">
        <h3>{{ param.title }}</h3>
        <p>Count: {{ param.count }}</p>
      </div>
      <div class="dashboard__operations">
        <div v-for="oper in operations" :key="oper.title" class="dashboard__operations-item">
          <h3>{{ oper.title }}</h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import { type Parameter, type Operation } from "./../types/types";

const parameters = reactive<Parameter[]>([
  { title: "Total Cassettes on Stock", count: 100 },
  { title: "Cassettes Issued", count: 50 },
  { title: "Total Films", count: 200 },
  { title: "Total Clients", count: 150 },
]);

const operations = reactive<Operation[]>([
  { title: "Add Films", description: "", apiPath: "", method: "POST" },
  { title: "Add Client", description: "", apiPath: "", method: "POST" },
  { title: "Add Rental", description: "", apiPath: "", method: "POST" },
  { title: "Download movies", description: "", apiPath: "", method: "POST" },
  { title: "Clear All databases", description: "", apiPath: "", method: "DELETE" },
]);
</script>

<style scoped lang="scss"></style>
