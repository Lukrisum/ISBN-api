package dao

type BookInfo struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Subname       string `json:"subname"`
	Author        string `json:"author"`
	Translator    string `json:"translator"`
	Publishing    string `json:"publishing"`
	Published     string `json:"published"`
	Designed      string `json:"designed"`
	Code          string `json:"code"`
	Douban        int64  `json:"douban"`
	DoubanScore   int    `json:"doubanScore"`
	Brand         string `json:"brand"`
	Weight        string `json:"weight"`
	Size          string `json:"size"`
	Pages         string `json:"pages"`
	PhotoUrl      string `json:"photoUrl"`
	LocalPhotoUrl string `json:"localPhotoUrl"`
	Price         string `json:"price"`
	Froms         string `json:"froms"`
	Num           int32  `json:"num"`
	CreateTime    string `json:"createTime"`
	Uptime        string `json:"uptime"`
	AuthorIntro   string `json:"authorIntro"`
	Description   string `json:"description"`
}
