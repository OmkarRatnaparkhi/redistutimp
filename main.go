package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type SwitchAccountModel struct {
	ParentId   string `json:"ParentId`
	ChildId    string `json:"ChildId"`
	RecordType int    `json:"RecordType"`
}

type SwitchAccountModel1 struct {
	ParentId   string   `json:"ParentId`
	ChildId    []string `json:"ChildId"`
	RecordType int      `json:"RecordType"`
}

func GenerateRedisKey(userKey string) string {
	redisKey := userKey + "_SA"
	return redisKey
}

func AddUserInSwitchAccountRedis(NewUser SwitchAccountModel) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	arrobj := []SwitchAccountModel{}

	redisgetdata, err := rdb.Get(ctx, GenerateRedisKey(NewUser.ParentId)).Result()
	fmt.Println(redisgetdata)
	if err != nil {
		data, err := json.Marshal(NewUser)
		if err != nil {
			return
		}
		err = rdb.Set(ctx, GenerateRedisKey(NewUser.ParentId), data, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		redisgetmodel, err := rdb.Get(ctx, GenerateRedisKey(NewUser.ParentId)).Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		userdatamodel := SwitchAccountModel{}

		err1 := json.Unmarshal([]byte(redisgetmodel), &userdatamodel)
		if err1 != nil {
			fmt.Println(err1)
		}
		arrobj = append(arrobj, userdatamodel)
		if userdatamodel.ParentId == NewUser.ParentId {

			arrobj = append(arrobj, NewUser)

			adddata, err := json.Marshal(arrobj)
			if err != nil {
				return
			}
			err = rdb.Set(ctx, GenerateRedisKey(NewUser.ParentId), adddata, 0).Err()
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}

func AddUserInSwitchAccountRedis1(NewUser SwitchAccountModel) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	redisgetdata, err := rdb.Get(ctx, GenerateRedisKey(NewUser.ParentId)).Result()
	fmt.Println(redisgetdata)
	if err != nil {
		data, err := json.Marshal(NewUser)
		if err != nil {
			return
		}
		err = rdb.Set(ctx, GenerateRedisKey(NewUser.ParentId), data, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {

		exists, err := rdb.Exists(ctx, GenerateRedisKey(NewUser.ParentId)).Result()
		if err != nil {
			fmt.Println("Failed to check key existence:", err)
			return
		}

		if exists == 1 {
			fmt.Println("Key exists in Redis.")

			data, err := json.Marshal(NewUser)
			if err != nil {
				return
			}
			err = rdb.Append(ctx, GenerateRedisKey(NewUser.ParentId), string(data)).Err()
			if err != nil {
				fmt.Println("Failed to append value:", err)
				return
			}
			fmt.Println("Value appended successfully.")
		} else {
			fmt.Println("Key does not exist in Redis.")
			return
		}

	}

}

type Record struct {
	ParentID   string `json:"ParentId"`
	ChildID    string `json:"ChildId"`
	RecordType string `json:"RecordType"`
}

func marshalFunction() {

	// Input string
	input := "{\"ParentId\":\"Sumant\",\"ChildId\":\"Saresh\",\"RecordType\":2}{\"ParentId\":\"Sumant\",\"ChildId\":\"Rahul\",\"RecordType\":2}{\"ParentId\":\"Sumant\",\"ChildId\":\"Prasad\",\"RecordType\":2}{\"ParentId\":\"Sumant\",\"ChildId\":\"Varun\",\"RecordType\":2}"

	// Split the input string into individual JSON objects
	objects := strings.Split(input, "}{")

	// Modify each object by adding braces
	for i := 0; i < len(objects); i++ {
		if i == 0 {
			objects[i] = "{" + objects[i] + "}"
		} else if i == len(objects)-1 {
			objects[i] = "{" + objects[i] + "}"
		} else {
			objects[i] = "{" + objects[i] + "}"
		}
	}

	// Join the modified objects into a JSON array string
	jsonArray := "[" + strings.Join(objects, ",") + "]"

	// Print the JSON array
	fmt.Println("JSON Array:", jsonArray)

	// Unmarshal the JSON array into a slice of Record objects
	var records []Record
	err := json.Unmarshal([]byte(jsonArray), &records)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON array:", err)
		return
	}

	// Print the records
	for _, record := range records {
		fmt.Println("Parent ID:", record.ParentID)
		fmt.Println("Child ID:", record.ChildID)
		fmt.Println("Record Type:", record.RecordType)
		fmt.Println()
	}
}

func main() {

	// iobj := SwitchAccountModel{"Sumant", "Saresh", 2}
	// AddUserInSwitchAccountRedis1(iobj)
	// iobj1 := SwitchAccountModel{"Sumant", "Rahul", 2}
	// AddUserInSwitchAccountRedis1(iobj1)
	// iobj2 := SwitchAccountModel{"Sumant", "Prasad", 2}
	// AddUserInSwitchAccountRedis1(iobj2)
	// iobj4 := SwitchAccountModel{"Sumant", "Varun", 2}
	// AddUserInSwitchAccountRedis1(iobj4)

	marshalFunction()

}
