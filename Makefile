all: scolor/scolor

scolor/scolor: scolor/main.go
	go build -o $@ ./scolor

run:
	./scolor/scolor

clean:
	rm scolor/scolor
