package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAIC69s9vEk8sDu", "K9K6iPbIecvPGBHcXkELCgK6lxvwrG")
	if err != nil {
		handleError(err)
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		handleError(err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("bucket:", bucket.Name)
	}

	bucket, err := client.Bucket("jmdmedia")
	if err != nil {
		handleError(err)
	}
	x, err := bucket.ListObjects()
	if err != nil {
		handleError(err)
	}

	for _, object := range x.Objects {
		fmt.Println("Object: ", object.Key)
	}

	err = bucket.PutObject("hello", strings.NewReader("MyObjectValue"))
	if err != nil {
		handleError(err)
	}

	var nextPos int64 = 13
	nextPos, err = bucket.AppendObject("hello", strings.NewReader("YourObjectValue"), nextPos)
	if err != nil {
		handleError(err)
	}
}
