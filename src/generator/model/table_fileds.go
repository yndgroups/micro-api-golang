package model

type TableFieldList struct {
	FieldName string `json:"field"`
	FieldType string `json:"type"`
}

type TableField struct {
	Field      string `json:"field"`
	Type       string `json:"type"`
	Collation  string `json:"collation"`
	Null       string `json:"null"`
	Key        string `json:"key"`
	Default    string `json:"default"`
	Extra      string `json:"extra"`
	Privileges string `json:"frivileges"`
	Comment    string `json:"comment"`
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}
