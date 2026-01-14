<script setup lang="ts">
  import { ref } from 'vue';
  import Icons from '@/components/Icons.vue';

  const isVisible = ref(false);
</script>

<template>
  <main>
    <div class="relative cursor-pointer active:scale-95" @click="isVisible = true">
      <Icons icon="bell" alt="Notification" title="Notification" />
      <span
        class="absolute top-0 right-0 inline-flex h-2 w-2 rounded-full bg-indigo-500 ring-2 ring-white"
      ></span>
      <span
        class="absolute top-0 right-0 inline-flex h-2 w-2 rounded-full bg-indigo-500 ring-2 ring-white animate-ping"
      ></span>
    </div>

    <div
      class="fixed top-0 right-0 h-full w-full bg-white shadow-lg z-50 transition-transform duration-500 ease-in-out"
      :class="isVisible ? 'translate-x-0' : 'translate-x-full'"
    >
      <div class="h-full overflow-auto">
        <div
          class="flex items-center p-4 justify-between sticky top-0 bg-white z-50"
        >
          <div class="text-xl">Notifications</div>
          <div class="active:scale-95" @click="isVisible = false">
            <Icons icon="close" alt="Close" title="Close" class='active:scale-95' />
          </div>
        </div>

        <div class="px-4 pb-4 flex flex-col gap-4">
          <div class="" v-for="(item, index) in 30" :key="index">
            <div class="font-bold">Header</div>
            <div class="flex gap-2 items-center justify-between">
              <div class="text-sm">
                Lorem ipsum dolor sit, amet consectetur adipisicing elit.
                Doloremque saepe placeat iure nobis.
              </div>
              <div class="grow flex justify-end">
                <div
                  class="h-2 w-2 rounded-full bg-indigo-500 ring-2 ring-white"
                ></div>
              </div>
            </div>
            <div
              class="text-xs text-gray-500 flex items-center justify-between"
            >
              <div class="">Activity</div>
              <div class="">5 mins ago</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
