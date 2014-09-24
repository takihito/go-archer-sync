go-archer-sync
==============

[Archer::Plugin::Rsyn](https://github.com/tokuhirom/Archer/blob/master/lib/Archer/Plugin/Rsync.pm) for Go

Usage
-----

```
$ archer-sync --help

Usage of ./archer-sync:
  -config="deploy_config.yaml": yaml confige file
  -parallel=2: parallel worker

$ archer-sync --config=sample.yaml --parallel=2
```

## Example Config

```
global:
  work_dir: /home/deploy_trunk/
  dest_dir: /home/deploy/

tasks:
  init:
    - module: Confirm
      name: confirm
      config:
        msg: "really deploy app? [y/n]"

  process:
    - module: Rsync
      name: deploy_app
      config:
        user: akihito
        source: "[% work_dir %][% project %]"
        dest: "[% server %]:[% dest_dir %]"
        dry_run: 0
        archive:  1
        compress: 1
        rsh:      ssh
        update:   1
        verbose:  1
        delete:   0
        progress: 1
        include: ['*/','*.go']
        exclude: ['*']
        filter: ['+ lib', '- .svn', '- tmp/*']

projects:
  example.com:
    servers:
      - app001
      - app002
      - app003
```


SEE ALSO
------

[Archer](https://github.com/tokuhirom/Archer)

LICENCE
-------

MIT

Author
------

Takeda Akihito <takeda.akihito@gmail.com>

