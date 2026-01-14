<script setup lang="ts">
  import { ref } from 'vue';
  import Icons from '@/components/Icons.vue';
  import { useRoute, useRouter } from 'vue-router';

  const route = useRoute();
  const router = useRouter();

  const isVisible = ref(false);

  const onOpenMenu = (payload: boolean) => {
    isVisible.value = payload;

    if (payload) {
      document.body.style.overflow = 'hidden';
      return;
    }

    document.body.style.overflow = '';
  };

  const onRedirect = (name: string) => {
    document.body.style.overflow = '';
    router.push({ name });
  };

  interface RouteItem {
    name: string;
    icon: string;
    title: string;
    callback: () => void;
  }

  interface RouteItems extends Array<RouteItem> {}

  const items: RouteItems = [
    {
      name: 'dashboard',
      icon: 'home',
      title: 'Home',
      callback: () => {
        onRedirect('dashboard');
      },
    },
    {
      name: 'all-donations',
      icon: 'wallet',
      title: 'All Donations',
      callback: () => {
        onRedirect('all-donations');
      },
    },
    {
      name: 'my-donations',
      icon: 'heart',
      title: 'My Donations',
      callback: () => {
        onRedirect('my-donations');
      },
    },
    {
      name: 'profile',
      icon: 'user',
      title: 'Profile',
      callback: () => {
        onRedirect('profile');
      },
    },
  ];
</script>

<template>
  <main>
    <div
      class="relative cursor-pointer active:scale-95"
      @click="onOpenMenu(true)"
    >
      <Icons icon="menu" alt="Menu" title="Menu" />
    </div>

    <div
      class="fixed top-0 left-0 h-full w-[250px] bg-white shadow-lg z-50 transition-transform duration-500 ease-in-out"
      :class="isVisible ? '-translate-x-0' : '-translate-x-full'"
    >
      <div class="h-screen overflow-auto pb-15">
        <div
          class="flex items-center p-4 justify-between sticky top-0 bg-white z-50"
        >
          <div class="text-xl">TITLE</div>
          <div class="active:scale-95" @click="onOpenMenu(false)">
            <Icons
              icon="close"
              alt="Close"
              title="Close"
              class="active:scale-95"
            />
          </div>
        </div>

        <div class="p-4 flex flex-col gap-2">
          <div
            class="cursor-pointer"
            v-for="({ name, icon, title, callback }, index) in items"
            :key="index"
          >
            <div
              class="flex gap-2 items-center p-2"
              :class="{ 'text-indigo-900 font-bold': route.name === name }"
              @click="callback"
              v-if="title"
            >
              <Icons :icon="icon" :alt="title" :title="title" />
              <div class="text-indigo-500">
                {{ title }}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="p-4 bottom-0 left-0 fixed w-full bg-white shadow">
        <div class="px-2">Logout</div>
      </div>
    </div>
  </main>
</template>
