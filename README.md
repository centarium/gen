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