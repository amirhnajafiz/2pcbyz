<template>
  <div>
    <app-navigator />
    <router-view>
    </router-view>
  </div>
</template>

<script>
import AppNavigator from './components/AppNavigator.vue'
export default {
  components: { AppNavigator },

}
</script>

<style>

</style>