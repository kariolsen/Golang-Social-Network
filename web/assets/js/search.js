function setSearchResults() {
    var search_content = $("#search_content").val();
    if (search_content != "") {
        var form = new FormData();
        form.append("search", search_content);

        $.ajax({
            url: "/api/search_content",
            type: "POST",
            timeout: 0,
            processData: false,
            mimeType: "multipart/form-data",
            contentType: false,
            dataType: 'json',
            data: form,
            success: function(r) {
                localStorage.removeItem("search_res");
                localStorage.setItem("search_res", JSON.stringify(r));
                if (window.location.pathname != "/search") {
                    window.location.replace("/search");
                }
            },
            error: function(e) {
                notifyUser('Error!', 'Search Failed !', 'danger', 3000);
            }
        });
    }
}


function setSearchMainResults() {
    var search_content = $("#search_content_main").val();
    if (search_content != "") {
        var form = new FormData();
        form.append("search", search_content);

        $.ajax({
            url: "/api/search_content",
            type: "POST",
            timeout: 0,
            processData: false,
            mimeType: "multipart/form-data",
            contentType: false,
            dataType: 'json',
            data: form,
            success: function(r) {
                localStorage.removeItem("search_res");
                localStorage.setItem("search_res", JSON.stringify(r));
                if (window.location.pathname != "/search") {
                    window.location.replace("/search");
                }
            },
            error: function(e) {
                notifyUser('Error!', 'Search Failed !', 'danger', 3000);
            }
        });
    }
}