package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var addr = "127.0.0.1:6379"

func main() {

	ctx := context.Background()
	cli := redis.NewClient(&redis.Options{Addr: addr})
	defer cli.Close()

	CalcValuesMemory(ctx, cli, 127, 10000)
	CalcValuesMemory(ctx, cli, 512, 10000)
	CalcValuesMemory(ctx, cli, 1024, 10000)
	CalcValuesMemory(ctx, cli, 5120, 10000)

	fmt.Println("--------------------------------------------------------")
	CalcValuesMemory(ctx, cli, 127, 500000)
	CalcValuesMemory(ctx, cli, 512, 500000)
	CalcValuesMemory(ctx, cli, 1024, 500000)
	CalcValuesMemory(ctx, cli, 5120, 500000)
}

func GetCurrentRedisUsedMemory(ctx context.Context, cli *redis.Client) uint64 {
	memoryInfoStr, _ := cli.Info(ctx, "memory").Result()

	r1 := regexp.MustCompile("used_memory:([0-9]+)\r\n")
	r2 := regexp.MustCompile("[0-9]+")

	usedMemoryBytes := r1.Find([]byte(memoryInfoStr))
	curUsedMemory := r2.Find(usedMemoryBytes)

	iMemory, _ := strconv.ParseUint(string(curUsedMemory), 10, 64)

	return iMemory
}

func CalcValuesMemory(ctx context.Context, cli *redis.Client, valueSize int, insertNum int) {
	fmt.Println("--------------------------------------------------------")
	if valueSize < 1024 {
		fmt.Printf("Start Insert Values. ValueSize: %d byte, Num: %d\n", valueSize, insertNum)
	} else {
		fmt.Printf("Start Insert Values. ValueSize: %.02f k, Num: %d\n", float64(valueSize)/float64(1024), insertNum)
	}

	cli.FlushDB(ctx)
	beforeMemory := GetCurrentRedisUsedMemory(ctx, cli)
	insertToRedis(ctx, cli, valueSize, insertNum)
	afterMemory := GetCurrentRedisUsedMemory(ctx, cli)
	cli.FlushDB(ctx)

	averageSize := (float64(afterMemory) - float64(beforeMemory)) / float64(insertNum)

	fmt.Printf("Before Memory: %d byte\n", beforeMemory)
	fmt.Printf("After Memory: %d byte\n", afterMemory)
	fmt.Printf("Average Memory Usage for key: %f byte\n", averageSize)
}

func insertToRedis(ctx context.Context, cli *redis.Client, valueSize int, insertNum int) {

	iValue := buildInsertValue(valueSize)
	pipe := cli.Pipeline()

	for i := 0; i < insertNum; i++ {
		key := fmt.Sprintf("%d", i)
		pipe.Set(ctx, key, iValue, -1)

		if i%1000 == 0 {
			_, err := pipe.Exec(ctx)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildInsertValue(bSize int) string {
	return strings.Repeat("1", bSize)
}
