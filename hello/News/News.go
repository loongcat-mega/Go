package News

type TencentNews struct {
	Title, Content, Author string
	Pubtime                int
}

var Reps = make(map[int]TencentNews, 100)

func init() {
	//Reps[0] = TencentNews{Title: "a", Author: "author1", content: "news", pubtime: 20}
	//Reps[1] = TencentNews{Title: "b", Author: "author2", content: "news2", pubtime: 30}
}
