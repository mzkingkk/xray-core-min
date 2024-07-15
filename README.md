# 说明

xray的包越来越大，小内存设备快要没法玩了，所以在前人的思路下自己删减搞一个能用的版本。
尚未删除到极致，还有好多保留的待删除部分，有空了再研究研究。  
修改前：29m  
修改后：19m  
压缩后：4.76m  

## 保留功能
- [x] http，端口监听可用
- [x] dokodemo-door，透明代理可用
- [x] grpc
- [x] dns，router相关均未删除

## 支持的配置
打勾的表示测试ok，未打勾的没好好测过，也许支持吧
- [x] VLESS reality_gRPC
- [ ] VLESS reality_vision
- [ ] VLESS gRPC TLS
- [ ] VLESS TCP TLS_Vision

## 删除内容
全在提交里面了，想要还原哪个模块，直接revert应该就行
- [x] ~~delete vmess~~
- [x] ~~delete quic~~
- [x] ~~delete trojan~~
- [x] ~~delete shadowsocks~~
- [x] ~~delete splithttp~~
- [x] ~~delete wireguard~~
- [x] ~~delete wechat~~
- [x] ~~delete kcp~~
- [x] ~~delete httpupgrade~~
- [x] ~~delete domainsocket~~
- [x] ~~delete websocket~~
- [x] ~~delete srtp~~
- [x] ~~delete yaml~~
- [x] ~~delete toml~~
- [x] ~~delete udp~~
- [x] ~~delete socks~~
- [x] ~~delete server.go~~
- [x] ~~delete test.go~~

# Linux / macOS / WSL

```bash
CGO_ENABLED=0 GOARCH=mipsle GOMIPS=softfloat go build -o xray -trimpath -ldflags "-s -w -buildid=" ./main
```

# 鸣谢
原始代码全部来自xray-core，v1.8.17，这里只是删删删：  
https://github.com/XTLS/Xray-core

思路来源是如下项目(原项目没研究出来怎么支持dokodemo-door且不会内存溢出，所以自己删删删，删出了本项目)：  
https://github.com/wangz-code/xray-core-min
