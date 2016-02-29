package request

type CommentFormRequest struct {
	Content string `form:"content"`
}
