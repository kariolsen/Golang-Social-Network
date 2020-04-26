package routes

import (
	UT "Golang-Social-Network/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"os"
)

func CreateImages(c *gin.Context, my_id interface{}, post_id int64) bool{
	form := c.Request.MultipartForm
	files := form.File["images"]
	img_sql := "INSERT INTO Images (post_id, user_id, image_name, image_size) VALUES "
	for _, file := range files {
		//fmt.Println("name:", file.Filename, file.Size)
		//fmt.Println("ext:", filepath.Ext(fileHeader.Filename))
		//fileType := "other"
		//fileExt := filepath.Ext(fileHeader.Filename)
		//if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		//    fileType = "img"
		//}
		fileDir := "./web/users/" + my_id.(string) + "/posts/" + strconv.FormatInt(post_id, 10)
		err := os.MkdirAll(fileDir, 0655)
		if err != nil{return false}
		filePathStr := filepath.Join(fileDir, file.Filename)
		c.SaveUploadedFile(file,filePathStr)
		img_sql = img_sql + "( " + strconv.FormatInt(post_id, 10) + ", " + my_id.(string) + ", " + "'" + file.Filename + "'" + ", " + strconv.FormatInt(file.Size, 10) +  " ),"
	}
	img_sql = img_sql[0: len(img_sql)-1]
	db := UT.Conn_DB()
	defer db.Close()
	_, err := db.Exec(img_sql)
	if err != nil{return false}
	return true
}

func ShowImages(c *gin.Context){
	is_loggedin(c, "")
	id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	var (
		image_name string
		post_id int
	)
	images := []interface{}{}
	stmt, err := db.Prepare("SELECT image_name, post_id FROM Images WHERE user_id = ? ORDER BY created_date DESC LIMIT 30")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false,})}
	rows, err := stmt.Query(id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false,})}
	for rows.Next(){
		rows.Scan(&image_name, &post_id)
		image := map[string]interface{}{
			"image_name": image_name,
			"post_id": post_id,
			"user_id": id,
		}
		images = append(images, image)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Images successfully retrieved",
		"success": true,
		"images": images,
	})
}

func GetHottestImages(c *gin.Context){
	is_loggedin(c, "")
	id, _ := UT.Get_Id_and_Username(c)
	db := UT.Conn_DB()
	defer db.Close()
	var (
		post_id int
		user_id string
		image_name string
	)
	images := []interface{}{}
	stmt, err := db.Prepare("SELECT Posts.post_id, Posts.created_by, Images.Image_name FROM Posts INNER JOIN Images USING (post_id) WHERE Posts.created_by != ? GROUP BY (Posts.post_id) ORDER BY Posts.created_date DESC, Posts.likes DESC LIMIT 20")
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false,})}
	rows, err := stmt.Query(id)
	if err != nil{c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false,})}
	for rows.Next(){
		rows.Scan(&post_id, &user_id, &image_name)
		image := map[string]interface{}{
			"image_name": image_name,
			"post_id": post_id,
			"user_id": user_id,
		}
		images = append(images, image)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hottest Images successfully retrieved",
		"success": true,
		"images": images,
	})
}

func ShowProfileImages(target_id interface{}, images_num int) interface{}{
	db := UT.Conn_DB()
	defer db.Close()
	var (
		image_name string
		post_id int
	)
	images := []interface{}{}
	stmt, err := db.Prepare("SELECT image_name, post_id FROM Images WHERE user_id = ? ORDER BY created_date DESC LIMIT ?")
	if err != nil{return false}
	rows, err := stmt.Query(target_id, images_num)
	if err != nil{return false}
	for rows.Next(){
		rows.Scan(&image_name, &post_id)
		image := map[string]interface{}{
			"image_name": image_name,
			"post_id": post_id,
			"user_id": target_id,
		}
		images = append(images, image)
	}
	return images
}

func ShowHashtagImages(hashtag_id interface{}, images_num int) interface{}{
	db := UT.Conn_DB()
	defer db.Close()
	var (
		image_name string
		user_id int
		post_id int
	)
	images := []interface{}{}
	stmt, err := db.Prepare("SELECT image_name, user_id, Posts_Hashtags.post_id FROM Images INNER JOIN Posts_Hashtags USING (post_id) WHERE Posts_Hashtags.hashtag_id = ? ORDER BY Posts_Hashtags.created_date DESC LIMIT ?")
	if err != nil{return false}
	rows, err := stmt.Query(hashtag_id, images_num)
	if err != nil{return false}
	for rows.Next(){
		rows.Scan(&image_name, &user_id, &post_id)
		image := map[string]interface{}{
			"image_name": image_name,
			"post_id": post_id,
			"user_id": user_id,
		}
		images = append(images, image)
	}
	return images
}

func ShowPostImages(c *gin.Context, post_id interface{}, user_id interface{}) []interface{}{
	var image_name string
	images := []interface{}{}
	db := UT.Conn_DB()
	defer db.Close()
	stmt, _ := db.Prepare("SELECT image_name FROM Images WHERE user_id = ? AND post_id = ? ORDER BY created_date DESC")
	rows, _ := stmt.Query(user_id, post_id)
	for rows.Next(){
		rows.Scan(&image_name)
		image := map[string]interface{}{
			"image_name": image_name,
		}
		images = append(images, image)
	}
	return images
}