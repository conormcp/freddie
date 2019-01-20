export default function formatDateVuetify(dateString) {
  // first convert to local time
  var d = new Date(dateString).toString();
  // then convert to ISO format and clip
  d = new Date(d).toISOString();
  return d.substr(0, 10);
}
