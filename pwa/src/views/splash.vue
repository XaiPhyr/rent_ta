<script setup lang="ts">
  import { useRouter } from 'vue-router';

  const router = useRouter();

  setTimeout(() => {
    router.push({ name: 'dashboard' });
  }, 2000);

  const title = 'Rent Ta'.split('');
</script>

<template>
  <main
    class="flex flex-col gap-2 items-center justify-center min-h-screen bg-gray-50"
  >
    <div class="text-3xl text-center font-bold flex space-x-2 text-indigo-500">
      <span
        v-for="(letter, index) in title"
        :key="index"
        class="animate-bounce inline-block"
        :style="{ animationDelay: `${index * 0.1}s` }"
      >
        {{ letter }}
      </span>
    </div>
  </main>
</template>
