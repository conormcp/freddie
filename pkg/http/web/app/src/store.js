import Vue from "vue";
import Vuex from "vuex";
import api from "./api";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    papers: [],
    meetings: []
  },
  mutations: {
    SET_PAPERS(state, p) {
      state.papers = p;
    },
    SET_MEETINGS(state, m) {
      state.meetings = m;
    }
  },
  actions: {
    fetchPapers({ commit }) {
      // get list of papers from api
      api
        .getPapers()
        .then(resp => {
          // update store with new data
          commit("SET_PAPERS", resp.data);
        })
        .catch(err => {
          console.log(err);
        });
    },
    fetchMeetings({ commit }) {
      // get list of meetings from api
      api.getMeetings().then(resp => {
        // update store with new data
        commit("SET_MEETINGS", resp.data);
      });
    }
  }
});
