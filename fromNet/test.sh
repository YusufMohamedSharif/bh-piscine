go run . "hello"  | cat -e
go run . "hello world"  | cat -e
go run . "nice 2 meet you"  | cat -e
go run . "you & me"  | cat -e
go run . "123"  | cat -e
go run . "/(")"  | cat -e
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"  | cat -e
go run . ""#$%&/()*+,-./"  | cat -e
go run . "It's Working"  | cat -e