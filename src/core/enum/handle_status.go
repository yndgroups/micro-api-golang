package enum

type Status int

// iota 初始化后会自动递增
const (
	Insert Status = 1
	Delete Status = 2
	Update Status = 3
	Query  Status = 4
)
