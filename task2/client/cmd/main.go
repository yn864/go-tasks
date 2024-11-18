package cmd

import "web_app/client"

func main() {

	client := client.NewClient("http://localhost:8080")
	client.GetVersion()
	client.DecodeMessage("SGVsbG8gV29ybGQh")
	client.HardOp()

}
