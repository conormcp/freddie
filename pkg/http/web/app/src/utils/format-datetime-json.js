function pad(n) {
  return n < 10 ? "0" + n : n;
}

function timezoneOffset() {
  var offset = new Date().getTimezoneOffset();
  var sign;
  if (offset === 0) {
    return "Z";
  }
  sign = offset > 0 ? "-" : "+";
  offset = Math.abs(offset);
  return sign + pad(Math.floor(offset / 60)) + ":" + pad(offset % 60);
}

export default function formatDatetimeJSON(date, time) {
  return `${date}T${time}:00${timezoneOffset()}`;
}
