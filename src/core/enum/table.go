package enum

type TableModel struct {
	TableName  string `json:"tableName"`
	PrimaryKey string `json:"primaryKey"`
	Permission string `json:"permission"`
}

// TableEnum 构造函数
func TableEnum(tableName string, primaryKey string) TableModel {
	return TableModel{
		TableName:  tableName,
		PrimaryKey: primaryKey,
	}
}
