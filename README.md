go-archer-sync
==============

[Archer::Plugin::Rsyn](https://github.com/tokuhirom/Archer/blob/master/lib/Archer/Plugin/Rsync.pm) for Go

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

SEE ALSO
------

[Archer](https://github.com/tokuhirom/Archer)

Author
------

takeda akihito <takeda.akihito@gmail.com>

LICENCE
-------

The MIT License (MIT)

