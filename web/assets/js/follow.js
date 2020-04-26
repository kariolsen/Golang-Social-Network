function followUser(follow_button, user_id) {
    follow_button.classList.toggle("not-following");
    $(follow_button).html("");

    $.ajax({
        url: "/api/user/followPressed/" + user_id.toString(),
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) {},
        error: function(e) {
            notifyUser('Error!', 'Search Failed !', 'danger', 3000);
        }
    });

    if ($(follow_button).hasClass("not-following")) {
        $(follow_button).html(
            $(follow_button).html() + '<i class="fa"></i>  Follow'
        );
    } else {
        $(follow_button).html(
            $(follow_button).html() + '<i class="fa fa-check"></i> Following'
        );
    }
}

function followHashtag(follow_button, hashtag_id) {
    follow_button.classList.toggle("not-following");
    $(follow_button).html("");

    $.ajax({
        url: "/api/hashtag/followPressed/" + hashtag_id.toString(),
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) {},
        error: function(e) {
            notifyUser('Error!', 'Search Failed !', 'danger', 3000);
        }
    });

    if ($(follow_button).hasClass("not-following")) {
        $(follow_button).html(
            $(follow_button).html() + '<i class="fa"></i>  Follow'
        );
    } else {
        $(follow_button).html(
            $(follow_button).html() + '<i class="fa fa-check"></i> Following'
        );
    }
}

function onlyFollowUser(user_id) {
    $.ajax({
        url: "/api/user/followPressed/" + user_id.toString(),
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) {},
        error: function(e) {
            notifyUser('Error!', 'Search Failed !', 'danger', 3000);
        }
    });
}