<script setup lang="ts">
  import Card from '@/components/Card.vue';
  import Icons from '@/components/Icons.vue';
  import { useRoute } from 'vue-router';
  import { news, categories } from '@/sample/data';
  import { computed, ref } from 'vue';

  const route = useRoute();

  const total = (Math.floor(Math.random() * 99999) / 100).toFixed(2);

  const selectedCategories = ref<string[]>(['all']);

  const setSelectedCategories = (payload: string) => {
    if (payload === 'all') {
      selectedCategories.value = ['all'];
      return;
    }

    if (selectedCategories.value.includes('all')) {
      selectedCategories.value = [];
    }

    const categoryIndex = selectedCategories.value.indexOf(payload);
    if (categoryIndex > -1) {
      selectedCategories.value.splice(categoryIndex, 1);
    } else {
      selectedCategories.value.push(payload);
    }

    const selectedCategoriesWithoutAll = selectedCategories.value.filter(
      (category) => category !== 'all'
    );

    if (selectedCategoriesWithoutAll.length === categories.length - 1) {
      setTimeout(() => {
        selectedCategories.value = ['all'];
      }, 100);
    }

    if (selectedCategories.value.length === 0) {
      selectedCategories.value = ['all'];
    }
  };

  const filteredItems = computed(() => {
    const data = news.filter((x) =>
      x.tags.some((y) => selectedCategories.value.includes(y))
    );

    return data.length > 0 ? data : news;
  });
</script>

<template>
  <main>
    <component :is="route.meta.layout">
      <div class="flex flex-col gap-4">
        <div class="relative w-full">
          <Icons
            icon="magnifying-glass"
            alt="Search"
            title="Search"
            class="absolute left-3 top-1/2 transform -translate-y-1/2"
          />
          <input
            class="w-full pl-10 pr-4 py-3 text-sm text-center shadow border border-indigo-500 rounded-xl focus:outline-none focus:ring focus:ring-indigo-500"
            type="text"
            placeholder="Search"
          />
        </div>

        <div
          class="w-full h-auto shadow bg-gray-50 rounded-xl p-4 bg-indigo-400"
        >
          <div class="flex items-center justify-between">
            <div class="text-white">
              <div class="font-bold">Start your</div>
              <div class="font-bold text-2xl">Adventure here</div>
            </div>

            <div
              class="cursor-pointer bg-white rounded-xl px-4 py-2 font-bold active:scale-95"
            >
              Start here <span class="italic">!</span>
            </div>
          </div>
        </div>

        <div
          class="w-full h-auto shadow bg-gray-50 rounded-xl p-4 bg-green-900"
        >
          <div class="flex items-center">
            <div class="font-bold text-white text-xl">
              Total Donations: ${{ total }}
            </div>
          </div>
        </div>

        <div class="flex items-center font-bold text-xl">Causes</div>

        <div class="w-full overflow-auto">
          <div class="flex items-center gap-2">
            <div
              v-for="({ text, value }, index) in categories"
              :key="index"
              class="w-auto h-auto cursor-pointer px-8 py-2 rounded-xl active:scale-95"
              :class="
                selectedCategories.includes(value)
                  ? 'bg-indigo-900 text-white'
                  : 'bg-white'
              "
              @click="setSelectedCategories(value)"
            >
              {{ text }}
            </div>
          </div>
        </div>

        <Card
          v-for="(
            { title, content, image, published_by }, index
          ) in filteredItems"
          :key="index"
          :title="title"
          :content="content"
          :published_by="published_by"
          :image="image"
        />
      </div>
    </component>
  </main>
</template>
