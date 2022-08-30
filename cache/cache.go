package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"time"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
)

var bm cache.Cache

var nilCtx = context.TODO()

func Get(key string, e interface{}) error {

	val, err := bm.Get(nilCtx, key)

	if err != nil {
		return errors.New("get cache error:" + err.Error())
	}

	if val == nil {
		return errors.New("cache does not exist")
	}
	if b, ok := val.([]byte); ok {
		buf := bytes.NewBuffer(b)

		decoder := gob.NewDecoder(buf)

		err := decoder.Decode(e)

		if err != nil {
			logs.Error("反序列化对象失败:", err)
		}
		return err
	} else if s, ok := val.(string); ok && s != "" {

		buf := bytes.NewBufferString(s)

		decoder := gob.NewDecoder(buf)

		err := decoder.Decode(e)

		if err != nil {
			logs.Error("反序列化对象失败:", err)
		}
		return err
	}
	return errors.New("value is not []byte or string")
}

func Set(key string, val interface{}, timeout time.Duration) error {

	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(val)
	if err != nil {
		logs.Error("序列化对象失败:", err)
		return err
	}

	return bm.Put(nilCtx, key, buf.String(), timeout)
}

func Delete(key string) error {
	return bm.Delete(nilCtx, key)
}
func Incr(key string) error {
	return bm.Incr(nilCtx, key)
}
func Decr(key string) error {
	return bm.Decr(nilCtx, key)
}
func IsExist(key string) (bool, error) {
	return bm.IsExist(nilCtx, key)
}
func ClearAll() error {
	return bm.ClearAll(nilCtx)
}

func StartAndGC(config string) error {
	return bm.StartAndGC(config)
}

// initialize cache
func Init(c cache.Cache) {
	bm = c
}
