<template>
  <v-dialog v-model="dialog" max-width="600px">
    <v-btn slot="activator" icon color="primary"> <v-icon>add</v-icon> </v-btn>
    <v-card light>
      <v-card-title> <span class="headline">Add Meeting</span> </v-card-title>
      <v-card-text>
        <v-container grid-list-md>
          <v-layout wrap>
            <v-flex xs12>
              <v-date-picker
                v-model="datePicker"
                landscape
                reactive
              ></v-date-picker>
            </v-flex>
            <v-flex xs12>
              <v-text-field
                label="Speaker"
                v-model="meeting.speaker"
                required
              ></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-text-field
                label="Title"
                v-model="meeting.title"
              ></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-text-field
                label="Afilliation"
                required
                v-model="meeting.affil"
              ></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-expansion-panel popout>
                <v-expansion-panel-content>
                  <div slot="header">Time</div>
                  <v-time-picker
                    v-model="timePicker"
                    landscape
                    reactive
                  ></v-time-picker>
                </v-expansion-panel-content>
              </v-expansion-panel>
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
import formateDatetimeJSON from "@utils/format-datetime-json";
import { mapActions } from "vuex";

export default {
  data() {
    return {
      dialog: false,
      datePicker: new Date().toISOString().substr(0, 10),
      timePicker: "15:30",
      meeting: {
        date: null,
        speaker: null,
        affil: null,
        title: null
      }
    };
  },
  methods: {
    reset() {
      const keys = Object.keys(this.meeting);
      for (const key of keys) {
        this.meeting[key] = null;
      }
    },
    submit() {
      this.formateDate();
      console.log(this.meeting);
      axios
        .post("/api/meetings/", this.meeting)
        .then(res => {
          console.log(res);
          this.reset();
          this.fetchMeetings();
        })
        .catch(err => {
          console.log(err);
        })
        .finally(() => {
          this.dialog = false;
        });
    },
    formateDate() {
      this.meeting.date = formateDatetimeJSON(this.datePicker, this.timePicker);
    },
    ...mapActions([
      "fetchMeetings" // map `this.fetchMeetings()` to `this.$store.dispatch('fetchMeetings')`
    ])
  }
};
</script>
