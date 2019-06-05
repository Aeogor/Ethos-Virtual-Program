export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


export ETHOSROOT=server/rootfs
export MINIMALTDROOT=server/minimaltdfs


.PHONY: all install clean
all:  myRpcClient myRpcService myRpcVirtualPrograms

message.go: Message.t
	$(ETN2GO) . message main $^

myRpcService: myRpcService.go message.go
	ethosGo $^ 

myRpcClient: myRpcClient.go message.go
	ethosGo $^ 

myRpcVirtualPrograms:  virtualProgramRead.go message.go
	ethosGo $^

# install types, service,
install: clean message.go myRpcClient myRpcService myRpcVirtualPrograms
	(ethosParams server && cd server && ethosMinimaltdBuilder)
	ethosTypeInstall message
	ethosServiceInstall myRpc message/Message all
	install -D  myRpcService myRpcClient         $(ETHOSROOT)/programs
	install -D virtualProgramRead 			     $(ETHOSROOT)/virtualPrograms
	ethosStringEncode /programs/myRpcService    > $(ETHOSROOT)/etc/init/services/myRpcService
	

# remove build artifacts
clean:
	sudo rm -rf server
	rm -rf message/ messageIndex/
	rm -f message.go
	rm -f myRpcService
	rm -f myRpcService.goo.ethos
	rm -f myRpcClient
	rm -f myRpcClient.goo.ethos
	rm -f virtualProgramRead
	rm -f virtualProgramRead.goo.ethos

