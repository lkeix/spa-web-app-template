<template>
  <div>
    <v-navigation-drawer
      app
      v-model="drawer"
      @transitionend="drawerFlip"
      clipped
    >
      <v-list flat>
        <v-list-item
          v-for="item in items"
          v-bind:key="item.i"
          @click="to(item.link)"
        >
          <v-list-item-content>
            <v-list-item-title>
              {{ item.title }}
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ item.subtitle }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  created() {
    this.items.push({
      title: "ログイン",
      subtitle: "ユーザ認証を行います",
      link: "/",
    });
  },
  data() {
    return {
      drawer: false,
      items: [],
    };
  },
  methods: {
    to(link) {
      this.$router.push({ path: link });
    },
    drawerFlip() {
      this.drawer = !this.drawer;
    },
  },
};
</script>