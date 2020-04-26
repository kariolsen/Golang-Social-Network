package routes

import(
	UT "Golang-Social-Network/utils"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowTopic(c *gin.Context) {
	is_loggedin(c, "")
	var (
		topic_id int
		topic_count int
	)
	db := UT.Conn_DB()
	defer db.Close()
	topic_name := strings.TrimSpace(c.PostForm("topic_name"))
	db.QueryRow("SELECT COUNT(*), topic_id FROM Topics WHERE topic_name = ?", topic_name).Scan(&topic_count, &topic_id)
	if topic_count != 1 || topic_id == 0 {panic("Invalid topic name")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	
	_, err := db.Exec("INSERT INTO Users_Topics (user_id, topic_id) VALUES(?, ?)", my_id, topic_id)
	UT.Err(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Followed topic "+ topic_name +" successfully",
		"success": true,
	})
}

func UnFollowTopic(c *gin.Context){
	is_loggedin(c, "")
	var (
		topic_id int
		topic_count int
	)
	db := UT.Conn_DB()
	defer db.Close()
	topic_name := strings.TrimSpace(c.PostForm("topic_name"))
	db.QueryRow("SELECT COUNT(*), topic_id FROM Topics WHERE topic_name = ?", topic_name).Scan(&topic_count, &topic_id)
	if topic_count != 1 || topic_id == 0 {panic("Invalid topic name")}
	my_id, _ := UT.Get_Id_and_Username(c)
	if my_id == 0 {panic("Invalid user id")}
	_, err := db.Exec("DELETE FROM Users_Topics WHERE user_id = ? AND topic_id = ?", my_id, topic_id)
	UT.Err(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Unfollowed topic "+ topic_name +" successfully",
		"success": true,
	})

}