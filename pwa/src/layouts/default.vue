<script setup lang="ts">
  import { onMounted, onBeforeUnmount } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import Icons from '@/components/Icons.vue';
  import Notifications from '@/components/Notifications.vue';
  import Navigations from '@/components/Navigations.vue';
  import type { RoutesInterface } from '@/interfaces/RoutesInterfaces';

  const router = useRouter();
  const route = useRoute();

  const items: RoutesInterface = [
    {
      path: '',
      name: 'dashboard',
      icon: 'home',
      title: 'Home',
      callback: () => {
        router.push({ name: 'dashboard' });
      },
    },
    {
      path: '',
      name: 'all-donations',
      icon: 'wallet',
      title: 'All Donations',
      callback: () => {
        router.push({ name: 'all-donations' });
      },
    },
    {
      path: '',
      name: 'my-donations',
      icon: 'heart',
      title: 'My Donations',
      callback: () => {
        router.push({ name: 'my-donations' });
      },
    },
    {
      path: '',
      name: 'profile',
      icon: 'user',
      title: 'Profile',
      callback: () => {
        router.push({ name: 'profile' });
      },
    },
    {
      path: '',
      name: 'court-dashboard',
      icon: 'court',
      title: 'Courts',
      callback: () => {
        router.push({ name: 'court-dashboard' });
      },
    },
  ];

  function setViewportHeight() {
    const height = window.visualViewport?.height ?? window.innerHeight;
    const vh = height * 0.01;
    document.documentElement.style.setProperty('--vh', `${vh}px`);
  }

  onMounted(() => {
    setViewportHeight();

    // Update when viewport changes (works on Samsung)
    window.visualViewport?.addEventListener('resize', setViewportHeight);
    window.visualViewport?.addEventListener('scroll', setViewportHeight);

    // Fallback listeners
    window.addEventListener('resize', setViewportHeight);
    window.addEventListener('orientationchange', setViewportHeight);
  });

  onBeforeUnmount(() => {
    window.visualViewport?.removeEventListener('resize', setViewportHeight);
    window.visualViewport?.removeEventListener('scroll', setViewportHeight);

    window.removeEventListener('resize', setViewportHeight);
    window.removeEventListener('orientationchange', setViewportHeight);
  });
</script>

<template>
  <div class="flex flex-col min-h-[calc(var(--vh,1vh)*100)] bg-gray-100">
    <header class="p-4 flex justify-between">
      <Navigations />
      <Notifications />
    </header>

    <main class="flex-grow p-6">
      <slot />
    </main>

    <footer
      class="w-full bg-white shadow rounded-t-xl text-center text-xs p-2 sticky bottom-0"
    >
      <div class="flex gap-2 items-center justify-around">
        <div
          class="cursor-pointer"
          v-for="({ name, icon, title, callback }, index) in items"
          :key="index"
        >
          <div
            class="flex flex-col items-center"
            :class="{ 'text-indigo-900 font-bold': route.name === name }"
            @click="callback"
            v-if="title"
          >
            <Icons :icon="icon" :alt="title" :title="title" />
            <div class="text-xs text-indigo-500">
              {{ title }}
            </div>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>
