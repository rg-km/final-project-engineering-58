package request

type CreatePostPayload struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Content string `json:"content"`
	UrlVideo string `json:"url_video"`
	CategoryID string `json:"category_id"`
	UserID string `json:"user_id"`
	ParentID string `json:"parent_id"`
}