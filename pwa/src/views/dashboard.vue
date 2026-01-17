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
