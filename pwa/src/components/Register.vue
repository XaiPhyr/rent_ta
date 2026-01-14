<script setup lang="ts">
  import { ref, type Ref } from 'vue';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const showPassword: Ref<boolean> = ref(false);
  let isLoading: Ref<boolean> = ref(false);

  const onSubmit = () => {
    isLoading.value = true;

    setTimeout(() => {
      isLoading.value = false;
      router.push({ name: 'dashboard' });
    }, 2000);
  };
</script>

<template>
  <main>
    <form class="flex flex-col gap-6 text-center">
      <div class="w-[350px]">
        <input
          id="register-username"
          name="register-username"
          class="w-full pr-10 pl-4 py-3 border border-gray-400 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
          placeholder="Username"
          autocomplete="username"
          required
        />
      </div>
      <div class="relative w-[350px]">
        <input
          id="register-password"
          name="register-password"
          :type="showPassword ? 'text' : 'password'"
          class="w-full pr-10 pl-4 py-3 border border-gray-400 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
          placeholder="Password"
          autocomplete="current-password"
          required
          maxlength="16"
        />
        <span
          class="text-xs absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-500 hover:cursor-pointer active:cursor-pointer active:text-gray-700"
          @click="showPassword = !showPassword"
        >
          {{ showPassword ? 'hide' : 'show' }}
        </span>
      </div>
      <div class="relative w-[350px]">
        <input
          id="register-confirm-password"
          name="register-confirm-password"
          :type="showPassword ? 'text' : 'password'"
          class="w-full pr-10 pl-4 py-3 border border-gray-400 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500"
          placeholder="Confirm Password"
          autocomplete="current-password"
          required
          maxlength="16"
        />
        <span
          class="text-xs absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-500 hover:cursor-pointer active:cursor-pointer active:text-gray-700"
          @click="showPassword = !showPassword"
        >
          {{ showPassword ? 'hide' : 'show' }}
        </span>
      </div>
      <div class="">
        <button
          :disabled="isLoading"
          type="button"
          class="w-full py-3 bg-indigo-500 text-white rounded shadow-md hover:cursor-pointer hover:bg-indigo-600 active:bg-indigo-600 active:scale-102 disabled:scale-100 disabled:opacity-25 disabled:cursor-not-allowed"
          @click="onSubmit"
        >
          Sign Up
        </button>
      </div>
    </form>
  </main>
</template>
