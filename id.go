package mgl

import "github.com/mgl-tech/mgl-go-utils/id"

type idUtils struct {
}

var IdUtils = &idUtils{}

func (receiver idUtils) getId() (*id.Snowflake, error) {
	return id.NewSnowflake(int64(0), int64(0))
}
