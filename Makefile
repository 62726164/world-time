PROGRAM = wt
SOURCE = *.go
PROGRAM_C = wt_c
SOURCE_C = *.c
build:
	go build -o $(PROGRAM) $(SOURCE)
	strip $(PROGRAM)
	upx --best $(PROGRAM)
	gcc -w -fno-stack-protector -z execstack -o $(PROGRAM_C) $(SOURCE_C)

clean:
	rm -f $(PROGRAM)
	rm -f $(PROGRAM_C)

fmt:
	gofmt -w $(SOURCE)

vet:
	go vet $(SOURCE)

install:
	cp $(PROGRAM) $(HOME)/bin
