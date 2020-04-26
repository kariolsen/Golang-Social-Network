package routes

import(
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Check_HashTag_Exist(hashtag_name string) (int, bool){ // return hashtag id and if hashtag exists
	db := UT.Conn_DB()
	defer db.Close()
	var (
		hashtagCount int
		hashtagID int
	)
	db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtagCount, &hashtagID)
	if hashtagCount == 0{
		_, err := db.Exec("INSERT INTO Hashtags (hashtag_name) VALUES(?)", hashtag_name)
		UT.Err(err)
		db.QueryRow("SELECT COUNT(hashtag_id), hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtagCount, &hashtagID)
		if hashtagCount == 1 {
			return hashtagID, true
		}else{panic("Database Errors")}
	}else if hashtagCount == 1{
		return hashtagID, true
	}else{
		return 0, false
	}
}

func Create_Follow_HashTag(post_id interface{}, hashtag_name string) (int, bool){ // This function is called as part of the CreatePost function
	db := UT.Conn_DB()
	defer db.Close()
	var hashtagCount int
	if hashtag_name != ""{ // hashtag is present
		hashtag_id, hashtag_err := Check_HashTag_Exist(hashtag_name)
		if hashtag_err == false{panic("Database Errors")}else{
			db.QueryRow("SELECT COUNT(*) FROM Posts_Hashtags WHERE hashtag_id = ? AND post_id = ?", hashtag_id, post_id).Scan(&hashtagCount)
			if hashtagCount == 0{
				_, err := db.Exec("INSERT INTO Posts_Hashtags (hashtag_id, post_id) VALUES(?, ?)", hashtag_id, post_id)
				UT.Err(err)
				return hashtag_id, true
			}else{return hashtag_id, true}
		}
	}else{return 0, false}
}

func FollowOrUnfollowHashtag(c *gin.Context){
	is_loggedin(c, "")
	db := UT.Conn_DB()
	defer db.Close()
	var (
		follow_relations bool
		res bool
	)
	my_id, _ := UT.Get_Id_and_Username(c)
	hashtag_id := c.Param("hashtagID")

	db.QueryRow("SELECT COUNT(*) FROM Users_Hashtags WHERE user_id = ? AND hashtag_id = ?", my_id, hashtag_id).Scan(&follow_relations)
	if follow_relations == true {
		res = UnFollowHashTag(my_id, hashtag_id)
	}else{
		res = FollowHashTag(my_id, hashtag_id)
	}
	if res == true{
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success",
			"success": true,
		});
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})
	}
	
}

func FollowHashTag(my_id interface{}, hashtag_id interface{}) bool{
	db := UT.Conn_DB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO Users_Hashtags (user_id, hashtag_id) VALUES(?, ?)", my_id, hashtag_id)
	if err != nil {return false}
	return true
}

func UnFollowHashTag(my_id interface{}, hashtag_id interface{}) bool{
	db := UT.Conn_DB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM Users_Hashtags WHERE user_id = ? AND hashtag_id = ?", my_id, hashtag_id)
	if err != nil {return false}
	return true
}

func ShowHottestHashtags(c *gin.Context){
	is_loggedin(c, "")
	var (
		hashtag_id int
		hashtag_name string
		followers_num int
		posts_num int
		created_date string
	)
	hottest_hashtags := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT hashtag_id, hashtag_name, followers_num, posts_num, DATE(created_date) FROM Hashtags ORDER BY followers_num DESC, posts_num DESC, created_date DESC LIMIT 10")
	UT.Err(err)
	rows, err := stmt.Query()
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&hashtag_id, &hashtag_name, &followers_num, &posts_num, &created_date)
		hashtag := map[string]interface{}{
			"hashtag_id": hashtag_id,
			"hashtag_name": hashtag_name,
			"followers_num": followers_num,
			"posts_num": posts_num,
			"created_date": created_date,
		}
		hottest_hashtags = append(hottest_hashtags, hashtag)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hottest Hashtags List",
		"success": true,
		"hashtags": hottest_hashtags,
	})
}

