export default function formatDateVuetify(dateString) {
  // first convert to local time
  var date = new Date(dateString);
  // then convert to ISO format and clip

  var isoDate = new Date(date.getTime() - date.getTimezoneOffset() * 60000)
    .toISOString()
    .substr(0, 10);
  return isoDate;
}
