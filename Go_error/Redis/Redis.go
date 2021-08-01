//Package Redis is for redis
package Redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func OpenRedis() (*redis.Conn, error) {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return nil, errors.Wrap(err, "redis open err")
	}
	return &c, nil
}
