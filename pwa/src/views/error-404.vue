<script setup lang="ts"></script>

<template>
  <main class="flex flex-col gap-2 items-center justify-center min-h-screen bg-gray-50">
    <div class="text-3xl text-center font-bold text-red-700">404 - Not Found</div>
    <RouterLink
      to="/"
      class="py-2 px-4 rounded hover:bg-blue-100 hover:text-blue-500"
    >
      Home
    </RouterLink>
  </main>
</template>
