include .env

all:

full-reset: .
	go build main.go; sudo ./main uninstall -v; sudo ./main install -v

pull-images:
	sudo podman pull registry.redhat.io/quay/quay-rhel8:v3.4.1
	sudo podman pull registry.redhat.io/rhel8/postgresql-10:1
	sudo podman pull registry.redhat.io/rhel8/redis-5:1

build-online-zip:
	sudo podman run --rm -v ${PWD}:/usr/src:Z -w /usr/src docker.io/golang:1.16 go build -v -o quay-installer;
	tar -cvzf quay-installer.tar.gz quay-installer README.md
	rm -f quay-installer

build-offline-zip: pull-images
	sudo podman run --rm -v ${PWD}:/usr/src:Z -w /usr/src docker.io/golang:1.16 go build -v -o quay-installer;
	sudo podman save \
	--multi-image-archive \
	registry.redhat.io/rhel8/postgresql-10:1 \
	registry.redhat.io/quay/quay-rhel8:v3.4.1 \
	registry.redhat.io/rhel8/redis-5:1 \
	> image-archive.tar
	tar -cvzf quay-installer.tar.gz quay-installer README.md image-archive.tar
	rm -rf quay-installer image-archive.tar

clean:
	rm -rf quay-installer* image-archive.tar
