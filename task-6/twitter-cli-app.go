package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	handlePtr := flag.String("handle", "TwitterAPI", "Twitter handle")
	api_client := anaconda.NewTwitterApiWithCredentials("<your token>", "<your token>", "<your token>", "<your token>")
	var user anaconda.User
	user, _ = api_client.GetUsersShow(*handlePtr, nil)
	userJSON, _ := json.MarshalIndent(user, "", " ")
	//fmt.Println(string(userJSON))
	f, _ := os.Create("./output.txt")
	f.WriteString(string(userJSON))
	f.Sync()
	f.Close()
}
