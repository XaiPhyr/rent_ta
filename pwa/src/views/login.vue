<script setup lang="ts">
  import Login from '@/components/Login.vue';
  import Register from '@/components/Register.vue';
  import { ref, type Ref } from 'vue';

  const showLoginForm: Ref<boolean> = ref(true);

  const onRegister = () => {
    showLoginForm.value = !showLoginForm.value;
  };
</script>

<template>
  <main class="flex flex-col min-h-screen bg-gray-50">
    <div class="flex-grow flex flex-col items-center justify-center gap-4">
      <div class="text-3xl text-center font-bold text-indigo-500">
        Rent Ta
      </div>

      <template v-if="showLoginForm">
        <Login />

        <div class="text-xs">or sign in with</div>
      </template>

      <template v-else>
        <Register />
      </template>
    </div>

    <div class="text-center text-sm text-gray-500 pb-3">
      {{
        showLoginForm ? "Don't have any account?" : 'Already have an account?'
      }}
      <span
        class="text-indigo-500 hover:underline hover:cursor-pointer"
        @click="onRegister"
      >
        {{ showLoginForm ? 'Sign Up' : 'Sign In' }}
      </span>
    </div>
  </main>
</template>
