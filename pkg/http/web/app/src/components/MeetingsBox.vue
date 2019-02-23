<template>
  <v-card>
    <v-card-title>
      <h1 class="display-2">Meetings</h1>
      <v-spacer />
      <new-meeting-dialog></new-meeting-dialog>
    </v-card-title>
    <v-card v-for="item in meetings" :key="item.index" class="my-2" light>
      <v-card-title>
        <h2>{{ item.title }}</h2>
      </v-card-title>
      <v-card-text>
        <h3>{{ item.speaker }} &mdash; {{ item.affil }}</h3>
        <h4>Date: {{ formatDate(item.date) }}</h4>
        <h4>Time: {{ formatTime(item.date) }}</h4>
        <br v-if="item.abstract" />
        <h4 v-if="item.abstract">Abstract: {{ item.abstract }}</h4>
      </v-card-text>
    </v-card>
  </v-card>
</template>

<script>
import { mapState, mapActions } from "vuex";
import formatDate from "@utils/format-date";
import formatTime from "@utils/format-time";
import NewMeetingDialog from "@components/NewMeetingDialog";

export default {
  components: {
    NewMeetingDialog
  },
  computed: mapState([
    // map this.papers to store.state.papers
    "meetings"
  ]),
  created: function() {
    this.fetchMeetings();
  },
  methods: {
    ...mapActions([
      "fetchMeetings" // map `this.fetchMeetings()` to `this.$store.dispatch('fetchMeetings')`
    ]),
    formatTime(d) {
      return formatTime(d);
    },
    formatDate(d) {
      return formatDate(d);
    }
  }
};
</script>
