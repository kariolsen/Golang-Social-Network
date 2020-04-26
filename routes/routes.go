package routes

import (
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"net/http"
)

func hash(password string) []byte{ //encrypt passwords
	hashed_pwd, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	UT.Err(hashErr)
	return hashed_pwd
}

func is_loggedin(c *gin.Context, urlRedirect string){
	var URL string
	if urlRedirect == ""{URL = "/login"}else{URL = urlRedirect}
	id, _ := UT.Get_Id_and_Username(c)
	if id == nil{
		c.Redirect(http.StatusFound, URL)
		c.JSON(http.StatusResetContent, map[string]interface{}{
			"message": "Please login First",
			"success": false,
		})
		panic("Please login First")
	}
}

func is_not_loggedin(c *gin.Context){
	id, _ := UT.Get_Id_and_Username(c)
	if id != nil {
		c.Redirect(http.StatusFound, "/")
		c.JSON(http.StatusResetContent, map[string]interface{}{
			"message": "Please log out first",
			"success": false,
		})
		panic("Please log out first")
	}
}

func id_name_in_json(c *gin.Context) interface{} {
	id, username := UT.Get_Id_and_Username(c)
	return map[string]interface{}{
		"id":       id,
		"username": username,
	}
}

func is_Blacked(my_id interface{}, target_id interface{}) bool{
	var res int
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(*) FROM Blacklist WHERE black_by = ? AND black_to = ?", target_id, my_id).Scan(&res) // is my_id blocked by target_id ?
	if res > 0{return true} else {return false}
}

func is_Following(my_id interface{}, target_id interface{}) bool{
	var res int
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(*) FROM Follow WHERE follow_by = ? AND follow_to = ?", my_id, target_id).Scan(&res) // is my_id blocked by target_id ?
	if res > 0{return true} else {return false}
}

func open_for_Unfollowers(target_id interface{}) bool{
	var(
		resCount int
		res bool
	)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(*), allow_unfollowed_views FROM Profile WHERE user_id = ?", target_id).Scan(&resCount, &res)
	if resCount != 1 {panic("Invalid user ID")} else { return res }
}

func extractTags_Mentions(content string) ([]string, []string) {
	var tags []string
	var mentions []string
	words := strings.Fields(content)  
	for _, word := range words{
		if strings.HasPrefix(word, "#") && strings.HasSuffix(word, "#"){
			word = strings.TrimLeft(word, "#")
			word = strings.TrimRight(word, "#")
			tags = append(tags, word)
		}else if strings.HasPrefix(word, "@"){
			word = strings.TrimLeft(word, "@")
			mentions = append(mentions, word)
		}
	}
	return tags, mentions
}
