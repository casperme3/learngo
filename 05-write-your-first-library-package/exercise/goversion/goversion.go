package goversion

import "runtime"

//GetVersion is wrapper to runtime Version function
func GetVersion() string {
	return "You are running GO version [" + runtime.Version() + "]"
}
