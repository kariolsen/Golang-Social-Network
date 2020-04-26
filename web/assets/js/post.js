function labelEssentials(content) {
    if (content == "") {
        return "Empty Content";
    }
    var word_arr = content.split(" ");
    var res = "";
    var hashtag_name = "";
    var user_name = "";
    for (let i = 0; i < word_arr.length; i++) { // hashtag
        if (word_arr[i].startsWith('#') && word_arr[i].endsWith('#')) {
            hashtag_name = word_arr[i].slice(1, -1);
            res = res + '<a href="/view_hashtag?name=' + hashtag_name + '"><span class="special_content">#' + hashtag_name + '#</span></a>';
            res = res + ' ';
        } else if (word_arr[i].startsWith('@')) { // @
            user_name = word_arr[i].slice(1);
            res = res + '<a href="/view_users?name=' + user_name + '"><span class="special_content">@' + user_name + '</span></a>';
            res = res + ' ';
        } else {
            res = res + word_arr[i];
            res = res + ' ';
        }
    }
    return res;
}