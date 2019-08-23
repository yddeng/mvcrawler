exe:
	test -d bin || mkdir -p bin
	cd bin;go build ../main/crawler.go;cd ../