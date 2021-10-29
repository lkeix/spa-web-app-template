<template>
  <div>
    <v-form>
      <v-text-field label="ID" v-model="user.id" />
      <v-text-field label="Name" v-model="user.name" />
      <v-text-field label="Mail" v-model="user.mail" />
      <v-btn color="primary" @click="onSubmit">送信</v-btn>
    </v-form>
  </div>
</template>

<script>
import post from '@/lib/post.js'

export default {
  data: () => {
    return {
      user: {
        id: '',
        name: '',
        mail: '',
      }
    }
  },
  methods: {
    onSubmit() {
      let params = new URLSearchParams()
      params.append("userID", this.user.id)
      params.append("name", this.user.name)
      params.append("mail", this.user.mail)
      post("/api/insertUser", params)
    }
  }
}
</script>