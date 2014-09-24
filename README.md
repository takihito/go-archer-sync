go-archer-sync
==============

archer sync for Go

Usage
-----

build and run

```
$ go build ./*.go

$ archer-sync --help

Usage of ./archer-sync:
  -config="deploy_config.yaml": yaml confige file
  -parallel=2: parallel worker

$ archer-sync --config=sample.yaml --parallel=2
```

config yaml

See sample (deploy_config.yaml).

Author
------

takeda akihito <takeda.akihito@gmail.com>

LICENCE
-------

The MIT License (MIT)

