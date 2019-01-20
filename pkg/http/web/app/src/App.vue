<template>
  <v-app dark> <RouterView :key="$route.fullPath"></RouterView> </v-app>
</template>

<script>
export default {
  name: "App",
  metaInfo: {
    // if no subcomponents specify a metaInfo.title, this title will be used
    title: "Home",
    // all titles will be injected into this template
    titleTemplate: "%s | FREDDIE"
  },
  data() {
    return {
      //
    };
  }
};
</script>
