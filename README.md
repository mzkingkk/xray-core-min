# 说明

xray的包越来越大，小内存设备快要没法玩了，所以在前人的思路下自己删减搞一个能用的版本。  
尚未删除到极致，还有好多保留的应该也都能删的，有空了再继续研究。  
修改前：29m  
修改后：二进制包->18m，压缩后->4.43m，内存占用->减少40%吧

## 保留功能
- [x] inbound->http，端口监听可用
- [x] inbound->dokodemo-door，透明代理可用
- [x] protocol->vless
- [x] grpc/tls/http
- [x] routing，没删但也没测

## 支持的配置->[查看](https://github.com/mzkingkk/xray-core-min/blob/main/examples/vless_reality_grpc.json)
打勾的表示测试ok，未打勾的没好好测过，从配置上看应该是支持的  
- [x] VLESS reality_gRPC
- [ ] VLESS reality_vision
- [x] VLESS gRPC TLS
- [ ] VLESS TCP TLS_Vision

## 删除内容
全在提交里面了，想要还原哪个模块，直接revert应该就行（如果冲突，记得拉分支处理）

<details>
<summary>v1.8.17.03-4.43m</summary>

- [x] delete stats
- [x] delete freedom
- [x] delete observatory
- [x] delete app/metrics
- [x] delete app/log/command
- [x] delete app/router/command
- [x] delete commands-all-tls
- [x] delete delete commands-all-api

</details>

<details>
<summary>v1.8.17.02-4.65m</summary>

- [x] delete utp
- [x] delete loopback
- [x] delete blackhole
- [x] delete bittorrent
- [x] delete dns and fakedns

</details>

<details>
<summary>v1.8.17.01-4.76m</summary>

- [x] delete vmess
- [x] delete quic
- [x] delete trojan
- [x] delete shadowsocks
- [x] delete splithttp
- [x] delete wireguard
- [x] delete wechat
- [x] delete kcp
- [x] delete httpupgrade
- [x] delete domainsocket
- [x] delete websocket
- [x] delete srtp
- [x] delete yaml
- [x] delete toml
- [x] delete udp
- [x] delete socks
- [x] delete server.go
- [x] delete test.go

</details>

# Linux / macOS / WSL

```bash
CGO_ENABLED=0 GOARCH=mipsle GOMIPS=softfloat go build -o xray -trimpath -ldflags "-s -w -buildid=" ./main
```

# 鸣谢
原始代码全部来自xray-core，v1.8.17，这里只是删删删：  
https://github.com/XTLS/Xray-core

思路来源如下项目(原项目刚开始没研究出来怎么支持dokodemo-door且不会内存溢出，所以自己删删删，删出了本项目)：  
https://github.com/wangz-code/xray-core-min
