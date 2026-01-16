<script setup lang="ts">
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";

const router = useRouter();
const route = useRoute();

// Mock Data for Dashboard
const stats = ref([
  {
    title: "Total Courts",
    value: "12",
    trend: "+2 this month",
    icon: "M13.5 16.875h3.375m0 0h3.375m-3.375 0V13.5m0 3.375v3.375M6 10.5h2.25a2.25 2.25 0 002.25-2.25V6a2.25 2.25 0 00-2.25-2.25H6A2.25 2.25 0 003.75 6v2.25A2.25 2.25 0 006 10.5zm0 9.75h2.25A2.25 2.25 0 0010.5 18v-2.25a2.25 2.25 0 00-2.25-2.25H6a2.25 2.25 0 00-2.25 2.25V18A2.25 2.25 0 006 20.25zm9.75-9.75H18a2.25 2.25 0 002.25-2.25V6A2.25 2.25 0 0018 3.75h-2.25A2.25 2.25 0 0013.5 6v2.25a2.25 2.25 0 002.25 2.25z",
    color: "bg-blue-50 text-blue-600",
  },
  {
    title: "Active Bookings",
    value: "84",
    trend: "12 ending today",
    icon: "M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121 7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25 2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5",
    color: "bg-emerald-50 text-emerald-600",
  },
  {
    title: "Monthly Revenue",
    value: "$12.5k",
    trend: "+15% vs last month",
    icon: "M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z",
    color: "bg-amber-50 text-amber-600",
  },
]);

const recentActivity = ref([
  {
    action: "New Booking",
    details: "Court A - 2 hours",
    time: "10 mins ago",
    type: "booking",
  },
  {
    action: "Maintenance",
    details: "Court C - Resurfacing",
    time: "2 hours ago",
    type: "maintenance",
  },
  {
    action: "New Court",
    details: 'Added "Riverside Tennis"',
    time: "1 day ago",
    type: "add",
  },
  {
    action: "Payment",
    details: "Booking #8821 - $45.00",
    time: "1 day ago",
    type: "payment",
  },
]);

const quickActions = [
  {
    label: "Add New Court",
    icon: "M12 4.5v15m7.5-7.5h-15",
    route: "/court-wizard/add",
    color: "from-blue-500 to-indigo-600",
  },
  {
    label: "Manage Schedule",
    icon: "M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121 7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25 2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5",
    route: "#",
    color: "from-emerald-500 to-teal-600",
  },
  {
    label: "View Reports",
    icon: "M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z",
    route: "#",
    color: "from-purple-500 to-pink-600",
  },
];

const navigateTo = (route: string) => {
  if (route !== "#") {
    router.push(route);
  }
};
</script>

<template>
  <main>
    <component :is="route.meta.layout">
      <div class="min-h-screen pb-20">
        <!-- Header -->

        <main class="mx-auto p-6 space-y-8">
          <!-- Quick Stats -->
          <section class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div
              v-for="(stat, index) in stats"
              :key="index"
              class="bg-white p-6 rounded-2xl shadow-[0_2px_10px_-3px_rgba(6,81,237,0.1)] border border-gray-100 hover:shadow-lg transition-shadow duration-300"
            >
              <div class="flex items-start justify-between">
                <div>
                  <p class="text-sm font-medium text-gray-500 mb-1">
                    {{ stat.title }}
                  </p>
                  <h3 class="text-3xl font-bold text-gray-900">
                    {{ stat.value }}
                  </h3>
                </div>
                <div :class="`p-3 rounded-xl ${stat.color}`">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    class="w-6 h-6"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      :d="stat.icon"
                    />
                  </svg>
                </div>
              </div>
              <div class="mt-4 flex items-center text-sm">
                <span
                  class="text-emerald-600 font-medium bg-emerald-50 px-2 py-0.5 rounded-md"
                  v-if="stat.trend.includes('+')"
                  >{{ stat.trend }}</span
                >
                <span class="text-gray-500 ml-2" v-else>{{ stat.trend }}</span>
              </div>
            </div>
          </section>

          <!-- Quick Actions -->
          <section>
            <h2 class="text-lg font-semibold text-gray-800 mb-4">
              Quick Actions
            </h2>
            <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <button
                v-for="(action, index) in quickActions"
                :key="index"
                @click="navigateTo(action.route)"
                class="group relative overflow-hidden rounded-2xl p-6 text-left shadow-sm hover:shadow-md transition-all duration-300 bg-white border border-gray-100"
              >
                <div
                  class="absolute inset-0 bg-gradient-to-br opacity-0 group-hover:opacity-5 transition-opacity duration-300"
                  :class="action.color"
                ></div>
                <div class="flex items-center gap-4">
                  <div
                    class="p-3 rounded-xl bg-gradient-to-br text-white shadow-lg shadow-gray-200"
                    :class="action.color"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke-width="1.5"
                      stroke="currentColor"
                      class="w-6 h-6"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        :d="action.icon"
                      />
                    </svg>
                  </div>
                  <span
                    class="font-semibold text-gray-700 group-hover:text-gray-900 transition-colors"
                    >{{ action.label }}</span
                  >
                </div>
              </button>
            </div>
          </section>

          <!-- Recent Activity -->
          <section
            class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden"
          >
            <div
              class="p-6 border-b border-gray-100 flex items-center justify-between"
            >
              <h2 class="text-lg font-semibold text-gray-800">
                Recent Activity
              </h2>
              <button
                class="text-sm text-indigo-600 font-medium hover:text-indigo-700"
              >
                View All
              </button>
            </div>
            <div class="divide-y divide-gray-50">
              <div
                v-for="(item, index) in recentActivity"
                :key="index"
                class="p-4 hover:bg-gray-50 transition-colors flex items-center justify-between"
              >
                <div class="flex items-center gap-4">
                  <div
                    class="h-10 w-10 rounded-full flex items-center justify-center text-lg"
                    :class="{
                      'bg-blue-100 text-blue-600': item.type === 'booking',
                      'bg-amber-100 text-amber-600':
                        item.type === 'maintenance',
                      'bg-green-100 text-green-600': item.type === 'add',
                      'bg-purple-100 text-purple-600': item.type === 'payment',
                    }"
                  >
                    <!-- Simple type icons -->
                    <span v-if="item.type === 'booking'">📅</span>
                    <span v-else-if="item.type === 'maintenance'">🔧</span>
                    <span v-else-if="item.type === 'add'">✨</span>
                    <span v-else-if="item.type === 'payment'">💰</span>
                  </div>
                  <div>
                    <p class="text-sm font-medium text-gray-900">
                      {{ item.action }}
                    </p>
                    <p class="text-xs text-gray-500">{{ item.details }}</p>
                  </div>
                </div>
                <span class="text-xs text-gray-400 font-medium">{{
                  item.time
                }}</span>
              </div>
            </div>
          </section>
        </main>
      </div>
    </component>
  </main>
</template>
