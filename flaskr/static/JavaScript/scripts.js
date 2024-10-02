$(document).ready(function() {
    if ($('.flash-error').length > 0) {
        alert($('.flash-error').text());
    }
});

function confirmClearDatabase() {
    if (confirm("Are you sure you want to clean up the All database?")) {
        fetch('/clear_database', { method: 'POST' });
    }
}
