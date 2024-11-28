# port-forwarder

ポートフォワーディングするプログラム。


## Usage:

```sh
NAME:
   port-forwarder - 指定されたアドレス間でポートフォワーディングを行います

USAGE:
   port-forwarder [global options]

VERSION:
   1.0.0

GLOBAL OPTIONS:
   --license                                show licensesa.
   --source value, -s value, -l value       source port. (ex: 127.0.0.1:8080)
   --destination value, -d value, -f value  destination port. (ex: example.com:443)
   --help, -h                               show help
   --version, -v                            print the version
```

`0.0.0.0:8888` で待ち受け、ローカルホストの `9999` 番ポートにフォワードする場合、以下のようにする。

```sh
port-forwarder -l 0.0.0.0:8888 -f localhost:9999
```


## Limitation:

- 現状は TCP のみ


## Install:

### binary download

[Latest version](https://github.com/mikoto2000/port-forwarder/releases/latest)


### go install

```sh
go install github.com/mikoto2000/port-forwarder@latest
```


## License:

Copyright (C) 2024 mikoto2000

This software is released under the MIT License, see LICENSE

このソフトウェアは MIT ライセンスの下で公開されています。 LICENSE を参照してください。


## Author:

mikoto2000 <mikoto2000@gmail.com>
