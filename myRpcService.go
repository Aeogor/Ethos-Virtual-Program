package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"log"
)


func main() {
	
	altEthos.LogToDirectory("test/myRpcServer")
	log.Printf("ChatServer: Initializing...\n")
	
	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Printf("Advertising service failed: %s\n", status)
		altEthos.Exit(status)
	}
	log.Printf("ChatServer: Done advertising...\n")

	for {
		var s Message
		userName, netFd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Println("Import failed")
			return
		}

		status = altEthos.PeekStream(netFd, &s)
		if status != syscall.StatusOk {
			log.Println("Peek Failed")
			return
		}

		if(s.fromUser != string(userName)) {
			log.Println("From User isnt same!")
			return
		}

		log.Println("from username: " + string(userName))
		log.Println("to   username: " + string(s.toUser))

		status = altEthos.FdSend([]syscall.Fd {netFd}, string(s.toUser), "virtualProgramRead")
		if status != syscall.StatusOk {
			log.Println("FdSend Failed", status)
			return
		}

		
		log.Println("after fdSend")

		altEthos.Close(netFd)
		
	
	}
}
