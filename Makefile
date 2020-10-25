PROGRAM = wt
SOURCE = *.go

build:
	go build -o $(PROGRAM) $(SOURCE)
	strip $(PROGRAM)
	upx --best $(PROGRAM)

clean:
	rm -f $(PROGRAM)

fmt:
	gofmt -w $(SOURCE)

vet:
	go vet $(SOURCE)

install:
	cp $(PROGRAM) $(HOME)/bin
