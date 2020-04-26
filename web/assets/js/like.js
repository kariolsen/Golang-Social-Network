function likeHomePost(heart, post_id) {
    var likes = parseInt($(heart).parent().siblings(".likes_num").text());
    heart.classList.toggle("liked_post");
    $(heart).parent().siblings(".likes_num").toggleClass("liked_post");

    $.ajax({
        url: "/api/post/likePressed/" + post_id.toString(),
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) {
            console.log(r.likes_num);
        },
        error: function(e) {
            notifyUser('Error!', 'Search Failed !', 'danger', 3000);
        }
    });

    if ($(heart).hasClass("liked_post")) {
        likes = likes + 1;
        $(heart).parent().siblings(".likes_num").text(likes);
    } else {
        likes = likes - 1;
        $(heart).parent().siblings(".likes_num").text(likes);
    }
}


function likePost(heart, post_id) {
    var likes = parseInt($(heart).siblings(".likes_num").text());
    heart.classList.toggle("liked_post");
    $(heart).siblings(".likes_num").toggleClass("liked_post");

    $.ajax({
        url: "/api/post/likePressed/" + post_id.toString(),
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        success: function(r) {
            console.log(r.likes_num);
        },
        error: function(e) {
            notifyUser('Error!', 'Search Failed !', 'danger', 3000);
        }
    });

    if ($(heart).hasClass("liked_post")) {
        likes = likes + 1;
        $(heart).siblings(".likes_num").text(likes);
    } else {
        likes = likes - 1;
        $(heart).siblings(".likes_num").text(likes);
    }
}