Name  : Srinivas C Lingutla
UIN   : 655115444
NetID : slingu2

Platform USED: ETHOS

------------------------------ Homework 3 - CS 485 -----------------------------------

 === BUILD AND RUN ====
 make install
 cd server
 sudo -E ethosRun

This currently cleans the folder and builds all the files
This will start the ethos instance

In an another terminal, navigate to server folder

 etAl server.ethos

This will give you the terminal access for the ethos instance
Navigate to /programs folder and run myRpcClient

 cd /programs
 myRpcClient

Similarly you can connect another user using 
 
 et server.ethos

--------------- IMPLEMENTATION ---------------

Currently on start the client program will ask for the user to enter a 
command. There are three commands to choose from. One to get all the 
messages from your inbox, one to compose a new message and send it. 
Another to exit the program. The following are the commands. 

Commands
---------------------
Enter (\n)   : get all messages
-compose      : send the message
-exit        : exit program
---------------------


The the user tries to compose a new message, they will be asked to enter
the following details. 
toUser: 
Subject:
Body: 

The body can be multiple lines of text, so the user can enter to get to
a new line. In order to finish the body text, the user will have to 
enter on a empty line, or press enter again. Once the message is written, 
the client makes an IPC connection with the service which verifies it and
passes it along to the Virtual programs. The virtual program creates the 
required directories and enters the message in the user's directory. 

-------------------------------------------------------
FILES INCLUDED

|-- Makefile
|-- Message.t
|-- README.txt
|-- myRpcClient.go
|-- myRpcService.go
|-- virtualProgramRead.go


0 directories, 6 files

-------------------------------
Sample Run of the program (Running on SrinivasL)

[SrinivasL@server.ethos /programs]$ myRpcClient
Enter Input (?? for commands) : ??


Commands
---------------------
Enter (\n)   : get all messages
-compose      : send the message
-exit        : exit program
---------------------

Enter Input (?? for commands) : -compose


Composing New Message
---------------------
toUser: sid
Subject: Hello

Press Enter on a Empty Message to Finish
---------------------
Body: Testing
Body: Message
Body:
Enter Input (?? for commands) : -compose


Composing New Message
---------------------
toUser: SrinivasL
Subject: Hello Again

Press Enter on a Empty Message to Finish
---------------------
Body: Testing
Body: Another
Body: Message
Body:
Enter Input (?? for commands) :


Getting All Messages
---------------------
From: SrinivasL
To: SrinivasL
Subject: Hello Again

Testing
Another
Message
---------------------
Enter Input (?? for commands) : -exit
[SrinivasL@server.ethos /programs]$
