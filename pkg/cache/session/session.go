package session

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

type Cache interface {
	Get(token string) (Session, error)
	New(userId uint32) (Session, error)
	Delete(token string) error
	Update(token string) error
}

type cacheImpl struct {
	client redis.UniversalClient
	ttl    time.Duration
}

func NewCache(client redis.UniversalClient, ttl time.Duration) Cache {
	return &cacheImpl{
		client: client,
		ttl:    ttl,
	}
}

func (s *cacheImpl) sessionKey(token string) string {
	return "HORDE_USER_" + token
}

func (s *cacheImpl) Get(token string) (sess Session, err error) {
	if err = s.client.Get(s.sessionKey(token)).Scan(&sess); err != nil {
		if err == redis.Nil {
			err = ErrSessionTokenNotFound
		}
		return
	}
	sess.Token = token
	return
}

func (s *cacheImpl) New(userId uint32) (Session, error) {
	now := time.Now()
	sess := Session{
		UserId:     userId,
		Token:      strings.ReplaceAll(uuid.New().String(), "-", ""),
		CreateTime: now,
		ExpireAt:   now.Add(s.ttl),
	}
	res, err := s.client.SetNX(s.sessionKey(sess.Token), &sess, s.ttl).Result()
	if err != nil {
		return sess, err
	}
	if !res {
		err = ErrDuplicatedSessionToken
	}
	return sess, nil
}

func (s *cacheImpl) Delete(token string) error {
	res, err := s.client.Del(s.sessionKey(token)).Result()
	if err != nil {
		return err
	}
	if res == 0 {
		return ErrSessionTokenNotFound
	}
	return nil
}

func (s *cacheImpl) Update(token string) error {
	res, err := s.client.Expire(s.sessionKey(token), s.ttl).Result()
	if err != nil {
		return err
	}
	if !res {
		return ErrSessionTokenNotFound
	}
	return nil
}

type Session struct {
	UserId     uint32
	Token      string
	CreateTime time.Time
	ExpireAt   time.Time
}

func (s *Session) MarshalBinary() ([]byte, error) {
	return jsoniter.Marshal(s)
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return jsoniter.Unmarshal(data, s)
}
