<script setup lang="ts">
  import { ref } from 'vue';
  import Icons from './Icons.vue';

  const props = defineProps({
    image: { type: String, required: true, default: '' },
    content: { type: String, required: true, default: '' },
    title: { type: String, required: true, default: '' },
    published_by: { type: String, required: true, default: '' },
  });

  const quickView = ref(false);

  const onOpenQuickView = (payload: boolean) => {
    quickView.value = payload;

    if (payload) {
      document.body.style.overflow = 'hidden';
    } else {
      document.body.style.overflow = '';
    }
  };
</script>

<template>
  <main>
    <div
      class="w-full h-auto shadow bg-white rounded-xl p-4 flex flex-col gap-2"
    >
      <div
        class="flex items-center justify-between"
        @click="onOpenQuickView(true)"
      >
        <div class="text-sm">By: {{ published_by }}</div>
        <Icons
          icon="arrows-pointing-out"
          alt="Quick View"
          title="Quick View"
          class="active:scale-95"
        />
      </div>

      <div class="flex flex-col gap-2 h-auto">
        <div class="relative" v-if="image">
          <img :src="image" alt="" srcset="" class="rounded-xl w-full h-full" />
          <div
            class="absolute bottom-0 left-0 inline-flex p-4 font-semibold text-sm"
          >
            {{ title }}
          </div>
        </div>

        <div class="font-semibold text-sm" v-else>
          {{ title }}
        </div>

        <div class="flex items-center">
          <div class="flex-1">
            <div class="flex items-center gap-1 text-sm">
              <Icons icon="clock" class="w-4 h-4 flex items-center" />
              <div class="">20 days</div>
            </div>
          </div>

          <div class="flex-1">
            <div class="flex items-center gap-1 text-sm">
              <Icons icon="pin" class="w-4 h-4 flex items-center" />
              <div class="">Location</div>
            </div>
          </div>
        </div>
        <div class="flex items-center gap-1 text-sm">
          <Icons icon="email" class="w-4 h-4 flex items-center" />
          <div class="">reachout@customercare.com</div>
        </div>
      </div>
    </div>

    <div
      class="fixed bottom-0 left-0 h-full w-full bg-white shadow-lg z-50 transition-transform duration-500 ease-in-out"
      :class="quickView ? 'translate-y-0' : 'translate-y-full'"
    >
      <div class="h-full overflow-auto" :style="{ height: '100dvh' }">
        <div
          class="flex items-center p-4 justify-end sticky top-0 bg-white z-50"
        >
          <div class="active:scale-95" @click="onOpenQuickView(false)">
            <Icons
              icon="close"
              alt="Close"
              title="Close"
              class="active:scale-95"
            />
          </div>
        </div>

        <div class="p-4 flex flex-col gap-2" v-if="image">
          <img :src="image" alt="" srcset="" class="rounded-xl w-full h-full" />

          <div class="text-xl">{{ title }}</div>
          <div class="flex items-center justify-between">
            <div class="">By: {{ published_by }}</div>

            <div class="flex items-center gap-1 text-sm">
              <Icons icon="clock" class="w-4 h-4 flex items-center" />
              <div class="">20 days</div>
            </div>
          </div>
        </div>

        <div class="px-4 pb-17 flex flex-col gap-4 text-justify">
          {{ content }}
        </div>

        <div class="p-4 bottom-0 left-0 fixed w-full bg-white shadow">
          Bottom
        </div>
      </div>
    </div>
  </main>
</template>
