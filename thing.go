package thingiverse

type ThingRes struct {
	Hits  []Thing `json:"hits"`
	Total int     `json:"total"`
}

type Thing struct {
	CommentCount int         `json:"comment_count"`
	CreatedAt    string      `json:"created_at"`
	Creator      Creator     `json:"creator"`
	ID           int         `json:"id"`
	IsNsfw       bool        `json:"is_nsfw"`
	IsPrivate    interface{} `json:"is_private"`
	IsPublished  interface{} `json:"is_published"`
	IsPurchased  interface{} `json:"is_purchased"`
	LikeCount    int         `json:"like_count"`
	Name         string      `json:"name"`
	PreviewImage string      `json:"preview_image"`
	PublicURL    string      `json:"public_url"`
	Tags         []Tag       `json:"tags"`
	Thumbnail    string      `json:"thumbnail"`
	URL          string      `json:"url"`
}

type Creator struct {
	AcceptsTips      bool   `json:"accepts_tips"`
	CountOfDesigns   int    `json:"count_of_designs"`
	CountOfFollowers int    `json:"count_of_followers"`
	CountOfFollowing int    `json:"count_of_following"`
	Cover            string `json:"cover"`
	FirstName        string `json:"first_name"`
	ID               int    `json:"id"`
	IsFollowing      bool   `json:"is_following"`
	LastName         string `json:"last_name"`
	Location         string `json:"location"`
	Name             string `json:"name"`
	PublicURL        string `json:"public_url"`
	Thumbnail        string `json:"thumbnail"`
	URL              string `json:"url"`
}

type Tag struct {
	AbsoluteURL string `json:"absolute_url"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	ThingsURL   string `json:"things_url"`
	URL         string `json:"url"`
}
