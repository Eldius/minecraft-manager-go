# minecraft-manager-go

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/Eldius/minecraft-manager-go)
![Go](https://github.com/Eldius/minecraft-manager-go/workflows/Go/badge.svg)

## snippets ##

```shell
minecraft-manager-go install \
		-n mineserver001 \
		-u ssh_user \
		-k ~/.ssh/id_ed25519 \
		-p 2222 \
		-c ssh \
		127.0.0.1

```

```shell
go run main.go install -n mineserver001 -u ssh_user -k ~/.ssh/id_ed25519 -p 2222 -c ssh 127.0.0.1
```

## links ##

- [ansible-go](https://github.com/apenella/go-ansible)
- [gobuffalo/packr](https://github.com/gobuffalo/packr)
- [Mojang Versions API](https://launchermeta.mojang.com/mc/game/version_manifest.json)