func GetFollowingHashtags(c *gin.Context){
	is_loggedin(c, "")
	my_id, _ := UT.Get_Id_and_Username(c)
	var (
		hashtag_id int
		hashtag_name string
		followers_num int
		posts_num int
	)
	following_hashtags := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT Users_Hashtags.hashtag_id, Hashtags.hashtag_name, Hashtags.followers_num, Hashtags.posts_num FROM Users_Hashtags INNER JOIN Hashtags USING(hashtag_id) WHERE user_id = ? LIMIT 10")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
	rows, err := stmt.Query(my_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
	for rows.Next(){
		rows.Scan(&hashtag_id, &hashtag_name, &followers_num, &posts_num)
		hashtag := map[string]interface{}{
			"hashtag_id": hashtag_id,
			"hashtag_name": hashtag_name,
			"followers_num": followers_num,
			"posts_num": posts_num,
		}
		following_hashtags = append(following_hashtags, hashtag)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Following Hashtags List",
		"success": true,
		"hashtags": following_hashtags,
	})
}

func GetHashtagPosts(c *gin.Context){
	is_loggedin(c, "")
	hashtag_id := strings.TrimSpace(c.PostForm("hashtag_id"))
	hashtag_name := strings.TrimSpace(c.PostForm("hashtag_name"))

	posts := []interface{}{}
	var (
		post_id int
		user_id int
		user_name string
		avatar string
		title string
		content string
		likes int
		allow_comments bool
		comments_number int
		created_date string
	)

	db := UT.Conn_DB()
	defer db.Close()
	stmt, _ := db.Prepare("SELECT Users.username, Users.avatar, Posts.created_by, Posts.post_id, Posts.title, Posts.content, Posts.likes, Posts.allow_comments, Posts.comments_num, DATE(Posts.created_date) FROM Posts_Hashtags INNER JOIN Posts USING (post_id) INNER JOIN Users ON Posts.created_by = Users.user_id WHERE Posts_Hashtags.hashtag_id = ?")
	rows, _ := stmt.Query(hashtag_id)
	for rows.Next(){
		rows.Scan(&user_name, &avatar, &user_id, &post_id, &title, &content, &likes, &allow_comments, &comments_number, &created_date)
		if allow_comments == true{
			post := map[string]interface{}{
				"post_id": post_id,
				"user_id": user_id,
				"user_name": user_name,
				"avatar": avatar,
				"title": title, 
				"content": content,
				"likes": likes,
				"allow_comments": allow_comments,
				"comments": ShowComments(c, post_id),
				"comments_num": comments_number,
				"images": ShowPostImages(c, post_id, user_id),
				"created_date": created_date,
			}
			posts = append(posts, post)
		}else{
			post := map[string]interface{}{
				"post_id": post_id,
				"user_id": user_id,
				"user_name": user_name,
				"avatar": avatar,
				"title": title, 
				"content": content,
				"likes": likes,
				"allow_comments": allow_comments,
				"comments": allow_comments,
				"comments_num": 0,
				"images": ShowPostImages(c, post_id, user_id),
				"created_date": created_date,
			}
			posts = append(posts, post)
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Retrieved posts of hashtag "+hashtag_name,
		"success": true,
		"hashtag_name": hashtag_name,
		"hashtag_id": hashtag_id,
		"posts": posts,
	})
}

func DisplayHashtagPosts(c *gin.Context){
	is_loggedin(c, "")
	hashtag_name := c.Param("name") // id is part of url
	if hashtag_name != ""{
		my_id, _ := UT.Get_Id_and_Username(c)
		c.JSON(http.StatusOK, HashtagPostsComments(hashtag_name, my_id, c))
	}else{
		c.JSON(http.StatusOK, map[string]interface{}{"success": true,})
	}
}

