package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	handlePtr := flag.String("handle", "TwitterAPI", "Twitter handle")
	api_client := anaconda.NewTwitterApiWithCredentials("93374758-MEuv7Zv4GWQkKZV7L6cTKBXJJmAMNxdJl6mKN6AFs", "NhHZ7YTEs6GwmrJ8n58hP6QCjAoVZV0r4vIpM2yprO9Vc", "3xeTO8kCsuUOonGFBpKCz4UD0", "P5h7NcN5VoQwWEeUDMkeUsB4xebtmwxAiQIx7zXrS79JPAOnig")
	var user anaconda.User
	user, _ = api_client.GetUsersShow(*handlePtr, nil)
	userJSON, _ := json.MarshalIndent(user, "", " ")
	//fmt.Println(string(userJSON))
	f, _ := os.Create("./output.txt")
	f.WriteString(string(userJSON))
	f.Sync()
	f.Close()
}
