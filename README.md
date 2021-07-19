## docker-registry-search

### Usage
```bash
$ docker-registry-search search --registry registry.example.com [--https] <term>
```
Returns all the images that contain the search term.
You can select the one you want from the list and it will be pulled.

### Create alias
```bash
alias drs="docker-registry-search search --registry registry.example.com --https ${@}"
```
