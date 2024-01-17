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
4) add other google.protobuf.BoolValue for serialization consistency
5) add update mapping(generics?)
6) add http api gateway server with json transport
7) add integration tests
8) add postman grpc/http schemas
9)  versioning


1. Why are wrappedFields used everywhere instead of standard types:
   1.1 AddRequest/UpdateRequest - It is impossible to clearly judge which json was passed for
   zero values when mapping Protobuf into Go structures. For example, for Protobuf - structures
message Test {
   bool is_checked = 1;
   int64 count = 2;
   tmp_name = 3; 
}
Json object {} и Json object { "is_checked" : false, "count" : 0, "tmp_name" : "" }
Matches the same in Go - structure
type Test struct {
   IsChecked bool
   Count int64
   TmpName string
}
   Which leads to the following critical problems in the database:
1) There is a field is_checked of type bool in the database - we cannot judge unambiguously, switched
   whether the user has some kind of checkbox or the field was simply not passed.
2) There is a field count of type bool - for example, this is the balance in the warehouse. Again, we cannot judge
   whether the user has established that the warehouse is out of stock or the field was simply not transmitted
3) In the database, the tmp_name field is of type string - the user can “zero” the field, or it can be
   it is simply not conveyed - it is not possible to distinguish one situation from another.
   Possible solutions: 
  А) Just make every single required field and run it every time
   the whole table back and forth. If someone somewhere does this, then you can immediately see the person whom
   needs to be fired.
  Б) GRPC - client. What if you suddenly need to provide an API for a client that grpc does not support?
   or categorically refuses to use it? The option is also absolutely not suitable.
   1.2 GetRequest/Model - imagine that we have a table with 200 fields. I have encountered cases when, when changing
   of one field at the front, all fields were updated at once. I think there is no need to explain that updating one field
   is not the same as updating 200 fields, even if their values have not changed. Unfortunately, very many
   people (especially in php) do this. Not all of them were fired on time.
I.e. mapping is needed - get all explicitly passed fields from the request (it is assumed that the front knows that they
   have changed and does not send the same thing back and forth) and generate an update line.
   Next - how to make it easier - load AddTest, UpdateTestByID, UpdateTestByCode...N into the model and call
   mapping function once or call the function for each request N times?
   GetRequest - since everything is already in wrapped fields, then GetRequest is easier to do like everything else

2. What is integer for - fields with (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {type: INTEGER}] -
in swagger - the proto generator there is a known error when an int64 field is signed with the "string" type. This
tag fixes this bug.