all: scolor/scolor

scolor/scolor: scolor/main.go
	go build -o $@ ./scolor

run: scolor/scolor
	./scolor/scolor -f README.md

clean:
	rm scolor/scolor
