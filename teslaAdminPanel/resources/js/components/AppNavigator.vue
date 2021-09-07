<template>
  <div>
    <v-toolbar
      dark
      src="https://cdn.vuetifyjs.com/images/backgrounds/vbanner.jpg"
    >
      <v-toolbar-title> Tesla Admin Panel </v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn class="mr-2" @click="() => this.$router.push('addCategory')">
        Add Category
      </v-btn>
      <v-btn @click="() => this.$router.push('categories')">
        Categories
      </v-btn>
      <v-btn class="ml-2"> Products </v-btn>
    </v-toolbar>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>