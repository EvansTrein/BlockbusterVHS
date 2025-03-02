<template>
  <div class="form">
    <form class="form-content" @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="name">Name:</label>
        <input type="text" id="name" v-model="formData.name" />
      </div>
      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="formData.email" />
      </div>
      <div class="form-group">
        <label for="phone">Phone:</label>
        <input
          type="tel"
          id="phone"
          placeholder="+7"
          v-model="formData.phone"
        />
      </div>
      <div class="form-group">
        <label for="password">Password:</label>
        <input
          type="password"
          id="password"
          placeholder="min 8 symbols"
          v-model="formData.password"
        />
      </div>
      <button type="submit">Send</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import ApiService from "./../services/api";

interface FormData {
  name: string;
  email: string;
  phone: string;
  password: string;
}

const formData = ref<FormData>({
  name: "",
  email: "",
  phone: "",
  password: "",
});

const handleSubmit = async () => {
  try {
    const response = await ApiService.register(formData.value);

    if (response.status === 201) {
      alert(`successful registration, user id - ${response.data.user_id}`);
    } else {
      alert(`fail register  ${response.data.message}\n\n ${response.data.error}`);
    }
  } catch (error: any) {
    alert(`registration failed: ${error.message}`);
  }
};
</script>

<style scoped lang="scss"></style>
