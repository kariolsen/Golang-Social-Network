function notifyUser(title, message, type, length) {
    var notify = $.notify({
        title: '<strong>' + title + '</strong>',
        message: message
    }, {
        type: type
    });
    setTimeout(function() {
        notify.close();
    }, length);
}