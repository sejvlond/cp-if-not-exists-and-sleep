IMAGE=sejvlond/cp-if-not-exists-and-sleep

run:
	go build -o run -tags=netgo

build: run

image: build
	docker build -t ${IMAGE} .

push: image
	docker push ${IMAGE}

clean:
	rm -f run
