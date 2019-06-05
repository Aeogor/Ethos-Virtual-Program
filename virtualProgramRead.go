package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)



func main() {

	altEthos.LogToDirectory("test/myRpcVP")

	var composedMessage Message
	var path = "/user/" + altEthos.GetUser() + "/mailbox/"

	fd, status := altEthos.FdReceive()
	if status != syscall.StatusOk {
		log.Println("FdReceive Failed")
		return
	}

	log.Println("FdReceive Successful")

	status = altEthos.ReadStream(fd, &composedMessage)
	if status != syscall.StatusOk {
		log.Println("Read Failed")
		return
	}

	checkDirectory := altEthos.IsDirectory(path + "incoming/")
	if checkDirectory == false {
		log.Println("Directory does not exist ", path, checkDirectory)
		log.Println("Creating Directory")
		
		status = altEthos.DirectoryCreate(path, &composedMessage, "all")
		if status != syscall.StatusOk {
			log.Println("Directory Create 'mailbox' Failed ", path, status)
			return
		}

		status = altEthos.DirectoryCreate(path + "incoming/", &composedMessage, "all")
		if status != syscall.StatusOk {
			log.Println("Directory Create 'incoming' Failed ", path, status)
			return
		}
	}

	fd2, status1 := altEthos.DirectoryOpen(path + "incoming/")
	if status1 != syscall.StatusOk {
		log.Println("Directory Create Failed ", path, status1)
		return
	}

	status = altEthos.WriteStream(fd2, &composedMessage)
	if status != syscall.StatusOk {
		log.Println("Directory write Failed ", path, status)
		return
	}

	log.Println("After write stream")
}
