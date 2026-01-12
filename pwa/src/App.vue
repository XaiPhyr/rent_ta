<script setup lang="ts">
  import { RouterView, useRoute } from 'vue-router';
  import defaultLayout from './layouts/default.vue';

  const route = useRoute();
</script>

<template>
  <component :is="route.meta.layout || defaultLayout">
    <RouterView />
  </component>
</template>

<style scoped></style>
