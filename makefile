install:
	go install user/gfind

test:
	go test user/gfind

clean:
	rm -rf ./bin/*
	rm -rf ./pkg/*
