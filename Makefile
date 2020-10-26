PROGRAM = wt
SOURCE = *.go
PROGRAM_C = wt_c
SOURCE_C = *.c
build:
	go build -o $(PROGRAM) $(SOURCE)
	strip $(PROGRAM)
	upx --best $(PROGRAM)
	gcc -ggdb3 -O1 -w -z execstack -o $(PROGRAM_C) $(SOURCE_C) -fno-stack-protector -fno-mudflap -D_FORTIFY_SOURCE=0

clean:
	rm -f $(PROGRAM)
	rm -f $(PROGRAM_C)

fmt:
	gofmt -w $(SOURCE)

vet:
	go vet $(SOURCE)

install:
	cp $(PROGRAM) $(HOME)/bin
