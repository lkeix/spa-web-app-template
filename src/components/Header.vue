<template>
  <div>
    <app-navigation ref="navigation" />
    <v-app-bar dark color="indigo darken-4" app clipped-left>
      <v-app-bar-nav-icon @click="flipNavigation"></v-app-bar-nav-icon>
      <v-app-bar-title>
        <div>{{title}}</div>
      </v-app-bar-title>
    </v-app-bar>
  </div>
</template>

<script>
import Navigation from '@/components/Navigation.vue'
export default {
  components: {
    "app-navigation": Navigation
  },
  data() {
    return {
      title: "text p2p",
    }
  },
  methods: {
    flipNavigation() {
      this.$refs.navigation.drawerFlip()
    }
  }
}
</script>
