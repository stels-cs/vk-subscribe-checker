package main

import (
	"github.com/stels-cs/vk-api-tools"
)

import "flag"
import (
	"log"
	"strconv"
	"encoding/json"
	"fmt"
)

type Allow struct {
	Allowed int `json:"is_allowed"`
}

func PrintAllowedIds(api *Vk.Api, pack *Vk.ExecutePack, remap *map[int]int) int  {
	allow := 0
	res, err := api.Execute(pack.GetCode())
	if err != nil {
		log.Fatal(err)
	}
	allowResults := make([]Allow, 0)
	err = json.Unmarshal(*res.Response, &allowResults)
	if err != nil {
		println( string([]byte(*res.Response)) )
		println(res.ExecuteErrors[0].Message)
		log.Fatal(err)
	}
	for i,a := range allowResults {
		if a.Allowed == 1 {
			fmt.Println((*remap)[i])
			allow++
		}
	}
	return allow
}

func main() {

	inputFile := flag.String("input", "", "input file with vk user ids")
	//outputFile := flag.String("output", "", "output file with ids who's allow messages from community")
	token := flag.String("token", "", "access token for api vk")

	flag.Parse()

	if *inputFile == "" {
		log.Fatal("No input file passed, pass -h for see help")
	}

	if *token == "" {
		log.Fatal("No token file passed (It's used for call vk api), pass -h for see help")
	}

	r, err := GetReadet(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	pack := Vk.ExecutePack{}
	api := Vk.GetApi(Vk.AccessToken{Token: *token}, Vk.GetHttpTransport(), nil)
	group, err := api.Groups.GetMe()
	if err != nil {
		log.Println("Seems token is bad or it's not grouo token")
		log.Fatalln(err)
	}

	log.Println("Checking ids for group "+group.Name)

	var e error
	var id int
	remap := make(map[int]int, 0)

	totalIds := 0
	allowWriteIds := 0

	for id, e = r.GetNexId(); e == nil; id, e = r.GetNexId() {
		totalIds++
		req := Vk.GetApiMethod(
			"messages.isMessagesFromGroupAllowed",
			Vk.Params{"group_id": strconv.Itoa(group.Id), "user_id": strconv.Itoa(id)})
		index, err := pack.Add(req)
		if err != nil {
			log.Fatal(err)
		}
		if index != -1 {
			remap[index] = id
		} else {
			allowWriteIds += PrintAllowedIds(api, &pack, &remap)
			pack.Clear()
			remap = make(map[int]int, 25)
			index, err := pack.Add(req)
			if err != nil {
				log.Fatal(err)
			}
			remap[index] = id
		}
	}

	if pack.Count() > 0 {
		allowWriteIds += PrintAllowedIds(api, &pack, &remap)
	}

	if e != nil && e.Error() != "EOF" {
		log.Fatal(e)
	}

	log.Println("DONE")
	log.Println("Total ids in file", totalIds)
	log.Println("Ids who's allow write", allowWriteIds)
}
