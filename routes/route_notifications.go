package routes

import(
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotifications(c *gin.Context){
	is_loggedin(c, "")
	var (
		user_id int
		user_name string
		user_avatar string
		created_date string
		noti_type int
		post_title string
	)
	mentions := []interface{}{}
	my_id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()

	stmt, err := db.Prepare("SELECT Users.username, Users.user_id, Users.avatar, DATE(Mentions.created_date), Posts.title, Mentions.type FROM Mentions INNER JOIN Posts USING (post_id) INNER JOIN Users ON Posts.created_by = Users.user_id WHERE Mentions.user_id = ? AND Users.user_id != ? AND Mentions.viewed = false ORDER BY Mentions.created_date DESC")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}
	rows, err := stmt.Query(my_id, my_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}

	for rows.Next(){
		rows.Scan(&user_name, &user_id, &user_avatar, &created_date, &post_title, &noti_type)
		mention := map[string]interface{}{
			"user_id": user_id,
			"user_name": user_name,
			"user_avatar": user_avatar,
			"created_date": created_date,
			"title": post_title,
			"type": noti_type,
		}
		mentions = append(mentions, mention)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Retrieved notifications",
		"success": true,
		"notifications": mentions,
	});
}

func ClearNotifications(c *gin.Context){
	is_loggedin(c, "")
	my_id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	_, err := db.Exec("UPDATE Mentions SET viewed = ? WHERE user_id = ?", true, my_id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": "DB Error",})}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Cleared notifications",
		"success": true,
	});
}