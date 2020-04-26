package routes

import(
	UT "Golang-Social-Network/utils"
)

func Create_Mention(post_id interface{}, username string){
	var(
		userCount int
		user_id int
		mentionCount int
	)
	db := UT.Conn_DB()
	defer db.Close()
	db.QueryRow("SELECT COUNT(user_id), user_id FROM Users WHERE username = ?", username).Scan(&userCount, &user_id)
	if userCount != 1 || user_id == 0 {panic("Invalid username")}
	db.QueryRow("SELECT COUNT(*) FROM Mentions WHERE user_id = ? AND post_id = ?", user_id, post_id).Scan(&mentionCount)
	if mentionCount == 0{
		stmt, err := db.Prepare("INSERT INTO Mentions (user_id, post_id) VALUES(?, ?)")
		UT.Err(err)
		stmt.Exec(user_id, post_id)
	}
}