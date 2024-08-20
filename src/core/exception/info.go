package exception

import (
	"fmt"
	"regexp"
	"strings"
)

// 数据库错误异常信息
func DbMsg(errorMsg string, serverName string) string {
	sArr := strings.Split(errorMsg, " ")
	for _, s := range sArr {
		reg, _ := regexp.Compile(`^\d{3,6}$`)
		errCode := reg.FindString(s)
		if errCode != "" {
			msg := GetMysqlErrMesaage(errCode)
			if msg != "" {
				return msg
			}
			return fmt.Sprintf("[%s]SQL异常:%s, errCode:%s", serverName, GetMysqlErrMesaage(errCode), errCode)
		}
		println(reg.FindString(s))
	}
	if strings.Contains(errorMsg, "Error 1366") {
		return fmt.Sprintf("[%s]SQL异常, errCode:%s", serverName, errorMsg)
	}
	return errorMsg
}

// 数据库异常状态码对应的信息
func GetMysqlErrMesaage(errCode string) string {
	var maps = make(map[string]string)
	maps["130"] = "文件格式不正确。"
	maps["145"] = "文件无法打开"
	//  ER_NO; SQLSTATE： HY000
	maps["1002"] = "用于构建其他消息"
	// ER_YES; SQLSTATE： HY000
	// 扩展EXPLAIN格式生成注释消息。在后续输出中的这些消息ER_YES的Code列中使用SHOW WARNINGS。
	maps["1003"] = "扩展EXPLAIN格式生成注释消息"
	// ER_CANT_CREATE_FILE; SQLSTATE： HY000
	// 由于无法创建或复制某些操作所需的文件而发生。可能的原因：源文件的权限问题；目标文件已存在，但不可写。
	maps["1004"] = "源文件的权限问题；目标文件已存在，但不可写"
	//  ER_CANT_CREATE_TABLE; SQLSTATE： HY000
	// 创建表失败 InnoDB无法创建表时报告此错误。如果错误消息涉及错误150，则表创建失败，因为 未正确形成外键约束。如果错误消息指示错误-1，则表创建可能失败，因为该表包含与内部InnoDB表名称匹配的列名称。
	maps["1005"] = "无法创建表"
	// ER_CANT_CREATE_DB; SQLSTATE： HY000
	// 创建数据库失败
	maps["1006"] = "无法创建数据库"
	// ER_DB_CREATE_EXISTS; SQLSTATE： HY000
	// 尝试创建数据库失败，因为该数据库已经存在。如果您确实要替换现有数据库，请先删除数据库，或者如果要保留现有数据库而不使该语句产生错误，请在该语句中添加一个IF NOT EXISTS子句 CREATE DATABASE。
	maps["1007"] = "数据库已存在，创建数据库失败"
	// ER_DB_DROP_EXISTS; SQLSTATE： HY000
	// 无法删除数据库'％s'; 数据库不存在
	maps["1008"] = "数据库不存在，删除数据库失败"
	maps["1009"] = "不能删除数据库文件导致删除数据库失败"
	//ER_DB_DROP_RMDIR; SQLSTATE： HY000
	// 删除数据库时出错（无法rmdir'％s'，错误号：％d-％s）
	maps["1010"] = "不能删除数据目录导致删除数据库失败"
	maps["1011"] = "删除数据库文件失败"
	// ER_CANT_FIND_SYSTEM_REC; SQLSTATE： HY000
	// InnoDB如果尝试访问 InnoDB INFORMATION_SCHEMA 表InnoDB不可用，则 返回。
	maps["1012"] = "不能读取系统表中的记录"
	// 无法获取“％s”的状态 ER_CANT_GET_STAT; SQLSTATE： HY000
	maps["1013"] = "无法获取状态"
	// ER_CANT_LOCK; SQLSTATE： HY000
	maps["1015"] = "无法锁定文件"
	// ER_CANT_OPEN_FILE; SQLSTATE： HY000
	// InnoDB当找不到InnoDB 数据文件中的表时，报告此错误。
	maps["1016"] = "无法打开文件"
	// ER_FILE_NOT_FOUND; SQLSTATE： HY000
	maps["1017"] = "找不到文件"
	// ER_CANT_READ_DIR; SQLSTATE： HY000
	// 无法读取“％s”的目录
	maps["1018"] = "无法读取目录"
	// ER_CHECKREAD; SQLSTATE： HY000
	// 自上次读取表'％s'以来，记录已更改
	maps["1020"] = "记录已被其他用户修改"
	maps["1021"] = "硬盘剩余空间不足，请加大硬盘可用空间"
	// 不会写；表'％s'中的重复键
	maps["1022"] = "关键字重复，更改记录失败"
	// ER_DUP_KEY; SQLSTATE： 23000
	maps["1023"] = "关闭时发生错误"
	// ER_ERROR_ON_READ; SQLSTATE： HY000
	// 读取文件'％s'时出错（错误号：％d-％s）
	maps["1024"] = "读文件时错误"
	// ER_ERROR_ON_RENAME; SQLSTATE： HY000
	// 将“％s”重命名为“％s”时出错（错误号：％d-％s）
	maps["1025"] = "重命名时发生错误"
	// 写入文件'％s'时出错（错误号：％d-％s）ER_ERROR_ON_WRITE; SQLSTATE： HY000
	maps["1026"] = "写文件时错误"
	// ER_FILE_USED; SQLSTATE： HY000
	maps["1027"] = "被锁定，无法更改"
	// 排序已中止 ER_FILSORT_ABORT 在8.0.18之后被删除。 ER_FILSORT_ABORT; SQLSTATE： HY000
	maps["1028"] = "排序已中止,在8.0.18之后被删除"
	// 检查该%d值以查看操作系统错误的含义。例如，28表示您已用完磁盘空间。ER_GET_ERRNO; SQLSTATE： HY000
	maps["1030"] = "来自存储引擎的错误"
	// “％s”的表存储引擎没有此选项 ER_ILLEGAL_HA; SQLSTATE： HY000
	maps["1031"] = "表存储引擎没有此选项"
	// ER_KEY_NOT_FOUND; SQLSTATE： HY000
	// 在“％s”中找不到记录
	maps["1032"] = "记录不存在"
	// ER_NOT_FORM_FILE; SQLSTATE： HY000
	// 文件“％s”中的信息不正确
	maps["1033"] = "文件中的信息不正确"
	// 表'％s'的密钥文件不正确；尝试修复它 ER_NOT_KEYFILE; SQLSTATE： HY000
	maps["1034"] = "表的密钥文件不正确；尝试修复它"
	// 表'％s'的旧密钥文件；修复它！ ER_OLD_KEYFILE; SQLSTATE： HY000
	maps["1035"] = "表的旧密钥文件；修复它"
	// 表'％s'是只读的 ER_OPEN_AS_READONLY; SQLSTATE： HY000
	maps["1036"] = "数据表是只读的，不能对它进行修改"
	// 内存不足；重新启动服务器，然后重试（需要％d字节） ER_OUTOFMEMORY; SQLSTATE： HY001
	maps["1037"] = "系统内存不足，请重启数据库或重启服务器"
	// ER_OUT_OF_SORTMEMORY; SQLSTATE： HY001
	maps["1038"] = "用于排序的内存不足，请增大排序缓冲区"
	// ER_CON_COUNT_ERROR; SQLSTATE： 08004
	// 内存不足；检查mysqld或其他进程是否使用了所有可用内存；如果不是，则可能必须使用“ ulimit”来允许mysqld使用更多的内存，或者可以添加更多的交换空间
	maps["1040"] = "已到达数据库的最大连接数，请加大数据库可用连接数"
	// ER_OUT_OF_RESOURCES; SQLSTATE： HY000
	maps["1041"] = "系统内存不足"
	// ER_BAD_HOST_ERROR; SQLSTATE： 08S01
	// 无法获取您的地址的主机名
	maps["1042"] = "无效的主机名"
	// ER_HANDSHAKE_ERROR; SQLSTATE： 08S01
	maps["1043"] = "无效连接"
	// ER_DBACCESS_DENIED_ERROR; SQLSTATE：42000
	maps["1044"] = "当前用户没有访问数据库的权限"
	// ER_ACCESS_DENIED_ERROR; SQLSTATE： 28000
	maps["1045"] = "不能连接数据库，用户名或密码错误"
	// ER_NO_DB_ERROR; SQLSTATE： 3D000
	maps["1046"] = "未选择数据库"
	// ER_UNKNOWN_COM_ERROR; SQLSTATE： 08S01
	maps["1047"] = "未知命令"
	// ER_BAD_NULL_ERROR; SQLSTATE： 23000
	// 列'％s'不能为空
	maps["1048"] = "字段不能为空"
	// ER_BAD_DB_ERROR; SQLSTATE： 42000
	// 消息：未知数据库'％s'
	maps["1049"] = "数据库不存在"
	// ER_TABLE_EXISTS_ERROR; SQLSTATE： 42S01
	// 表'％s'已经存在
	maps["1050"] = "数据表已存在"
	// ER_BAD_TABLE_ERROR; SQLSTATE： 42S02
	// 未知表'％s'
	maps["1051"] = "数据表不存在"
	// ER_NON_UNIQ_ERROR; SQLSTATE： 23000
	// 查询中没有适当限定的列出现在选择列表或ON子句中
	maps["1052"] = "列不明确"
	// ER_SERVER_SHUTDOWN; SQLSTATE： 08S01
	maps["1053"] = "服务器正在关闭"
	// ER_BAD_FIELD_ERROR; SQLSTATE： 42S22
	maps["1054"] = "字段不存在"
	// ER_WRONG_FIELD_WITH_GROUP; SQLSTATE：42000
	maps["1055"] = "不在GROUP BY中"
	// 无法在“％s”上分组
	maps["1056"] = "无法分组"
	// ER_WRONG_SUM_SELECT; SQLSTATE： 42000
	maps["1057"] = "语句在同一语句中具有求和函数和列"
	// 列计数与值计数不匹配
	maps["1058"] = "列计数与值计数不匹配"
	// ER_TOO_LONG_IDENT; SQLSTATE： 42000
	maps["1059"] = "标识符名称太长"
	// ER_DUP_FIELDNAME; SQLSTATE： 42S21
	maps["1060"] = "重复的列名"
	//  ER_DUP_KEYNAME; SQLSTATE： 42000
	maps["1061"] = "重复的键名"
	// ER_DUP_ENTRY; SQLSTATE： 23000
	// 返回此错误的消息使用的格式字符串 ER_DUP_ENTRY_WITH_KEY_NAME。
	maps["1062"] = "键重复"
	// ER_WRONG_FIELD_SPEC; SQLSTATE： 42000
	maps["1063"] = "列的列说明符不正确"
	// ％s在第％d行靠近“％s”
	maps["1064"] = "SQL错误"
	// ER_EMPTY_QUERY; SQLSTATE： 42000 SQL语句为空
	maps["1065"] = "无效的SQL语句, SQL语句为空"
	// ER_NONUNIQ_TABLE; SQLSTATE： 42000
	maps["1066"] = "不是唯一的表/别名"
	// ER_INVALID_DEFAULT; SQLSTATE： 42000
	maps["1067"] = "默认值无效"
	// ER_MULTIPLE_PRI_KEY; SQLSTATE： 42000
	maps["1068"] = "定义了多个主键"
	// ER_TOO_MANY_KEYS; SQLSTATE： 42000
	maps["1069"] = "指定的密钥太多,允许的最大多少个键"
	// ER_TOO_MANY_KEY_PARTS; SQLSTATE： 42000
	maps["1070"] = "指定的关键部分太多；最多允许多少个键"
	// ER_TOO_LONG_KEY; SQLSTATE： 42000
	maps["1071"] = "指定密钥太长；最大密钥长度为％d个字节"
	// ER_KEY_COLUMN_DOES_NOT_EXITS; SQLSTATE：42000
	maps["1072"] = "表中不存在键列"
	// ER_BLOB_USED_AS_KEY; SQLSTATE： 42000
	maps["1073"] = "BLOB列不能用于已使用表类型的键规范中"
	// ER_TOO_BIG_FIELDLENGTH; SQLSTATE： 42000
	maps["1074"] = "列长度对于列'％s'太大（最大值=％lu）；使用BLOB或TEXT代替"
	// ER_WRONG_AUTO_KEY; SQLSTATE： 42000
	maps["1075"] = "错误的表定义；只能有一个自动列，并且必须将其定义为键"
	// ER_READY; SQLSTATE： HY000
	maps["1076"] = "已准备好进行连接"
	// ER_NORMAL_SHUTDOWN; SQLSTATE： HY000
	maps["1077"] = "正常关闭,在8.0.4之后被删除"
	// ER_SHUTDOWN_COMPLETE; SQLSTATE： HY000
	maps["1079"] = "关机完成"
	// ER_FORCING_CLOSE; SQLSTATE： 08S01
	maps["1080"] = "强制关闭线程"
	maps["1081"] = "不能建立Socket连接"
	// ER_NO_SUCH_INDEX; SQLSTATE： 42S12
	maps["1082"] = "表没有类似于CREATE INDEX中使用的索引,重新创建表"
	// ER_WRONG_FIELD_TERMINATORS; SQLSTATE：42000
	maps["1083"] = "字段分隔符参数不是预期的；查看手册"
	// ER_BLOBS_AND_NO_TERMINATED; SQLSTATE：42000
	maps["1084"] = "您不能对BLOB使用固定的行长,请使用“字段终止于”"
	// ER_TEXTFILE_NOT_READABLE; SQLSTATE：HY000
	maps["1085"] = "文件'％s'必须在数据库目录中，或者所有人都可以读取"
	// ER_FILE_EXISTS_ERROR; SQLSTATE： HY000
	maps["1086"] = "文件已存在"
	// ER_LOAD_INFO; SQLSTATE： HY000
	maps["1087"] = "记录:已删除，跳过，警告"
	// ER_ALTER_INFO; SQLSTATE： HY000
	maps["1088"] = "记录重复项"
	// ER_WRONG_SUB_KEY; SQLSTATE： HY000
	maps["1089"] = "消息：不正确的前缀密钥；使用的密钥部分不是字符串，使用的长度比密钥部分长，或者存储引擎不支持唯一的前缀密钥"
	// ER_CANT_REMOVE_ALL_FIELDS; SQLSTATE：42000
	maps["1090"] = "您不能使用ALTER TABLE删除所有列；改用DROP TABLE"
	maps["1114"] = "数据表已满，不能容纳任何记录"
	maps["1116"] = "打开的数据表太多"
	maps["1129"] = "数据库出现异常，请重启数据库"
	maps["1130"] = "连接数据库失败，没有连接数据库的权限"
	maps["1133"] = "数据库用户不存在"
	maps["1141"] = "当前用户无权访问数据库"
	maps["1142"] = "当前用户无权访问数据表"
	maps["1143"] = "当前用户无权访问数据表中的字段"
	maps["1146"] = "数据表不存在"
	maps["1147"] = "未定义用户对数据表的访问权限"
	maps["1149"] = "SQL语句语法错误"
	maps["1158"] = "网络错误，出现读错误，请检查网络连接状况"
	maps["1159"] = "网络错误，读超时，请检查网络连接状况"
	maps["1160"] = "网络错误，出现写错误，请检查网络连接状况"
	maps["1161"] = "网络错误，写超时，请检查网络连接状况"
	maps["1062"] = "字段值重复，入库失败"
	maps["1169"] = "字段值重复，更新记录失败"
	maps["1177"] = "打开数据表失败"
	maps["1180"] = "提交事务失败"
	maps["1181"] = "回滚事务失败"
	maps["1203"] = "当前用户和数据库建立的连接已到达数据库的最大连接数，请增大可用的数据库连接数或重启数据库"
	maps["1205"] = "加锁超时"
	maps["1211"] = "当前用户没有创建用户的权限"
	maps["1216"] = "外键约束检查失败，更新子表记录失败"
	maps["1217"] = "外键约束检查失败，删除或修改主表记录失败"
	maps["1226"] = "当前用户使用的资源已超过所允许的资源，请重启数据库或重启服务器"
	maps["1227"] = "权限不足，您无权进行此操作"
	maps["1235"] = "MySQL版本过低，不具有本功能"
	maps["1250"] = "客户端不支持服务器要求的认证协议，请考虑升级客户端。"
	maps["1251"] = "Client 不能支持 authentication protocol 的要求Client does not support authentication protocol requested by server; consider upgrading MySQL clientQuote:"
	maps["1264"] = "取值超出限定范围"
	maps["1267"] = "不合法的混合字符集。"
	maps["1364"] = "必传字段没有传值"
	maps["1366"] = "参数传递不合法"
	maps["2002"] = "服务器端口不对。"
	maps["2003"] = "MySQL 服务没有启动，请启动该服务。"
	maps["2008"] = "MySQL client ran out of memory错误指向了MySQL客户mysql。这个错误的原因很简单，客户没有足够的内存存储全部结果。"
	maps["2013"] = "远程连接数据库是有时会有这个问题，MySQL 服务器在执行一条 SQL 语句的时候失去了连接造成的。"
	maps["10048"] = `建议在my.ini文件中修改最大连接数， 把 mysql_connect() 方法都改成了 mysql_pconnect() 方法. 要修改mysql_pconnect()，可以在论坛的data目录的sql_config.php中 p c o n n e c t = 0 ; / / 是 否 持 久 连 接 修 改 成 pconnect = 0; //是否持久连接 修改成pconnect=0;//是否持久连接修改成pconnect = 1; 开启防刷新,严禁刷新太快.`
	maps["10055"] = `没有缓存空间可利用，查看下你的C盘空间是否已经满,清除一些没有用的文件. 可以在后台的"论坛核心设置","核心功能设置"里"进程优化"开启,"GZIP 压缩输出"关闭.查找了一下10055（没有缓存空间可利用）`
	maps["10061"] = `启动这台机器上的MySQL服务 如服务启动失败，一定是你的my.ini文件出了差错， MySQL服务不能正常启动 你删除了它后，MySQL就会按其默认配置运行， 那就没有问题了`
	return maps[errCode]
}
