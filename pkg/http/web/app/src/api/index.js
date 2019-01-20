import axios from "axios";

function getPapers() {
  return axios.get("/api/papers/");
}

function getMeetings() {
  return axios.get("/api/meetings/");
}

export default {
  getPapers,
  getMeetings
};
