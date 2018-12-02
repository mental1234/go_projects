MKDIR = mkdir -p
PRINT = echo
BUILD = bin
SOURCE = src

run:
	@go run $(FILE)

buildf:
	${MKDIR} ${BUILD} && cd ${BUILD}
	@go build ../$(FILE)

project:
	${MKDIR} $(PROJ)/${BUILD}
	${MKDIR} $(PROJ)/${SOURCE}
	cp go_helloWorld.go $(PROJ)/${SOURCE}/main.go

buildp:
	${MKDIR} $(PROJ)/${BUILD}
	cd $(PROJ)/${BUILD} && go build ../${SOURCE}/$(FILE)
