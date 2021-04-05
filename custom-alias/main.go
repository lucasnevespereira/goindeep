package main

import "fmt"

// Type Alias
type ErrorCode int

// Iota increments each tim it is called
const (
	ERR_CODE_OK ErrorCode = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_LOCKED
	ERR_CODE_GENERIC
)


func (ec ErrorCode) IsCritical() bool {
	return ec == ERR_CODE_LOCKED || ec == ERR_CODE_NOT_FOUND
}

func (ec ErrorCode) String() string {
	// [...] sets  the len of the array to the items assigned to the anonymous struct ("ok", "not found", etc..)
	return [...]string{
		"ok",
		"not found",
		"locked",
		"generic",
	}[ec]
	// if this is confusing we can also do a switch case
	//switch ec {
	//case ERR_CODE_NOT_FOUND:
	//	return "not found"
	// etc..
	//}
}



func printErrCode(c ErrorCode) {
	fmt.Printf("code=%v, critical=%v, detail=%v \n", c, c.IsCritical(), c)
}


func main() {
code := ERR_CODE_NOT_FOUND
printErrCode(code)
}