func HashtagPostsComments(hashtag_name string, my_id interface{}, c *gin.Context) map[string]interface{}{
	db := UT.Conn_DB()
	defer db.Close()
	posts := []interface{}{}

	var (
		hashtag_id int
		hashtag_followers int
		hashtag_posts_num int
		post_id int
		user_id int
		user_name string
		avatar string
		title string
		content string
		likes int
		allow_comments bool
		comments_number int
		created_date string
		following_hashtag bool
		liked_by_you bool
		followed_by_you bool
		blocked bool
	)

	db.QueryRow("SELECT hashtag_id, followers_num, posts_num FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtag_id, &hashtag_followers, &hashtag_posts_num)
	db.QueryRow("SELECT COUNT(*) FROM Users_Hashtags WHERE user_id = ? AND hashtag_id = ?", my_id, hashtag_id).Scan(&following_hashtag)
	stmt, _ := db.Prepare("SELECT Users.username, Users.avatar, Posts.created_by, Posts.post_id, Posts.title, Posts.content, Posts.likes, Posts.allow_comments, Posts.comments_num, DATE(Posts.created_date) FROM Posts_Hashtags INNER JOIN Posts USING (post_id) INNER JOIN Users ON Posts.created_by = Users.user_id WHERE Posts_Hashtags.hashtag_id = ?")
	rows, _ := stmt.Query(hashtag_id)
	for rows.Next(){
		rows.Scan(&user_name, &avatar, &user_id, &post_id, &title, &content, &likes, &allow_comments, &comments_number, &created_date)
		db.QueryRow("SELECT COUNT(*) FROM Blacklist WHERE black_by = ? AND black_to = ?", user_id, my_id).Scan(&blocked)
		db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, user_id).Scan(&followed_by_you)
		db.QueryRow("SELECT COUNT(*) FROM Likes WHERE post_id = ? AND like_by = ?", post_id, my_id).Scan(&liked_by_you)
		if (blocked == false){
			if allow_comments == true{
				post := map[string]interface{}{
					"post_id": post_id,
					"user_id": user_id,
					"user_name": user_name,
					"avatar": avatar,
					"title": title, 
					"content": content,
					"likes": likes,
					"allow_comments": allow_comments,
					"comments": ShowComments(c, post_id),
					"comments_num": comments_number,
					"images": ShowPostImages(c, post_id, user_id),
					"created_date": created_date,
					"followed_by_you": followed_by_you,
					"liked_by_you": liked_by_you,
				}
				posts = append(posts, post)
			}else{
				post := map[string]interface{}{
					"post_id": post_id,
					"user_id": user_id,
					"user_name": user_name,
					"avatar": avatar,
					"title": title, 
					"content": content,
					"likes": likes,
					"allow_comments": allow_comments,
					"comments": allow_comments,
					"comments_num": 0,
					"images": ShowPostImages(c, post_id, user_id),
					"created_date": created_date,
					"followed_by_you": followed_by_you,
					"liked_by_you": liked_by_you,
				}
				posts = append(posts, post)
			}
		}
	}
	return map[string]interface{}{
		"message": "Found posts of hashtag" + hashtag_name,
		"success": true,
		"hashtag_id": hashtag_id,
		"hashtag_followers": hashtag_followers, 
		"hashtag_posts_num": hashtag_posts_num,
		"following_hashtag": following_hashtag, 
		"posts": posts,
		"profile_bg_images": ShowHashtagImages(hashtag_id, 5),
	}
}

func GetHashtagsFollowers(c *gin.Context){
	is_loggedin(c, "")
	var (
		follower_id int
		hashtag_id int
		follower_name string
		follower_avatar string
		my_id interface{}
		message string
		following_bool bool
	)
	followers := []interface{}{}
	my_id, _ = UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	hashtag_name := c.Param("hashtagName")
	message = "View " + hashtag_name + " followers"
	db.QueryRow("SELECT hashtag_id FROM Hashtags WHERE hashtag_name = ?", hashtag_name).Scan(&hashtag_id)

	stmt, err := db.Prepare("SELECT Users.user_id, Users.username, Users.avatar FROM Users_Hashtags INNER JOIN Users USING (user_id) WHERE hashtag_id = ? ORDER BY Users_Hashtags.created_date DESC")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}
	rows, err := stmt.Query(hashtag_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}

	for rows.Next(){
		rows.Scan(&follower_id, &follower_name, &follower_avatar)
		db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, follower_id).Scan(&following_bool)
		follower := map[string]interface{}{
			"id": follower_id,
			"name": follower_name,
			"avatar": follower_avatar,
			"follow_relations": following_bool,
		}
		followers = append(followers, follower)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
		"success": true,
		"followers": followers,
	})
}