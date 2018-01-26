package main

import (
    "github.com/alienteam/omok-server/chat"
)
var userid int = 0

func auth(id, pw string) int {
    userid = userid+1
    return userid
}

func main(){
    s := chat.NewChatServer(":8080", "/ws", auth)
    s.Start()
}

