<template>
  <v-container>
    <v-form ref="form">
      <v-text-field :counter="35" label="Name" required dark></v-text-field>

      <v-text-field label="Order" required dark></v-text-field>

      <v-row class="mb-2">
          <font-awesome-icon icon="camera" class="white--text mt-4 ml-2" style="font-size: 30px; margin-right: -20px;" />
        <v-file-input label="Image" dark show-size></v-file-input>
      </v-row>

      <v-btn color="success" class="mr-4"> Submit </v-btn>
    </v-form>
  </v-container>
</template>

<script>
export default {};
</script>

<style>
</style>