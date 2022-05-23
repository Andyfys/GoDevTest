/*
评论：包含发送评论，获取评论列表
comment： include send comment, get all comments in the video
*/

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

/*
判断了用户是否登录
这里暂时可以不做修改

*/

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

/*
这里是获取评论的主要部分。
在这里
CommentList: DemoComments
这句话会把一个评论的列表赋给CommentList进行返回
所以，对于实现评论功能，主要的问题点在于如何实现列表的返回
*/

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
