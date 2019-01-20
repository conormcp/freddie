<template>
  <div>
    <nav-bar />
    <v-content> <slot /> </v-content>
  </div>
</template>

<script>
import NavBar from "../components/NavBar";

export default {
  components: {
    NavBar
  }
};
</script>
