package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	is_not_loggedin(c)
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title": "Signup For Free",
	})
}

func Login(c *gin.Context) {
	is_not_loggedin(c)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Log in",
	})
}

func Home(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
	})
}

func Exploring(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "explore.html", gin.H{
		"title": "Explore friends moments",
	})
}

func Stories(c *gin.Context){
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "stories.html", gin.H{
		"title": "Stories",
	})
}

func ViewProfile(c *gin.Context){
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"title": "Profile",
	})
}

func ViewUserProfile(c *gin.Context){
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "user_profile.html", gin.H{
		"title": "User Profile",
	})
}

func ViewPost(c *gin.Context){
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "post.html", gin.H{
		"title": "This Post",
	})
}

func Followers(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "followers.html", gin.H{
		"title": "Followers",
	})
}

func UserFollowers(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "user_followers.html", gin.H{
		"title": "User Followers",
	})
}

func HashtagFollowers(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "hashtag_followers.html", gin.H{
		"title": "Hashtag Followers",
	})
}

func Followings(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "followings.html", gin.H{
		"title": "Followings",
	})
}

func UserFollowings(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "user_followings.html", gin.H{
		"title": "User Followings",
	})
}

func Upload(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"title": "Upload",
	})
}

func Search(c *gin.Context) {
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "search.html", gin.H{
		"title": "Search",
	})
}

func ViewHashtag(c *gin.Context){
	is_loggedin(c, "")
	c.HTML(http.StatusOK, "hashtag.html", gin.H{
		"title": "Hashtag Posts",
	})
}