<script setup lang="ts">
  import { useRoute } from 'vue-router';

  const route = useRoute();
</script>

<template>
  <main>
    <component :is="route.meta.layout">
      <div class="text-3xl text-blue-500">My Donations</div>
    </component>
  </main>
</template>
