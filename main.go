package main  

import (
	"fmt"
	"github.com/PyMarcus/rpc_chat/view"
	"github.com/PyMarcus/rpc_chat/server"

)

func main(){
	var o string 
	fmt.Print("client or server? >>")
	fmt.Scanf("%s", &o)

	if o == "client"{
		view.Start()
	}else{
		server.RunServer()
	}
}