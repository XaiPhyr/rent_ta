<script setup lang="ts">
  import { useRoute } from 'vue-router';

  const route = useRoute();
</script>

<template>
  <main>
    <component :is="route.meta.layout">
      <div class="">All Donations</div>
    </component>
  </main>
</template>
