1) How to add own command and parameter to template:
```go
(v0.1.21)
main.go:
    var -> serviceName      = goopt.String([]string{"--serviceName"}, "", "Name for service in proto")
    func execTemplate -> data["serviceName"] = *serviceName
    func initialize -> conf.ServiceName = *serviceName
/dbmeta/codegen.go:
    type Config struct ->   ServiceName           string
	func WriteTemplate -> data["serviceName"] = c.ServiceName
```
2) add template functions(v0.1.22): https://github.com/Masterminds/sprig 
delAll, updateAll, AddAll
3)  Debug args:
   --sqltype=postgres \
   --connstr {connstr} \
   --exclude=schema_migrations \
   --out ./tmp \
   --overwrite \
   --json \
   --dao="model" \
   --gorm \
   --generate-dao \
   --templateDir=./tmp_templates \
   --protobuf \
   --module github.com/centarium/gen \
   --mapping=./tmp_templates/mapping.json \
   --proto_go_package="tmp" \
   --json-fmt="camel" \
   --serviceName="Test" \
   --proto-args="-I=./tmp/include/ --proto_path=. --go_out=." \
   --database code-generation