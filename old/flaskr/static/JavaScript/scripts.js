$(document).ready(function () {
  if ($(".flash-error").length > 0) {
    alert($(".flash-error").text());
  }
});

function confirmClearDatabase() {
  if (confirm("Are you sure you want to clean up the All database?")) {
    fetch("/clear_database", { method: "POST" })
      .then((response) => response.json())
      .then((data) => {
        if (data.redirect) {
          window.location.href = data.redirect;
        }
      });
  }
}

function DownloadMovies() {
  var number = prompt(
    "enter the quantity for EVERY movie to be added to the warehouse:"
  );
  if (number !== null) {
    number = parseInt(number);
    if (number >= 1 && number <= 1000) {
      // send a number
      fetch("/download_films", {
        method: "POST",
        body: number.toString(),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.redirect) {
            window.location.href = data.redirect;
          }
        });
    } else {
      alert("Enter a number between 1 and 1000");
    }
  }
}
