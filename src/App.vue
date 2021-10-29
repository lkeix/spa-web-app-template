<template>
  <v-app>
    <app-header />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import AppHeader from "@/components/Header.vue";
export default {
  components: {
    "app-header": AppHeader,
  },
};
</script>
