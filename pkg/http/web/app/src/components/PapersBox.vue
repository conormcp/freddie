<template>
  <v-card>
    <v-card-title>
      <h1 class="display-2">Papers</h1>
      <v-spacer />
      <new-paper-dialog> </new-paper-dialog>
    </v-card-title>
    <v-card v-for="item in papers" :key="item.index" class="my-2" light>
      <v-card-title>
        <h2>{{ item.title }}</h2>
      </v-card-title>
      <v-divider />
      <v-card-text>
        <h3>
          Link:
          <v-btn flat :href="item.url" target="_blank" icon>
            <v-icon>link</v-icon>
          </v-btn>
        </h3>
        <h3>
          <span v-for="(a, i) in item.authors" :key="a.index">
            <span v-if="i < 3"> <span v-if="i">; </span>{{ a }}</span>
            <span v-else-if="i == 3">; et al.</span>
          </span>
        </h3>
        <h4>Added: {{ formatDate(item.added) }}</h4>
      </v-card-text>
    </v-card>
  </v-card>
</template>

<script>
import { mapState, mapActions } from "vuex";
import formatDate from "@utils/format-date";
import NewPaperDialog from "@components/NewPaperDialog";

export default {
  components: { NewPaperDialog },
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
    ]),
    formatDate(d) {
      return formatDate(d);
    }
  }
};
</script>
