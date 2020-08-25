package main

import (
	"fmt"
	"log"
	"net/rpc"
)


type UserMessage struct{
	UserID string
	Message string
}
type RPC struct {
	DB map[string]string
}
func main(){
		var usrMsg UserMessage
		var reply RPC
		client, err := rpc.DialHTTP("tcp","localhost:8080")
		if err != nil{
			log.Fatal("error connecting",err)
		}
		usrMsg.UserID = "aryan"
		usrMsg.Message= "How are You"

		//TO ADD NEW KEY WITH MESSAGE
		client.Call("RPC.AddKeyMessage", usrMsg,&reply)
		fmt.Println("Data", reply.DB)

		usrMsg.UserID = "naman"
		usrMsg.Message= "not good"
		client.Call("RPC.AddKeyMessage", usrMsg,&reply)
		fmt.Println("Data", reply.DB)

		usrMsg.UserID = "RAHUKL"
		usrMsg.Message= "that's it"
		//TO EDIT THE MESSAGE PRESENT WITH THE KEY
		client.Call("RPC.EditMessage", usrMsg,&reply)
		//fmt.Println("Data", reply.DB[usrMsg.UserID])

		usrMsg.UserID = "aryan"
		//	GET THE MESSAGE BASED ON KEY
		client.Call("RPC.GetMessage", usrMsg,&reply)
		fmt.Println("Data", reply.DB[usrMsg.UserID])

}