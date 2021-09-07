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

      <v-btn color="success" v-on:click="submitForm()" class="mr-4">
        Submit
      </v-btn>
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
      let data = new FormData();
      data.append("name", this.name);
      data.append("order", this.order);
      data.append("image", this.image);
      axios
        .post("http://localhost:8000/api/add/category", data)
        .then((response) => {
          if (response.status >= 200 && response.status < 300) {
            this.$router.push("categories");
            alert("Category added successfully");
          } else {
            alert("Something went wrong.");
          }
        }).catch(error => {
            alert("Something went wrong with your request. Code " + error.response.status);
        });
      this.$refs.form.reset();
    },
  },
};
</script>

<style>
</style>