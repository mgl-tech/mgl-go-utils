package mgl

type idUtils struct {
}

var IdUtils = &idUtils{}

func (receiver idUtils) getId() (*Snowflake, error) {
	return NewSnowflake(int64(0), int64(0))
}
