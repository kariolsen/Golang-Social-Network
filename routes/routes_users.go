package routes

import (
	UT "Golang-Social-Network/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"strconv"
	"os"
	"io/ioutil"
)

func ToLogout(c *gin.Context){
	is_loggedin(c, "")
	session := UT.GetSession(c)
	delete(session.Values, "id")
	delete(session.Values, "username")
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "You have logged out",
		"success": true,
	})
}

func ToSignUp(c *gin.Context){
	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))
	password_repeated := strings.TrimSpace(c.PostForm("password_repeated"))
	email := strings.TrimSpace(c.PostForm("email"))

	if username == "" || password == "" || password_repeated == "" || email == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "You forgot some values",})
	}else if len(username) < 3 || len(username) > 32{
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Username length needs to be between 3 and 32",})
	}else if checkmail.ValidateFormat(email) != nil{
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Incorrect email",})
	}else if password != password_repeated{
		c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Passwords need to match",})
	}else{
		db := UT.Conn_DB()
		defer db.Close()
		rs, err := db.Exec("INSERT INTO Users(username, password, email) VALUES (?, ?, ?)", username, hash(password), email)
		if err != nil{
			c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "Duplicate username or password",})
		}else{
			user_id, _ := rs.LastInsertId()
			user_id_str := strconv.FormatInt(user_id,10)

			userPath := "./web/users/" + user_id_str
			err := os.MkdirAll(userPath, 0655)
			profilePath := "./web/users/" + user_id_str + "/profile"
			err = os.MkdirAll(profilePath, 0655)
			imgPath := "./web/users/" + user_id_str + "/images"
			err = os.MkdirAll(imgPath, 0655)
			postPath := "./web/users/" + user_id_str + "/posts"
			err = os.MkdirAll(postPath, 0655)
			if err != nil {c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "File Server Error",})}

			input, err := ioutil.ReadFile("./web/defaults/profile/avatar.png")
			if err != nil {c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "File Server Error",})}
			err = ioutil.WriteFile(userPath + "/profile/avatar.png", input, 0655)
			if err != nil {c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "File Server Error",})}

			session := UT.GetSession(c)
			session.Values["id"] = user_id_str
			session.Values["username"] = username
			session.Values["avatar"] = "avatar.png"
			session.Save(c.Request, c.Writer) 
			c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"message": "Welcome, " + username + " !!",
			})
		}
	}
}

func Basics(c *gin.Context){
	is_loggedin(c, "")
	my_id, my_username := UT.Get_Id_and_Username(c)
	my_avatar := UT.Get_Avatar(c)
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"id": my_id,
		"username": my_username,
		"avatar": my_avatar,
	})
}

func ToLogin(c *gin.Context){
	login_username := strings.TrimSpace(c.PostForm("username"))
	login_password := strings.TrimSpace(c.PostForm("password"))
	if login_username == "" || login_password == ""{
		panic("Please enter username and password")
	}else{
		var id int
		var count_id int
		var username string
		var password string
		var avatar string
		db := UT.Conn_DB()
		defer db.Close()
		db.QueryRow("SELECT COUNT(user_id), user_id, username, password, avatar FROM Users WHERE username = ?", login_username).Scan(&count_id, &id, &username, &password, &avatar)
		if count_id != 1{
			panic("Incorrect username or password")
		}else{
			err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login_password)) // check if hashed password match
			if err != nil{
				panic("Incorrect password")
			}else{
				session := UT.GetSession(c)
				session.Values["id"] = strconv.FormatInt(int64(id), 10)
				session.Values["username"] = username
				session.Values["avatar"] = avatar
				session.Save(c.Request, c.Writer)
				c.JSON(http.StatusOK, map[string]interface{}{
					"success": true,
					"message": username + ", you have logged in",
				})
 			}
		}
	}
}

func GetFollowers(c *gin.Context){
	is_loggedin(c, "")
	var (
		follower_id int
		follower_name string
		follower_avatar string
		following_bool bool
		my_id interface{}
		message string
	)
	followers := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	username := c.Param("userName")
	if username == ""{ // it means self
		my_id, _ = UT.Get_Id_and_Username(c)
		message = "View your followers"
	}else{ // it means others
		db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&my_id)
		message = "View "+ username +"'s followers"
	}
	stmt, err := db.Prepare("SELECT follow_by FROM Follow WHERE follow_to = ? ORDER BY created_date DESC")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}
	rows, err := stmt.Query(my_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}
	for rows.Next(){
		rows.Scan(&follower_id)
		db.QueryRow("SELECT username, avatar FROM Users WHERE user_id = ?", follower_id).Scan(&follower_name, &follower_avatar)
		db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, follower_id).Scan(&following_bool)
		follower := map[string]interface{}{
			"id": follower_id,
			"name": follower_name,
			"avatar": follower_avatar,
			"follow_relations":following_bool,
		}
		followers = append(followers, follower)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
		"success": true,
		"followers": followers,
	})
}

func GetFollowings(c *gin.Context){
	is_loggedin(c, "")
	var (
		following_id int
		following_name string
		following_avatar string
		my_id interface{}
		message string
	)
	followings := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	username := c.Param("userName")
	if username == ""{ // it means self
		my_id, _ = UT.Get_Id_and_Username(c)
		message = "View your followings"
	}else{ // it means others
		db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&my_id)
		message = "View "+ username +"'s followings"
	}
	stmt, err := db.Prepare("SELECT follow_to FROM Follow WHERE follow_by = ? ORDER BY created_date DESC")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}
	rows, err := stmt.Query(my_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}
	for rows.Next(){
		rows.Scan(&following_id)
		db.QueryRow("SELECT username, avatar FROM Users WHERE user_id = ?", following_id).Scan(&following_name, &following_avatar)
		following := map[string]interface{}{
			"id": following_id,
			"name": following_name,
			"avatar": following_avatar,
		}
		followings = append(followings, following)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
		"success": true,
		"followings": followings,
	})
}

