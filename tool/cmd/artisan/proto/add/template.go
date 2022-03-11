package add

const protoTemplate = `
{{$srvPath := .Service | toLower}}
syntax = "proto3";

package {{.Package}};

import "google/api/annotations.proto";

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";

service {{.Service}} {
	rpc Create{{.Service}} (Create{{.Service}}Request) returns (Create{{.Service}}Response) {
		option (google.api.http) = {
			post : "/{{$srvPath}}"
			body : "*"
		};
	};
	rpc Update{{.Service}} (Update{{.Service}}Request) returns (Update{{.Service}}Response) {
		option (google.api.http) = {
			put : "/{{$srvPath}}/{uid}"
			body : "*"
		};
	};
	rpc Delete{{.Service}} (Delete{{.Service}}Request) returns (Delete{{.Service}}Response) {
		option (google.api.http) = {
			delete : "/{{$srvPath}}/{uid}"
		};
	};
	rpc Get{{.Service}} (Get{{.Service}}Request) returns (Get{{.Service}}Response) {
		option (google.api.http) = {
			get : "/{{$srvPath}}/{uid}"
		};
	};
	rpc List{{.Service}} (List{{.Service}}Request) returns (List{{.Service}}Response) {
		option (google.api.http) = {
			get : "/{{$srvPath}}"
		};
	};
}

message {{.Service}}Object {
	string uid = 1;
}

message Create{{.Service}}Request { {{.Service}}Object obj = 1; }
message Create{{.Service}}Response {}

message Update{{.Service}}Request { {{.Service}}Object obj = 1;  string uid = 2; }
message Update{{.Service}}Response {}

message Delete{{.Service}}Request { string uid = 1; }
message Delete{{.Service}}Response { {{.Service}}Object obj = 1; }

message Get{{.Service}}Request { string uid = 1;}
message Get{{.Service}}Response { {{.Service}}Object obj = 1; }

message List{{.Service}}Request {}
message List{{.Service}}Response { repeated {{.Service}}Object objList = 1; }
`

const errTemplate = `
syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";

// @plugins=protoc-gen-go-errors
// 错误
enum Error {
  // @msg=未知类型
  // @code=UNKNOWN
  ERR_UNKNOWN = 0;

  // @msg=成功
  // @code=OK
  ERR_OK_STATUS = 1;

  // @msg=未找到资源
  // @code=NOT_FOUND
  ERR_NOT_FOUND = 2;

  // @msg=请求参数无效
  // @code=INVALID_ARGUMENT
  ERR_INVALID_ARGUMENT = 3;

  // @msg=请求后端存储错误
  // @code=INTERNAL
  ERR_INTERNAL_STORE = 4;

  // @msg=内部错误
  // @code=INTERNAL
  ERR_INTERNAL_ERROR = 5;
}
`
