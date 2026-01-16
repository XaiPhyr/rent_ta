<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useRouter, useRoute } from "vue-router";
// @ts-ignore - types might not be installed yet
import L from "leaflet";
import "leaflet/dist/leaflet.css";

const router = useRouter();
const route = useRoute();

// Form State
const form = ref({
  name: "",
  location: null as { lat: number; lng: number } | null,
  sports: [] as string[],
  isIndoor: false,
  surface: "",
  hasLighting: false,
  capacity: null as number | null,
});

// Options
const availableSports = [
  "Basketball",
  "Tennis",
  "Volleyball",
  "Badminton",
  "Soccer",
  "Pickleball",
];
const surfaceOptions = ["Wood", "Cement", "Turf", "Clay", "Acrylic"];

// Map Ref
const mapContainer = ref<HTMLElement | null>(null);
const map = ref<any>(null);
const marker = ref<any>(null);

// Initialize Map
onMounted(() => {
  if (mapContainer.value) {
    // Default to a central location (e.g., city center) - User should ideally update this from user location
    // Using a generic start point for now: Manila coordinates example or just 0,0
    const startLat = 14.5995;
    const startLng = 120.9842;

    map.value = L.map(mapContainer.value).setView([startLat, startLng], 13);

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(map.value);

    // Click handler to set location
    map.value.on("click", (e: any) => {
      const { lat, lng } = e.latlng;
      form.value.location = { lat, lng };

      if (marker.value) {
        marker.value.setLatLng([lat, lng]);
      } else {
        marker.value = L.marker([lat, lng]).addTo(map.value);
      }
    });
  }
});

// Form Submission
const submitForm = () => {
  // Basic validation
  if (!form.value.name || !form.value.location) {
    alert("Please fill in Name and pick a Location on the map.");
    return;
  }
  console.log("Submitting Court Data:", form.value);
  alert("Court Added! (Check console for data)");
  // router.push('/dashboard'); // distinct success action
};

const goBack = () => {
  router.back();
};

const toggleSport = (sport: string) => {
  if (form.value.sports.includes(sport)) {
    form.value.sports = form.value.sports.filter((s) => s !== sport);
  } else {
    form.value.sports.push(sport);
  }
};
</script>

