package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
)

type UserMessage struct{
	UserID string
	Message string
}
type RPC struct {
	DB map[string]string
}

var DATABASE []RPC
//var database =  map[User]Message{}

func (r *RPC) GetMessage(usrMsg UserMessage, reply *RPC) error{
	//fmt.Println("***GetMessage***")
	var str []string
	reply.DB = make(map[string]string)
	key := usrMsg.UserID
	for _,val := range DATABASE{
		//fmt.Printf("%s,%s\n","***key***", key)
		if v, ok := val.DB[key];ok {
		//	fmt.Printf("%s,%s\n","***value***", v)
			str = strings.Split(v,"|")
		}
	}
	message := ""
	for k,v := range str{
		index:= strconv.Itoa(k)
		message = message + index+"->"
		message = message + v + " "
	}
	reply.DB[key] = message

	return nil
}
func (r *RPC) AddKeyMessage(usrMsg UserMessage, reply *RPC) error{
	//fmt.Println("***AddMessage*** ")
	reply.DB = make(map[string]string)
	key := usrMsg.UserID
	msg := usrMsg.Message
	for _,val := range DATABASE{
		if _, ok := val.DB[key]; ok{
			log.Println(key , " Key is already present")
			return errors.New("key is already present")
		}
	}
	if usrMsg.UserID != ""{
		k := usrMsg.UserID
		reply.DB[k] = msg
	}
	DATABASE = append(DATABASE,*reply)
	return nil
}
func(r *RPC) EditMessage(usrMsg UserMessage, reply *RPC) error{
	message := ""
	key := usrMsg.UserID
	msg := usrMsg.Message
	flag := false
	reply.DB = make(map[string]string)
	for k,val := range DATABASE{
		if v, ok := val.DB[key]; ok{
			message = v + "|" + msg
			DATABASE[k].DB[key] = message
			flag = true
		}
	}
	if !flag{
		log.Println(key , " Key is not present")
		return errors.New("key is not present")
	}
	reply.DB[key] = message
	return nil
}
/*func(r *RPC) DeleteMessage(key *User, reply *Message) error{
	if key.UserID != "" {
		k := User{UserID: key.UserID}
		if _, ok := r.DB[k]; ok {
			r.DB[k] = Message{Message: ""}
		}
	}
	return nil
}*/
func main() {
		var goRPC = new(RPC)
		err := rpc.Register(goRPC)
		if err != nil{
			log.Fatal("error registering interface",err)
		}
		// whatever is registered in rpc above will become default server mux by using rpc.HandleHTTP()

		rpc.HandleHTTP()
		listener, err := net.Listen("tcp", ":8080")
		if err != nil{
			log.Fatal("Listener error",err)
		}
		log.Println("***starting***",8080)
		http.Serve(listener, nil)
		if err != nil{
			log.Fatal("error registering interface",err)
		}

	}
	func handle(conn net.Conn, ch chan string) {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan(){
			key := scanner.Text()
			ch <- key
		}
}
