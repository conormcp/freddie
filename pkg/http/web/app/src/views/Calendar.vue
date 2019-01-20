<template>
  <layout>
    <v-container>
      <v-layout>
        <v-flex xs12 md6 mx-2>
          <v-card light>
            <v-date-picker
              v-model="picker"
              full-width
              :events="dates"
            ></v-date-picker>
          </v-card>
        </v-flex>
        <v-flex xs12 md6 mx-2>
          <v-card light>
            <v-card-title>
              <h1>{{ picker }}</h1>
            </v-card-title>
            <v-divider />
            <v-card v-if="dates.includes(picker)">
              <v-card-title>
                <v-btn flat :href="currentMeeting.url" target="_blank" block>
                  <h2>{{ currentMeeting.title }}</h2>
                </v-btn>
              </v-card-title>
              <v-card-text>
                <h3>
                  {{ currentMeeting.speaker }} &mdash;
                  {{ currentMeeting.affil }}
                </h3>
                <h4>Date: {{ formatDate(currentMeeting.date) }}</h4>
                <h4>Time: {{ formatTime(currentMeeting.date) }}</h4>
              </v-card-text>
            </v-card>
            <v-card v-else>
              <v-card-title>
                <v-btn
                  flat
                  :href="currentMeeting.url"
                  target="_blank"
                  block
                  disabled
                >
                  <h2>No Meeting</h2>
                </v-btn>
              </v-card-title>
            </v-card>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </layout>
</template>

<script>
import { mapState, mapActions } from "vuex";
import formatDate from "@utils/format-date";
import isoTodayLocal from "@utils/iso-today-local";
import formatTime from "@utils/format-time";
import formatDateVuetify from "@utils/format-date-vuetify";
import Layout from "@layouts/main";

var isoDate = isoTodayLocal();

export default {
  metaInfo: {
    title: "Calendar"
  },
  components: {
    Layout
  },
  data() {
    return {
      picker: isoDate
    };
  },
  computed: {
    ...mapState([
      // map this.meetings to store.state.meetings
      "meetings"
    ]),
    dates: function() {
      return this.meetings.map(x => formatDateVuetify(x.date));
    },
    currentMeeting: function() {
      var idx = this.dates.indexOf(this.picker);
      if (idx != -1) {
        return this.meetings[idx];
      }
      return {};
    }
  },
  watch: {
    picker: function(val) {
      this.date = Date(val);
    }
  },
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
