$(document).ready(function() {
    if ($('.flash-error').length > 0) {
        alert($('.flash-error').text());
    }
});

function confirmClearDatabase() {
    if (confirm("Are you sure you want to clean up the All database?")) {
        fetch('/clear_database', { method: 'POST' })
            .then(response => response.json())
            .then(data => {
                if (data.redirect) {
                    window.location.href = data.redirect;
                }
            });
    }
}

function DownloadMovies() {
    if (confirm("To download movies?")) {
        fetch('/download_films', { method: 'POST' })
            .then(response => response.json())
            .then(data => {
                if (data.redirect) {
                    window.location.href = data.redirect;
                }
            });
    }
}
