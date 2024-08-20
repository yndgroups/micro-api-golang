package mapping

type userInfo struct {
	UserId   int16  `json:"userId"`
	UserName string `json:"userName"`
	Age      string `json:"age"`
}

var UserInfo = &userInfo{}

func (u *userInfo) GetIndexName() string {
	return "user_info2"
}

// mapping其实就是一种描述索引库字段类型的一种形式
// es 存储的数据的形态都是json。包括mapping字段描述也是json方式
// "analyzer":"ik_max_word"  如果这样写，es必须安装ik分词器，否则会报。此含义是未来如果搜索数据索引化的时候，它会把你的标题进行分词和你的文档进行关联
const UserMapping = `
{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            }
        }
    }
}`
