package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/kernelTypes"
	"log"
	"strings"	
)

var userName string


func getAllMessages(){
	printToScreen("\n\nGetting All Messages\n")
	printToScreen("---------------------\n")

	var path = "/user/" + altEthos.GetUser() + "/mailbox/incoming/"

	checkDirectory := altEthos.IsDirectory(path)
	if checkDirectory == false {
		log.Println("Directory does not exist ", path, checkDirectory)
		printToScreen("No messages to read\n\n")
		return
	}
	
	FileNames, status := altEthos.SubFiles(path)
	if status != syscall.StatusOk {
		log.Fatalf("Error fetching files in %v\n", path)
		printToScreen("Unable to get the files")
		return
	}
	for i := 0; i < len(FileNames); i++ {
		log.Printf(path, FileNames[i])
		var newMessage Message
		status = altEthos.Read(path + FileNames[i], &newMessage)
		if status != syscall.StatusOk {
			log.Fatalf("Error reading box file at %v/%v\n", path, FileNames[i])
		}

		boday := strings.Join(newMessage.body,"\n")

		message := "From: " + string(newMessage.fromUser) + 
					"\nTo: " + string(newMessage.toUser) + 
					"\nSubject: " + string(newMessage.subject) + 
					"\n\n" + boday

		printToScreen(kernelTypes.String(message))
		printToScreen("\n---------------------\n")

	}

}

func sendMessage(){
	log.Printf("Called send mail\n")

	composedMessage := getComposedMessage()

	fd, status := altEthos.Ipc("myRpc", "", &composedMessage)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	status = altEthos.WriteStream(fd, &composedMessage)
	if status != syscall.StatusOk {
		log.Println("Write failed")
		return
	}
}


func printToScreen(prompt kernelTypes.String) {  
	statusW := altEthos.WriteStream(syscall.Stdout, &prompt)
	if statusW != syscall.StatusOk {
		log.Printf("Error writing to syscall.Stdout: %v", statusW)
	}
}

func printCommands(){
	printToScreen("\n\nCommands\n")
	printToScreen("---------------------\n")
	printToScreen("Enter (\\n)   : get all messages\n")
	printToScreen("-compose      : send the message\n")
	printToScreen("-exit        : exit program\n")
	printToScreen("---------------------\n\n")

}

func userInputHandler(userInput string) {
	if (userInput == "\n"){
		getAllMessages()
	} else if (strings.Contains(userInput, "-compose")) {
		sendMessage()
	} else if (userInput == "-exit\n") {
		altEthos.Exit(syscall.StatusOk)
	} else if (userInput == "??\n"){
		printCommands()
	} else {
		printToScreen("Invalid command! Please try again\n")
		printCommands()
	}

}

func getComposedMessage() (c Message){

	printToScreen("\n\nComposing New Message\n")
	printToScreen("---------------------\n")

	printToScreen("toUser: ")

	var composedMessage Message

	composedMessage.fromUser = userName

	var toUser kernelTypes.String
	status := altEthos.ReadStream(syscall.Stdin, &toUser)
	if status != syscall.StatusOk {
			log.Printf("Error while reading syscall.Stdin: %v", status)
	}

	printToScreen("Subject: ")

	var subject kernelTypes.String
	status = altEthos.ReadStream(syscall.Stdin, &subject)
	if status != syscall.StatusOk {
			log.Printf("Error while reading syscall.Stdin: %v", status)
	}

	var body_slice []string

	printToScreen("\nPress Enter on a Empty Message to Finish \n")
	printToScreen("---------------------\n")

	for {
		printToScreen("Body: ")
		var body kernelTypes.String
		status = altEthos.ReadStream(syscall.Stdin, &(body))
		if status != syscall.StatusOk {
				log.Printf("Error while reading syscall.Stdin: %v", status)
		}

		trimmed := strings.TrimRight(string(body), "\n");
		if(trimmed == ""){
			break
		} else {
			body_slice = append(body_slice, trimmed)
		}

	}

	composedMessage.toUser = string(toUser)
	composedMessage.toUser = strings.TrimRight(composedMessage.toUser, "\n");

	composedMessage.subject = string(subject)
	composedMessage.subject = strings.TrimRight(composedMessage.subject, "\n");

	composedMessage.body = body_slice

	return composedMessage
}

func getInput(){
	for {
		printToScreen("Enter Input (?? for commands) : ")
		var userInput kernelTypes.String
		status := altEthos.ReadStream(syscall.Stdin, &userInput)
		if status != syscall.StatusOk {
				log.Printf("Error while reading syscall.Stdin: %v", status)
		}

		userInputHandler(string(userInput));
	}
}

func main () {

	altEthos.LogToDirectory("test/myRpcClient")
	
	log.Printf("EthosMailClient: before call\n")

	userName = altEthos.GetUser()

	getInput()

	log.Printf("EthosMailClient: done\n")
}