<template>
  <main>
    <component :is="route.meta.layout">
      <div class="min-h-screen pb-20">
        <!-- Header -->
        <div
          class="bg-white sticky top-0 z-20 shadow-sm px-4 py-3 flex items-center gap-3"
        >
          <!-- <button @click="goBack" class="p-2 rounded-full hover:bg-gray-100">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6 text-gray-700"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18"
              />
            </svg>
          </button> -->
          <h1
            class="text-xl font-bold text-gray-900 bg-clip-text text-transparent bg-gradient-to-r from-blue-600 to-indigo-600"
          >
            Add New Court
          </h1>
        </div>

        <div class="p-4 space-y-6">
          <!-- General Info -->
          <section
            class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100"
          >
            <h2
              class="text-lg font-semibold text-gray-800 mb-4 flex items-center gap-2"
            >
              <span class="bg-blue-100 text-blue-600 p-1.5 rounded-lg">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  class="w-5 h-5"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25"
                  />
                </svg>
              </span>
              Basic Information
            </h2>

            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1"
                  >Court Name</label
                >
                <input
                  v-model="form.name"
                  type="text"
                  placeholder="e.g. Downtown Hoops Center"
                  class="w-full px-4 py-2.5 rounded-xl border border-gray-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-100 outline-none transition-all placeholder:text-gray-400 bg-gray-50 focus:bg-white"
                />
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Capacity</label
                  >
                  <input
                    v-model.number="form.capacity"
                    type="number"
                    placeholder="Max players"
                    class="w-full px-4 py-2.5 rounded-xl border border-gray-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-100 outline-none transition-all placeholder:text-gray-400 bg-gray-50 focus:bg-white"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Surface</label
                  >
                  <select
                    v-model="form.surface"
                    class="w-full px-4 py-2.5 rounded-xl border border-gray-200 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-100 outline-none transition-all bg-gray-50 focus:bg-white"
                  >
                    <option value="" disabled>Select Surface</option>
                    <option v-for="s in surfaceOptions" :key="s" :value="s">
                      {{ s }}
                    </option>
                  </select>
                </div>
              </div>
            </div>
          </section>

          <!-- Location Map -->
          <section
            class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 overflow-hidden"
          >
            <h2
              class="text-lg font-semibold text-gray-800 mb-4 flex items-center gap-2"
            >
              <span class="bg-emerald-100 text-emerald-600 p-1.5 rounded-lg">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  class="w-5 h-5"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z"
                  />
                </svg>
              </span>
              Location
            </h2>
            <p class="text-sm text-gray-500 mb-3">
              Tap on the map to pin the exact court location.
            </p>

            <div
              class="h-64 sm:h-80 w-full rounded-xl overflow-hidden border border-gray-200 bg-gray-100 relative z-0"
            >
              <div ref="mapContainer" class="h-full w-full"></div>
              <!-- Map Overlay Hint if no loc selected -->
              <div
                v-if="!form.location"
                class="absolute bottom-4 left-1/2 -translate-x-1/2 bg-white/90 backdrop-blur-sm px-4 py-2 rounded-full shadow-lg text-xs font-medium text-gray-600 z-[1000] pointer-events-none"
              >
                Tap to set location
              </div>
            </div>
            <div
              v-if="form.location"
              class="mt-3 text-sm text-emerald-600 flex items-center gap-1 font-medium bg-emerald-50 w-fit px-3 py-1.5 rounded-lg"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                class="w-4 h-4"
              >
                <path
                  fill-rule="evenodd"
                  d="M9.69 18.933l.003.001C9.89 19.02 10 19 10 19s.11.02.308-.066l.002-.001.006-.003.018-.008a5.741 5.741 0 00.281-.14c.186-.096.446-.24.757-.433.62-.384 1.445-.966 2.274-1.765C15.302 14.988 17 12.493 17 9A7 7 0 103 9c0 3.492 1.698 5.988 3.355 7.625a19.015 19.015 0 002.274 1.765c.311.193.571.337.757.433.092.047.185.093.281.14l.018.008.006.003.002.001zM10 12a3 3 0 100-6 3 3 0 000 6z"
                  clip-rule="evenodd"
                />
              </svg>
              Pinned: {{ form.location.lat.toFixed(5) }},
              {{ form.location.lng.toFixed(5) }}
            </div>
          </section>

          <!-- Details -->
          <section
            class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100"
          >
            <h2
              class="text-lg font-semibold text-gray-800 mb-4 flex items-center gap-2"
            >
              <span class="bg-purple-100 text-purple-600 p-1.5 rounded-lg">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke-width="1.5"
                  stroke="currentColor"
                  class="w-5 h-5"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M9.53 16.122a3 3 0 00-5.78 1.128 2.25 2.25 0 01-2.4 2.245 4.5 4.5 0 008.4-2.245c0-.399-.078-.78-.22-1.128zm0 0a15.998 15.998 0 003.388-1.62m-5.04 2.825a15.971 15.971 0 01-2.45-6.522A5.996 5.996 0 011 .5 8.97 8.97 0 0113.824 10c0 1.282-.176 2.505-.513 3.655M1.502 14.5c1.063 1.063 2.54 1.655 4.128 1.655 1.588 0 3.065-.592 4.128-1.655"
                  />
                </svg>
              </span>
              Features & Sports
            </h2>

            <div class="space-y-6">
              <!-- Supported Sports -->
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Supported Sports</label
                >
                <div class="flex flex-wrap gap-2">
                  <button
                    v-for="sport in availableSports"
                    :key="sport"
                    @click="toggleSport(sport)"
                    class="px-4 py-2 rounded-full text-sm font-medium transition-all border"
                    :class="
                      form.sports.includes(sport)
                        ? 'bg-indigo-600 text-white border-indigo-600 shadow-md shadow-indigo-200'
                        : 'bg-white text-gray-600 border-gray-200 hover:border-gray-300 hover:bg-gray-50'
                    "
                  >
                    {{ sport }}
                  </button>
                </div>
              </div>

              <!-- Toggles -->
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                <!-- Indoor/Outdoor -->
                <div
                  class="flex items-center justify-between p-3 rounded-xl border border-gray-200 bg-gray-50"
                >
                  <div>
                    <span class="block text-sm font-medium text-gray-800"
                      >Environment</span
                    >
                    <span class="text-xs text-gray-500">{{
                      form.isIndoor ? "Indoor Court" : "Outdoor Court"
                    }}</span>
                  </div>
                  <button
                    @click="form.isIndoor = !form.isIndoor"
                    class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                    :class="form.isIndoor ? 'bg-indigo-600' : 'bg-gray-200'"
                  >
                    <span
                      class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      :class="form.isIndoor ? 'translate-x-6' : 'translate-x-1'"
                    />
                  </button>
                </div>

                <!-- Lighting -->
                <div
                  class="flex items-center justify-between p-3 rounded-xl border border-gray-200 bg-gray-50"
                >
                  <div>
                    <span class="block text-sm font-medium text-gray-800"
                      >Lighting</span
                    >
                    <span class="text-xs text-gray-500">{{
                      form.hasLighting ? "Available" : "Not Available"
                    }}</span>
                  </div>
                  <button
                    @click="form.hasLighting = !form.hasLighting"
                    class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                    :class="form.hasLighting ? 'bg-indigo-600' : 'bg-gray-200'"
                  >
                    <span
                      class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                      :class="
                        form.hasLighting ? 'translate-x-6' : 'translate-x-1'
                      "
                    />
                  </button>
                </div>
              </div>
            </div>
          </section>

          <!-- Submit Action -->
          <div class="pt-4">
            <button
              @click="submitForm"
              class="w-full bg-gradient-to-r from-indigo-600 to-blue-600 text-white font-bold text-lg py-4 rounded-2xl shadow-lg shadow-indigo-200 hover:shadow-xl hover:shadow-indigo-300 transform hover:-translate-y-0.5 active:translate-y-0 transition-all"
            >
              Save Court Details
            </button>
          </div>
        </div>
      </div>
    </component>
  </main>
</template>

<style scoped>
/* Leaflet map fix for z-index issues if any */
:deep(.leaflet-pane) {
  z-index: 10;
}
:deep(.leaflet-bottom.leaflet-right) {
  z-index: 20;
}
</style>
