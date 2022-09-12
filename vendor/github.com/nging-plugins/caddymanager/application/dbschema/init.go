// @generated Do not edit a file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

var WithPrefix = func(tableName string) string {
	return "" + tableName
}

var DBI = factory.DefaultDBI

func init() {

	DBI.FieldsRegister(map[string]map[string]*factory.FieldInfo{"nging_access_log": {"body_bytes": {Name: "body_bytes", DataType: "bigint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "响应body字节数", GoType: "uint64", MyType: "", GoName: "BodyBytes"}, "brower_name": {Name: "brower_name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 60, Options: []string{}, DefaultValue: "", Comment: "浏览器名", GoType: "string", MyType: "", GoName: "BrowerName"}, "brower_type": {Name: "brower_type", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 20, Options: []string{}, DefaultValue: "", Comment: "浏览器类型(spider/mobile/pc)", GoType: "string", MyType: "", GoName: "BrowerType"}, "created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "elapsed": {Name: "elapsed", DataType: "decimal", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 1e+08, Precision: 2, MaxSize: 10, Options: []string{}, DefaultValue: "0.00", Comment: "耗时(秒)", GoType: "float64", MyType: "", GoName: "Elapsed"}, "hit_status": {Name: "hit_status", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "缓存服务器命中状态", GoType: "uint", MyType: "", GoName: "HitStatus"}, "host": {Name: "host", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 60, Options: []string{}, DefaultValue: "", Comment: "Header中的Host, 一般会是域名", GoType: "string", MyType: "", GoName: "Host"}, "id": {Name: "id", DataType: "bigint", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint64", MyType: "", GoName: "Id"}, "local_addr": {Name: "local_addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 120, Options: []string{}, DefaultValue: "", Comment: "本机地址", GoType: "string", MyType: "", GoName: "LocalAddr"}, "method": {Name: "method", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "", Comment: "GET POST HEAD OPTIONS PUT DELETE...", GoType: "string", MyType: "", GoName: "Method"}, "minute": {Name: "minute", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 5, Options: []string{}, DefaultValue: "", Comment: "H:i", GoType: "string", MyType: "", GoName: "Minute"}, "referer": {Name: "referer", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "来源网址", GoType: "string", MyType: "", GoName: "Referer"}, "remote_addr": {Name: "remote_addr", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 30, Options: []string{}, DefaultValue: "", Comment: "IP地址", GoType: "string", MyType: "", GoName: "RemoteAddr"}, "scheme": {Name: "scheme", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "", Comment: "https/http", GoType: "string", MyType: "", GoName: "Scheme"}, "status_code": {Name: "status_code", DataType: "mediumint", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "状态码", GoType: "uint", MyType: "", GoName: "StatusCode"}, "time_local": {Name: "time_local", DataType: "datetime", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "1970-01-01 00:00:00", Comment: "本地时间格式", GoType: "string", MyType: "", GoName: "TimeLocal"}, "uri": {Name: "uri", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "网址", GoType: "string", MyType: "", GoName: "Uri"}, "user": {Name: "user", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 64, Options: []string{}, DefaultValue: "", Comment: "用户", GoType: "string", MyType: "", GoName: "User"}, "user_agent": {Name: "user_agent", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "用户代理", GoType: "string", MyType: "", GoName: "UserAgent"}, "version": {Name: "version", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 10, Options: []string{}, DefaultValue: "", Comment: "HTTP/1.0, HTTP/1.1 ...", GoType: "string", MyType: "", GoName: "Version"}, "vhost_id": {Name: "vhost_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "虚拟主机ID", GoType: "uint", MyType: "", GoName: "VhostId"}, "x_forward_for": {Name: "x_forward_for", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 120, Options: []string{}, DefaultValue: "", Comment: "IP", GoType: "string", MyType: "", GoName: "XForwardFor"}, "x_real_ip": {Name: "x_real_ip", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 30, Options: []string{}, DefaultValue: "", Comment: "IP", GoType: "string", MyType: "", GoName: "XRealIp"}}, "nging_vhost": {"created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "disabled": {Name: "disabled", DataType: "enum", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{"Y", "N"}, DefaultValue: "N", Comment: "是否停用", GoType: "string", MyType: "", GoName: "Disabled"}, "domain": {Name: "domain", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "域名", GoType: "string", MyType: "", GoName: "Domain"}, "group_id": {Name: "group_id", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "组", GoType: "uint", MyType: "", GoName: "GroupId"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "ID", GoType: "uint", MyType: "", GoName: "Id"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 200, Options: []string{}, DefaultValue: "", Comment: "网站名称", GoType: "string", MyType: "", GoName: "Name"}, "root": {Name: "root", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 500, Options: []string{}, DefaultValue: "", Comment: "网站物理路径", GoType: "string", MyType: "", GoName: "Root"}, "setting": {Name: "setting", DataType: "text", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "设置", GoType: "string", MyType: "", GoName: "Setting"}, "updated": {Name: "updated", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "更新时间", GoType: "uint", MyType: "", GoName: "Updated"}}, "nging_vhost_group": {"created": {Name: "created", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "创建时间", GoType: "uint", MyType: "", GoName: "Created"}, "description": {Name: "description", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 255, Options: []string{}, DefaultValue: "", Comment: "说明", GoType: "string", MyType: "", GoName: "Description"}, "id": {Name: "id", DataType: "int", Unsigned: true, PrimaryKey: true, AutoIncrement: true, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "", Comment: "", GoType: "uint", MyType: "", GoName: "Id"}, "name": {Name: "name", DataType: "varchar", Unsigned: false, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 120, Options: []string{}, DefaultValue: "", Comment: "组名", GoType: "string", MyType: "", GoName: "Name"}, "uid": {Name: "uid", DataType: "int", Unsigned: true, PrimaryKey: false, AutoIncrement: false, Min: 0, Max: 0, Precision: 0, MaxSize: 0, Options: []string{}, DefaultValue: "0", Comment: "用户ID", GoType: "uint", MyType: "", GoName: "Uid"}}})

	DBI.ColumnsRegister(map[string][]string{"nging_access_log": {"id", "vhost_id", "remote_addr", "x_real_ip", "x_forward_for", "local_addr", "elapsed", "host", "user", "time_local", "minute", "method", "uri", "version", "status_code", "body_bytes", "referer", "user_agent", "hit_status", "scheme", "brower_name", "brower_type", "created"}, "nging_vhost": {"id", "name", "group_id", "domain", "root", "created", "updated", "setting", "disabled"}, "nging_vhost_group": {"id", "uid", "name", "description", "created"}})

	DBI.ModelsRegister(factory.ModelInstancers{`NgingAccessLog`: factory.NewMI("nging_access_log", func(connID int) factory.Model { return &NgingAccessLog{base: *((&factory.Base{}).SetConnID(connID))} }, ""), `NgingVhost`: factory.NewMI("nging_vhost", func(connID int) factory.Model { return &NgingVhost{base: *((&factory.Base{}).SetConnID(connID))} }, "虚拟主机"), `NgingVhostGroup`: factory.NewMI("nging_vhost_group", func(connID int) factory.Model { return &NgingVhostGroup{base: *((&factory.Base{}).SetConnID(connID))} }, "虚拟主机组")})

}
