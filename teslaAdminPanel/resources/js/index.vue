<template>
  <v-app style="background-color:#000000; height:100vh;">
    <app-navigator />
    <router-view>
    </router-view>
  </v-app>
</template>

<script>
import AppNavigator from './components/AppNavigator.vue'
export default {
  components: { AppNavigator },

}
</script>

<style>

</style>