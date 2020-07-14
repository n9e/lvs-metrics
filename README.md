# **lvs-metrics** <sup><sub>_lvs metrics collector for n9e_</sub></sup>
[![Build Status](https://travis-ci.org/weizhenqian/lvs-metrics.svg?branch=master)](https://travis-ci.org/weizhenqian/lvs-metrics)

lvs metrics collector for n9e.

##Change Log
修改了配置文件，新增了特殊化的endpoint及自定义shell命令的两种方式。

## Dependencies

* [libnl3][]
* [ipvs][]
* [netlink][]
* [toolkits][]
* [glog][]

Dependencies are handled by [godep][], simple install it and type `godep restore` to fetch them.

## Install

#### The libnl3 needed in compile and production machine.

```console
# sudo yum install -y libnl3.x86_64
or
# sudo apt-get install libnl-3-dev libnl-genl-3-dev
```

#### Only needed in compile machine.
```console
$ git clone https://github.com/weizhenqian/lvs-metrics.git
$ cd lvs-metrics
$ make bin
```


## Configuration

Edit cfg.json configuration file.

```console
{
    "debug": false,
    "attachtags": "",
    "http": {
        "enable": false,
        "listen": "0.0.0.0:1987"
    },
    "endpoint": {
        "specify": "127.0.0.1",
        "shell": "/bin/jq -r '.local_ip' /etc/info.json"（此处我写的命令仅适配我自己的服务器，请编写适配你自己的shell命令）                       
    },
    "transfer": {
        "enable": true,
        "addr": "http://127.0.0.1:2058/api/collector/push",
        "interval": 30,
        "timeout": 1000
    }
}

or

{
    "debug": false,
    "attachtags": "",
    "http": {
        "enable": false,
        "listen": "0.0.0.0:1987"
    },
    "endpoint": {
        "specify": "127.0.0.1",
        "shell": "/bin/jq -r '.local_ip' /etc/info.json"（此处我写的命令仅适配我自己的服务器，请编写适配你自己的shell命令）                       
    },
    "transfer": {
        "enable": true,
        "addr": "127.0.0.1:8433", # Send metrics to transfer direct.
        "interval": 30,
        "timeout": 1000
    }
}
```



## Metrics

| Counters | Type | Notes |
|-----|-----|-----|
| lvs.in.bytes | GUAGE | network in bytes per host |
| lvs.out.bytes | GUAGE | network out bytes per host |
| lvs.in.packets | GUAGE | network in packets per host |
| lvs.out.packets | GUAGE | network out packets per host |
| lvs.total.conns | GUAGE | lvs total connections per host |
| lvs.vip.active_conn | GUAGE | lvs active connections per vip now |
| lvs.vip.inact_conn | GUAGE | lvs inactive connections per vip now |
| lvs.vip.rs_num | GUAGE | lvs live realserver num per vip now |
| lvs.vip.conns | COUNTER | lvs conns counter from service start per vip |
| lvs.vip.inbytes | COUNTER | lvs inbytes counter from service start per vip |
| lvs.vip.outbytes | COUNTER | lvs outpkts counter from service start per vip |
| lvs.vip.inpkts | COUNTER | lvs inpkts counter from service start per vip |
| lvs.vip.outpkts | COUNTER | lvs outpkts counter from service start per vip |


[libnl3]: https://www.infradead.org/~tgr/libnl
[ipvs]: https://github.com/google/seesaw
[netlink]: https://github.com/google/seesaw
[toolkits]: https://github.com/toolkits
[glog]: https://github.com/golang/glog
[godep]: https://github.com/tools/godep
