.PHONY: run

all: scolor/scolor

scolor/scolor: scolor/main.go
	go build -o $@ ./scolor

run: scolor/scolor
	./scolor/scolor -f fix.txt -pos 2,2

clean:
	rm scolor/scolor
