<template>
  <layout>
    <v-container> <papers-box /> </v-container>
  </layout>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Layout from "@layouts/main";
import PapersBox from "@components/PapersBox";

export default {
  metaInfo: {
    title: "Papers"
  },
  components: {
    Layout,
    PapersBox
  },
  computed: mapState([
    // map this.papers to store.state.papers
    "papers"
  ]),
  created: function() {
    this.fetchPapers();
  },
  methods: {
    ...mapActions([
      "fetchPapers" // map `this.fetch()` to `this.$store.dispatch('fetchPapers')`
    ])
  }
};
</script>
