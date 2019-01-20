export default function isoTodayLocal() {
  var date = new Date(); // Or the date you'd like converted.
  var isoDate = new Date(date.getTime() - date.getTimezoneOffset() * 60000)
    .toISOString()
    .substr(0, 10);
  return isoDate;
}
