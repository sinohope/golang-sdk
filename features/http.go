package features

import "github.com/sinohope/sinohope-golang-sdk/common"

type HTTP interface {
	Post(path string, request interface{}) (*common.Response, error)
}
