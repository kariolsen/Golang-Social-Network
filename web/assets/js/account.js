function logout() {
    $.ajax({
        url: "/user/logout",
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) { window.location.href = "/login"; },
        error: function(e) { console.log(e); }
    });
}