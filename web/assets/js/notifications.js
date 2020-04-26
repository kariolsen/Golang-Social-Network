function loadNotifications() {
    $.ajax({
        url: "/api/notifications",
        type: "GET",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) {
            $('#Slim').html("");
            $('#notifications_num').val("");
            var noti_messages = "";
            $.each(r.notifications, function(index, user) {
                noti_messages = noti_messages + '<a href="javascript:void(0);" class="dropdown-item notify-item"><div class="notify-icon bg-success"><i class="fa fa-comment"></i></div>'
                if (user.type == 0) {
                    noti_messages = noti_messages + '<p class="notify-details">' + user.user_name + ' mentioned you in post "' + user.title + '"<small class="text-muted">' + user.created_date + '</small></p></a>'
                } else if (user.type == 1) {
                    noti_messages = noti_messages + '<p class="notify-details">' + user.user_name + ' liked your post "' + user.title + '"<small class="text-muted">' + user.created_date + '</small></p></a>'
                } else if (user.type == 2) {
                    noti_messages = noti_messages + '<p class="notify-details">' + user.user_name + ' commented on your post "' + user.title + '"<small class="text-muted">' + user.created_date + '</small></p></a>'
                }
            });
            $('#Slim').html(
                $('#Slim').html() + noti_messages
            );
            $('#notifications_num').text(r.notifications.length);
        },
        error: function(e) {
            console.log(e);
        }
    });
}

function clearNotifications() {
    $.ajax({
        url: "/api/notifications/clear",
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) { loadNotifications(); },
        error: function(e) {
            console.log(e);
        }
    });
}