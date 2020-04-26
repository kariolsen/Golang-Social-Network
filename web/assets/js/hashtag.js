function showHashtagPosts(hashtag_name, hashtag_id) {
    var form = new FormData();
    form.append("hashtag_name", hashtag_name);
    form.append("hashtag_id", hashtag_id);
    $.ajax({
        url: "/api/hashtag/posts",
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        data: form,
        success: function(r) {
            var posts_num = r.posts.length;
            var hashtag_post_board = "";
            $('#hashtag_post_board').html("")
            if (posts_num > 0) {
                $.each(r.posts, function(index, post) {
                    if ((index + 1) % 3 == 1) {
                        hashtag_post_board = hashtag_post_board + '<div class="row"><div class="col-lg-4"><a href="#"><div class="storybox" style="background: linear-gradient( rgba(34,34,34,0.78), rgba(34,34,34,0.78)), url(users/' + post.user_id + '/posts/' + post.post_id + '/' + post.images[0].image_name + ') no-repeat;background-size: cover;background-position: center center;-webkit-background-size: cover;-moz-background-size: cover;-o-background-size: cover;"><div class="story-body text-center"><div class=""><a href="/view_users?name=' + post.user_name + '"><img class="img-circle" src="users/' + post.user_id + '/profile/' + post.avatar + '" alt="user"></div></a><h4>' + post.user_name + '</h4><a href="/view_hashtag?name=' + hashtag_name + '"><h4>#' + hashtag_name + '#</h4></a><p>' + post.created_date + '</p></div></div></a></div>'
                    } else if ((index + 1) % 3 == 0 || index + 1 == posts_num) {
                        hashtag_post_board = hashtag_post_board + '<div class="col-lg-4"><a href="#"><div class="storybox" style="background: linear-gradient( rgba(34,34,34,0.78), rgba(34,34,34,0.78)), url(users/' + post.user_id + '/posts/' + post.post_id + '/' + post.images[0].image_name + ') no-repeat;background-size: cover;background-position: center center;-webkit-background-size: cover;-moz-background-size: cover;-o-background-size: cover;"><div class="story-body text-center"><div class=""><a href="/view_users?name=' + post.user_name + '"><img class="img-circle" src="users/' + post.user_id + '/profile/' + post.avatar + '" alt="user"></div></a><h4>' + post.user_name + '</h4><a href="/view_hashtag?name=' + hashtag_name + '"><h4>#' + hashtag_name + '#</h4></a><p>' + post.created_date + '</p></div></div></a></div></div>'
                    } else {
                        hashtag_post_board = hashtag_post_board + '<div class="col-lg-4"><a href="#"><div class="storybox" style="background: linear-gradient( rgba(34,34,34,0.78), rgba(34,34,34,0.78)), url(users/' + post.user_id + '/posts/' + post.post_id + '/' + post.images[0].image_name + ') no-repeat;background-size: cover;background-position: center center;-webkit-background-size: cover;-moz-background-size: cover;-o-background-size: cover;"><div class="story-body text-center"><div class=""><a href="/view_users?name=' + post.user_name + '"><img class="img-circle" src="users/' + post.user_id + '/profile/' + post.avatar + '" alt="user"></div></a><h4>' + post.user_name + '</h4><a href="/view_hashtag?name=' + hashtag_name + '"><h4>#' + hashtag_name + '#</h4></a><p>' + post.created_date + '</p></div></div></a></div>'
                    }
                });
            } else {
                hashtag_post_board = '<p style="text-align: center;font-size: x-large;">#' + hashtag_name + '# do not have any posts</p>'
            }
            $('#hashtag_post_board').html(
                $('#hashtag_post_board').html() + hashtag_post_board
            )
        },
        error: function(e) {
            console.log(e);
        }
    });
}

function showOnePostsPerHashtag(hashtag_name, hashtag_id) {
    var form = new FormData();
    form.append("hashtag_name", hashtag_name);
    form.append("hashtag_id", hashtag_id);
    $.ajax({
        url: "/api/hashtag/posts",
        type: "POST",
        timeout: 0,
        processData: false,
        mimeType: "multipart/form-data",
        contentType: false,
        dataType: 'json',
        data: form,
        success: function(r) {
            var posts_num = r.posts.length;
            var home_hashtag_post_board = "";
            home_hashtag_post_board = home_hashtag_post_board + '<a href="#"><div class="storybox" style="background: linear-gradient( rgba(34,34,34,0.78), rgba(34,34,34,0.78)), url(users/' + r.posts[0].user_id + '/posts/' + r.posts[0].post_id + '/' + r.posts[0].images[0].image_name + ') no-repeat;background-size: cover;background-position: center center;-webkit-background-size: cover;-moz-background-size: cover;-o-background-size: cover;"><div class="story-body text-center"><div class=""><a href="/view_users?name=' + r.posts[0].user_name + '"><img class="img-circle" src="users/' + r.posts[0].user_id + '/profile/' + r.posts[0].avatar + '" alt="user"></a></div><h4>' + r.posts[0].user_name + '</h4><a href="/view_hashtag?name=' + hashtag_name + '"><h4>#' + hashtag_name + '#</h4></a><p>' + r.posts[0].created_date + '</p></div></div></a>'
            $('#home_hashtag_post_board').html(
                $('#home_hashtag_post_board').html() + home_hashtag_post_board
            )
        },
        error: function(e) {
            console.log(e);
        }
    });
}