<template>
  <div class="home">
    <v-container fluid>
      <v-card>
        <v-container fluid>
          <sample-form />
        </v-container>
      </v-card>
    </v-container>
  </div>
</template>

<script>
import form from '@/components/Form.vue'

export default {
  components: {
    'sample-form': form
  }
}
</script>
