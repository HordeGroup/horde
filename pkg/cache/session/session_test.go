package session

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func NewRedisClient(masterName string, port int64, password string) (redis.UniversalClient, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		MasterName: masterName,
		Addrs:      []string{":" + strconv.FormatInt(port, 10)},
		Password:   password,
	})
	res, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return client, nil
}

func TestSessionCache(t *testing.T) {
	client, _ := NewRedisClient("", 6379, "")

	t.Run("新建SESSION", func(t *testing.T) {
		ttl := time.Second
		sessCache := NewCache(client, ttl)
		sess, err := sessCache.New(1)
		require.NoError(t, err)
		fmt.Println(sess.Token)
		time.Sleep(2 * time.Second)
		sess, err = sessCache.Get(sess.Token)
		require.NoError(t, err)
		fmt.Println(sess)
	})

}
