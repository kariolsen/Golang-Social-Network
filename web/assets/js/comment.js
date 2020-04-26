function commentPost(button, post_id) {
    content = $(button).siblings(".post-comment-input").val();
    if (content != "") {
        var form = new FormData();
        form.append("content", content);

        $.ajax({
            url: "/api/comments/add/" + parseInt(post_id),
            type: "POST",
            timeout: 0,
            processData: false,
            mimeType: "multipart/form-data",
            contentType: false,
            dataType: 'json',
            data: form,
            success: function(r) {
                $(button).siblings(".post-comment-input").val("");
            },
            error: function(e) {
                notifyUser('Error!', 'Search Failed !', 'danger', 3000);
            }
        });
    }
}

function loadComments(comments, init, allow_comments) {
    var c_board = "";
    if (comments != 0 && comments != false) {
        if (init == true) {
            c_board = c_board + '<ul class="img-comment-list">'
        }
        $.each(comments, function(index, comment) {
            c_board = c_board + '<li><div class="comment-img"><a href="/view_users?name=' + comment.username + '"><img src="users/' + comment.user_id + '/profile/' + comment.avatar + '" class="img-responsive img-circle" alt="Image"/></a></div><div class="comment-text"><strong><a href="/view_users?name=' + comment.username + '">' + comment.username + '</a></strong><p class="">' + comment.content + '</p> <span class="date sub-text">' + comment.comment_date + '</span></div></li>'
        });
        if (init == true) {
            c_board = c_board + '</ul>'
        }
    } else {
        if (allow_comments == false) {
            c_board = c_board + '<ul class="img-comment-list"><li><p style="text-align: center;">This post does not allow comments</p></li></ul>'
        } else {
            c_board = c_board + '<ul class="img-comment-list"><li><p style="text-align: center;">This post has no comments</p></li></ul>'
        }
    }
    return c_board;
}

function reloadComments(submit_button, post_id, init, allow_comments) {
    var comments_content = "";
    comment_area = $(submit_button).parents(".modal-meta-bottom").siblings(".img-comment-list");
    comment_area.html("");
    $.ajax({
        url: "/api/comments/show/" + post_id.toString(),
        type: "GET",
        timeout: 0,
        dataType: 'json',
        success: function(r) {
            comments_content = loadComments(r.comments, init, allow_comments);
            comment_area.html(
                comment_area.html() + comments_content
            )
        },
        error: function(e) { console.log(e); }
    });
}

function publishComments(submit_button, post_id, init, allow_comments) {
    if (allow_comments == true) {
        commentPost(submit_button, post_id);
        setTimeout(function() { reloadComments(submit_button, post_id, init, allow_comments); }, 500);
    }
}