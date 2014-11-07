package main

import "io/ioutil"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	Makefile := []byte(
		`APP=appname
NS=namespace

build:
	docker build --tag=${NS}/${APP} .

debug:
	docker run --volumes-from=${APP} --interactive=true --tty=true ${NS}/${APP} bash

run:
	docker run --rm --name=${APP} --detach=true --publish=80:8000 ${NS}/${APP}

clean:
	docker stop ${APP} && docker rm ${APP}

interactive:
	docker run --rm --interactive --tty --name=${APP} ${NS}/${APP} bash

push:
	docker push ${NS}/${APP}
`)

	err := ioutil.WriteFile("Makefile", Makefile, 0644)
	check(err)

	Dockerfile := []byte(
		`FROM ubuntu:14.04
MAINTAINER foo@bar.org
`)

	err = ioutil.WriteFile("Dockerfile", Dockerfile, 0644)
	check(err)

	err = ioutil.WriteFile(".dockerignore", []byte(".git\n"), 0644)
	check(err)

}
