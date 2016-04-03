
server_tests:
	bash -c "cd src/server && go test"

tests: server_tests

stylecheck:
	bash -c "mkdir -p temp; cd temp && export GOPATH=`pwd`/temp && go get github.com/qiniu/checkstyle/gocheckstyle"
	temp/bin/gocheckstyle -config=.gostyle .
	@rm -fr temp
