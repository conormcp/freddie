<template>
  <v-dialog v-model="dialog" max-width="600px">
    <v-btn slot="activator" icon color="primary"> <v-icon>add</v-icon> </v-btn>
    <v-card light>
      <v-card-title> <span class="headline">Add Paper</span> </v-card-title>
      <v-card-text>
        <v-container grid-list-md>
          <v-layout wrap>
            <v-flex xs12>
              <v-text-field
                label="URL"
                v-model="url"
                required
                placeholder="http://ui.adsabs.harvard.edu/..."
                :rules="[rules.ads]"
              ></v-text-field>
            </v-flex>
          </v-layout>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" flat @click="dialog = false">Close</v-btn>
        <v-btn color="primary" flat @click="submit()">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import axios from "axios";
import { mapActions } from "vuex";

export default {
  data: () => ({
    dialog: false,
    url: "",
    rules: {
      ads: value =>
        value.includes("adsabs.harvard.edu") || "Must be an ADS URL"
    }
  }),
  computed: {
    error: function() {
      return !this.url.includes("adsabs.harvard.edu");
    }
  },
  methods: {
    submit() {
      console.log(this.url);
      axios
        .post("/api/papers/", { url: this.url })
        .then(res => {
          this.url = "";
          console.log(res);
          this.dialog = false;
          this.fetchPapers();
        })
        .catch(err => {
          console.log(err);
        });
    },
    ...mapActions([
      "fetchPapers" // map `this.fetchPapers()` to `this.$store.dispatch('fetchPapers')`
    ])
  }
};
</script>
