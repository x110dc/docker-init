### Why use this?

You're tired of typing some boilerplate for every docker-based project.  

### What does it do?

creates a `Dockerfile`, `Makefile`, and `.dockerignore` files so you don't have to
start them from scratch.  The meat of it is the `Makefile` which I found myself
be copying from project to project.

### Installation:

```
> mkdir -p ~/bin
> wget -O ~/bin/docker-init https://raw.githubusercontent.com/x110dc/docker-init/master/docker-init.sh
> chmod +x ~/bin/docker-init
> export PATH=~/bin:$PATH
```

### Usage:
```
> mkdir my-new-docker-project; cd my-new-docker-project
> docker-init
> ls -a
Dockerfile
Makefile
.dockerignore
```
### Caveats:
- this makes no attempt to check for existing files and will happily overwrite
whatever is there

### Prerequisites:
- bash
