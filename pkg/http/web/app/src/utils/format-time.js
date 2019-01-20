export default function formatTime(dateString) {
  var d = new Date(dateString);
  return d.toLocaleTimeString();
}
