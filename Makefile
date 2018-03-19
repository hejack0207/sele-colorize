.PHONY: run

all: scolor/scolor

scolor/scolor: scolor/main.go
	go build -o $@ ./scolor

run: scolor/scolor
	./scolor/scolor -f fix.txt

clean:
	rm scolor/scolor
