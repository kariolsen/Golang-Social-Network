function edit_profile_data(new_name, new_email, new_quote) {
    var form = new FormData();
    form.append("username", new_name);
    form.append("quote", new_quote);
    form.append("email", new_email);

    $.ajax({
        url: "/api/profile",
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        data: form,
        success: function(r) { window.location.href = "/view_profile"; },
        error: function(e) {
            $('#edit_profile_err_message').html($('#edit_profile_err_message').html() + 'Failed to update profile </br>');
            setTimeout(function() {
                $('#edit_profile_err_message').html("");
            }, 2000);
        }
    });
}

function edit_profile_data_and_avatar(new_name, new_email, new_quote, new_avatar) {
    var form = new FormData();
    form.append("avatar", new_avatar);
    $.ajax({
        url: "/api/profile_avatar",
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        data: form,
        success: function(r) { edit_profile_data(new_name, new_email, new_quote); },
        error: function(e) {
            $('#edit_profile_err_message').html($('#edit_profile_err_message').html() + 'You already have an avatar with the same name, please rename it </br>');
            setTimeout(function() {
                $('#edit_profile_err_message').html("");
            }, 2000);
        }
    });
}