func GetHashtags(c *gin.Context){
	is_loggedin(c, "")
	var (
		hashtag_id int
		hashtag_name string
		my_id interface{}
		message string
	)
	db := UT.Conn_DB()
	defer db.Close()
	hashtags := []interface{}{}
	username := c.Param("userName")
	if username == ""{ // it means self
		my_id, _ = UT.Get_Id_and_Username(c)
		message = "View your hashtags"
	}else{ // it means others
		db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&my_id)
		message = "View "+ username +"'s hashtags"
	}
	stmt, err := db.Prepare("SELECT hashtag_id FROM Users_Hashtags WHERE user_id = ? ORDER BY created_date DESC")
	UT.Err(err)
	rows, err := stmt.Query(my_id)
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&hashtag_id)
		db.QueryRow("SELECT hashtag_name FROM Hashtags WHERE hashtag_id = ?", hashtag_id).Scan(&hashtag_name)
		hashtag := map[string]interface{}{
			"id": hashtag_id,
			"name": hashtag_name,
		}
		hashtags = append(hashtags, hashtag)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": message,
		"success": true,
		"hashtags": hashtags,
	})
}

func FollowOrUnfollowUser(c *gin.Context){
	is_loggedin(c, "")
	db := UT.Conn_DB()
	defer db.Close()
	var (
		follow_relations bool
		res bool
	)
	my_id, _ := UT.Get_Id_and_Username(c)
	user_id := c.Param("userID")
	if my_id != user_id {
		db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, user_id).Scan(&follow_relations)
		if follow_relations == true {
			res = UnFollowUser(my_id, user_id)
		}else{
			res = FollowUser(my_id, user_id)
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
}

func FollowUser(my_id interface{}, target_id interface{}) bool{
	db := UT.Conn_DB()
	defer db.Close()
	if my_id != target_id{
		_, err := db.Exec("INSERT INTO Follow (follow_by, follow_to) VALUES(?, ?)", my_id, target_id)
		if err != nil{return false}
	}
	return true
}

func UnFollowUser(my_id interface{}, target_id interface{}) bool{
	db := UT.Conn_DB()
	defer db.Close()
	if my_id != target_id{
		_, err := db.Exec("DELETE FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, target_id)
		if err != nil{return false}
	}
	return true
}

func BlockUser(c *gin.Context){
	is_loggedin(c, "")
	var (
		target_id int
	)
	db := UT.Conn_DB()
	defer db.Close()
	username := c.Param("userName")
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&target_id)
	if target_id == 0 {panic("Invalid username")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	if my_id != target_id{
		_, err := db.Exec("INSERT INTO Blacklist (black_by, black_to) VALUES(?, ?)", my_id, target_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Blocked "+username+" successfully",
			"success": true,
		})
	}else{panic("You cannot block yourself")}
}

func UnBlockUser(c *gin.Context){
	is_loggedin(c, "")
	var (
		target_id int
	)
	db := UT.Conn_DB()
	defer db.Close()
	username := c.Param("userName")
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&target_id)
	if target_id == 0 {panic("Invalid username")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	if my_id != target_id{
		_, err := db.Exec("DELETE FROM Blacklist WHERE black_by = ? AND black_to = ?", my_id, target_id)
		UT.Err(err)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unblocked "+username+" successfully",
			"success": true,
		})
	}else{panic("You cannot unblock yourself")}	
}

func ShowHottestUsers(c *gin.Context){
	is_loggedin(c, "")
	my_id, _ := UT.Get_Id_and_Username(c)
	var (
		user_id int
		user_likes int
		user_name string
		user_avatar string
		latest_post_id int
		latest_image_name string
	)
	hottest_users := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, err := db.Prepare("SELECT created_by, sum(likes) likes FROM Posts WHERE created_by NOT IN (SELECT black_by FROM Blacklist WHERE black_to = ? UNION SELECT follow_to FROM Follow WHERE follow_by = ?) AND created_by != ? GROUP BY created_by ORDER BY likes DESC LIMIT 10")
	UT.Err(err)
	rows, err := stmt.Query(my_id, my_id, my_id)
	UT.Err(err)
	for rows.Next(){
		rows.Scan(&user_id, &user_likes)
		db.QueryRow("SELECT username, avatar, post_id, image_name from Images INNER JOIN Users using(user_id) where user_id = ? ORDER BY created_date DESC LIMIT 1", user_id).Scan(&user_name, &user_avatar, &latest_post_id, &latest_image_name)
		user := map[string]interface{}{
			"id": user_id,
			"name": user_name,
			"avatar": user_avatar,
			"likes": user_likes,
			"latest_post": latest_post_id,
			"latest_image": latest_image_name,
		}
		hottest_users = append(hottest_users, user)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hottest Users List",
		"success": true,
		"users": hottest_users,
	})
}

func GetUserID(c *gin.Context){
	is_loggedin(c, "")
	var (
		user_id int
	)
	username := c.Param("userName")
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&user_id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User ID",
		"success": true,
		"user_id": user_id,
	})
}