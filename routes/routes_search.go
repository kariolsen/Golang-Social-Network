package routes

import (
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
)


func SearchContent(c *gin.Context){
	is_loggedin(c, "")
	db := UT.Conn_DB()
	defer db.Close()
	var(
		user_name string
		user_id int
		user_avatar string
		hashtag_name string
		hashtag_id int
	)
	users := []interface{}{}
	hashtags := []interface{}{}
	search_words := strings.Fields(c.PostForm("search"))  

	for _, word := range search_words {
		stmt, err := db.Prepare("SELECT username, user_id, avatar FROM Users WHERE username LIKE ?")
		if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
		rows, err := stmt.Query("%"+word+"%")
		if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
		for rows.Next(){
			rows.Scan(&user_name, &user_id, &user_avatar)
			user := map[string]interface{}{
				"user_id": user_id,
				"user_name": user_name,
				"user_avatar": user_avatar,
			}
			users = append(users, user)
		}
		
		stmt, err = db.Prepare("SELECT hashtag_name, hashtag_id FROM Hashtags WHERE hashtag_name LIKE ?")
		if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
		rows, err = stmt.Query("%"+word+"%")
		if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "DB Error","success": false,})}
		for rows.Next(){
			rows.Scan(&hashtag_name, &hashtag_id)
			hashtag := map[string]interface{}{
				"hashtag_id": hashtag_id,
				"hashtag_name": hashtag_name,
			}
			hashtags = append(hashtags, hashtag)
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successful search",
		"success": true,
		"users": users,
		"hashtags": hashtags,
	});
}