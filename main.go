package main

import (
	route "Golang-Social-Network/routes"
	"github.com/gin-gonic/gin"
	"github.com/urfave/negroni"
	"net/http"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("web/*.html")
	router.StaticFS("/assets", http.Dir("web/assets"))
	router.StaticFS("/users", http.Dir("web/users"))

	user := router.Group("/user")
	{
		user.POST("/signup", route.ToSignUp)
		user.POST("/login", route.ToLogin)
		user.POST("/logout", route.ToLogout)
	}

	router.GET("/basics", route.Basics)
	router.GET("/signup", route.Signup)
	router.GET("/login", route.Login)
	router.GET("/home", route.Home) // direct to page that shows the posts of followings order by created_date only
	router.GET("/upload", route.Upload) // direct to page that can upload posts
	router.GET("/exploring", route.Exploring) // direct to page that shows the most popular posts and images of all users (not limited to your followings)
	router.GET("/stories", route.Stories) // direct to page that 
	router.GET("/view_post", route.ViewPost) // direct to page that display your profile info
	router.GET("/view_users", route.ViewUserProfile) // direct to page that display other user's profile info
	router.GET("/view_profile", route.ViewProfile) // direct to page that display your profile info
	router.GET("/view_hashtag", route.ViewHashtag)
	router.GET("/followers", route.Followers)
	router.GET("/followings", route.Followings)
	router.GET("/hashtag_followers", route.HashtagFollowers)
	router.GET("/user_followers", route.UserFollowers)
	router.GET("/user_followings", route.UserFollowings)
	router.GET("/search", route.Search)

	api := router.Group("/api")
	{
		api.GET("/explore", route.Explore)
		api.GET("/explore/hashtag/posts/all/:hashtagname", route.ExploreAllHashtagPosts)
		api.GET("/explore/hashtag/posts/following/:hashtagname", route.ExploreFriendsHashtagPosts)

		api.GET("/post/popular", route.ShowHottestPosts)
		api.POST("/post/add", route.CreatePost)
		api.POST("/post/delete/:postID", route.DeletePost)
		api.POST("/post/edit/:postID", route.UpdatePost)

		api.POST("/post/likePressed/:postID", route.LikeOrUnlike)

		api.GET("/images", route.ShowImages)
		api.GET("/images/popular", route.GetHottestImages)

		api.POST("/search_content", route.SearchContent)

		api.POST("/comments/add/:postID", route.CreateComments)
		api.POST("/comments/edit/:commentID", route.EditComments)
		api.POST("/comments/like/:commentID", route.LikeComments)
		api.POST("/comments/unlike/:commentID", route.UnlikeComments)
		api.POST("/comments/delete/:commentID", route.DeleteComments)
		api.GET("/comments/show/:postID", route.DisplayComments)

		api.GET("/user/popular", route.ShowHottestUsers)
		api.GET("/user/followers", route.GetFollowers)
		api.GET("/user/followings", route.GetFollowings)
		api.GET("/user/hashtags", route.GetHashtags)
		api.GET("/user/followers/:userName", route.GetFollowers)
		api.GET("/user/followings/:userName", route.GetFollowings)
		api.GET("/user/hashtags/:userName", route.GetHashtags)

		api.POST("/user/followPressed/:userID", route.FollowOrUnfollowUser)

		api.POST("/user/blacklist/:userName", route.BlockUser)
		api.POST("/user/unblacklist/:userName", route.UnBlockUser)
		api.POST("/user/ID/:userName", route.GetUserID)

		api.GET("/profile/:id", route.Profile)
		api.POST("/profile", route.EditProfile)
		api.POST("/profile_avatar", route.EditProfileAvatar)

//		api.POST("/follow_topic", route.FollowTopic)
//		api.POST("/unfollow_topic", route.UnFollowTopic)
		api.GET("/hashtag/display/:name", route.DisplayHashtagPosts) // including posts and comments in detail
		api.GET("/hashtag/popular", route.ShowHottestHashtags)
		api.GET("/hashtag/followers/:hashtagName", route.GetHashtagsFollowers)
		api.GET("/hashtag/following", route.GetFollowingHashtags)
		api.POST("/hashtag/posts", route.GetHashtagPosts) // only posts
		api.POST("/hashtag/followPressed/:hashtagID", route.FollowOrUnfollowHashtag)

		api.GET("/notifications/", route.GetNotifications)
		api.POST("/notifications/clear", route.ClearNotifications)
	}
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":8882")
}