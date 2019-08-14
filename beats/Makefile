RM = rm -rf
GOBUILD = go build
GOTEST = go test
APP = bin/offertrack

# go build -mod=vendor 开启vendor模式

.PHONY: main test clean

main:
	${GOBUILD} -mod=vendor -o ${APP} src/main.go

test:
	pushd src/util && ${GOTEST} && popd

clean:
	${RM} bin logs
