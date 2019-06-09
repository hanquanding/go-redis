package gredis

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisCli redis.Conn
	Addr     string
)

func init() {
	var err error
	Addr = "127.0.0.1:6379"
	RedisCli, err = redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
}

func getConn() (redis.Conn, error) {
	cli, err := redis.Dial("tcp", Addr)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func Get(key string) (interface{}, error) {
	cli, err := getConn()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	i, e := cli.Do("GET", key)
	return i, e
}

func GetBytes(key string) ([]byte, error) {
	return redis.Bytes(Get(key))
}

func GetString(key string) (string, error) {
	return redis.String(Get(key))
}

func GetInt(key string) (int, error) {
	return redis.Int(Get(key))
}

func GetInt64(key string) (int64, error) {
	return redis.Int64(Get(key))
}

func GetJson(key string, reply interface{}) error {
	jsonBytes, err := GetBytes(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, &reply)
}

func Set(key string, value interface{}) error {
	cli, err := getConn()
	if err != nil {
		return err
	}
	defer cli.Close()
	_, err = cli.Do("SET", key, value)
	return err
}

func SetNX(key string, value interface{}) (bool, error) {
	cli, err := getConn()
	if err != nil {
		return false, err
	}
	defer cli.Close()
	reply, err := cli.Do("SETNX", key, value)
	if err != nil {
		return false, err
	}
	return reply.(int64) > 0, nil
}

func SetJson(key string, value interface{}) error {
	cli, err := getConn()
	if err != nil {
		return err
	}
	defer cli.Close()
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = cli.Do("SET", key, string(jsonBytes))
	return err
}

func Exists(key string) (bool, error) {
	cli, err := getConn()
	if err != nil {
		return false, err
	}
	defer cli.Close()
	return redis.Bool(cli.Do("EXISTS", key))
}

func Remove(key string) (bool, error) {
	cli, err := getConn()
	if err != nil {
		return false, err
	}
	defer cli.Close()
	return redis.Bool(cli.Do("DEL", key))
}
