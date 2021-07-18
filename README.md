## docker-registry-search

```bash
$ ./install.sh
$ docker-registry-search search --registry registry.example.com [--https] <term>
```

### Create alias
```bash
alias drs="docker-registry-search search --registry registry.example.com --https \"$@\""
```
