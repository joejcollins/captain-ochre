[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/joejcollins/captain-ochre)

# Golang Examples

<https://lets-go.alexedwards.net/>


## Docker Demo

```
docker build -t captain-ochre .
docker run -it --rm --volume /workspace/captain-ochre:/workspace/captain-ochre --name captain-ochre captain-ochre
docker exec -it boring_perlman /bin/sh
```