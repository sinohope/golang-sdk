package features

import "github.com/sinohope/sinohope-golang-sdk/common"

type Gateway interface {
	Post(path string, request interface{}) (*common.Response, error)
}
