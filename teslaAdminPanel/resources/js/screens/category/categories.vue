<template>
  <v-simple-table dark>
    <template v-slot:default>
      <thead>
        <tr>
          <th class="text-center">Image</th>
          <th class="text-center">Name</th>
          <th class="text-center">Ordering</th>
          <th class="text-center">Edit</th>
          <th class="text-center">Delete</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="category in categories" :key="category.id">
          <td class="text-center">
            <img
              src="https://bordalo.observador.pt/v2/rs:fill:900/c:433:433:nowe:0:0/q:86/plain/https://s3.observador.pt/wp-content/uploads/2017/01/26172929/tesla-car-logo_433x433_acf_cropped.jpg"
              style="height: 100px; width: 100px"
            />
          </td>
          <td class="text-center">
            {{ category.name }}
          </td>
          <td class="text-center">
            {{ category.order }}
          </td>
          <td class="text-center">
            <v-btn class="mx-2" fab dark large color="cyan">
              <font-awesome-icon icon="pen" />
            </v-btn>
          </td>
          <td class="text-center">
            <v-btn class="mx-2" fab dark large color="red">
              <font-awesome-icon icon="trash" />
            </v-btn>
          </td>
        </tr>
      </tbody>
    </template>
  </v-simple-table>
</template>

<script>
export default {
  data() {
    return {
      categories: [],
    };
  },
  methods: {
    getCategories() {
      axios.get("http://localhost:8000/api/categories").then((response) => {
        if (response.status >= 200 && response.status < 300) {
          this.categories = response.data.categories;
        }
      });
    },
  },
  mounted() {
    this.getCategories();
  },
};
</script>

<style>
</style>