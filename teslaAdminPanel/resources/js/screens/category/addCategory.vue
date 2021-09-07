<template>
  <v-container>
    <v-form ref="form">
      <v-text-field
        v-model="name"
        :counter="35"
        label="Name"
        required
        dark
      ></v-text-field>

      <v-text-field v-model="order" label="Order" required dark></v-text-field>

      <v-row class="mb-2">
        <font-awesome-icon
          icon="camera"
          class="white--text mt-4 ml-2"
          style="font-size: 30px; margin-right: -20px"
        />
        <v-file-input
          v-model="image"
          label="Image"
          dark
          show-size
        ></v-file-input>
      </v-row>

      <v-btn color="success" v-on:click="submitForm()" class="mr-4"> Submit </v-btn>
    </v-form>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      name: "",
      order: null,
      image: null,
    };
  },
  methods: {
      submitForm() {
          axios.post('http://localhost:8000/api/add/category', {
              name: this.name,
              order: this.order,
              image: this.image
          }).then(response => {
              console.log(response.data.name);
          });
          this.$refs.form.reset();
      }
  }
};
</script>

<style>
</style>