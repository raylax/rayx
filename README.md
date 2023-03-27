# [xray](https://github.com/chaitin/xray) poc 执行器
[![Go](https://github.com/raylax/rayx/actions/workflows/build.yml/badge.svg)](https://github.com/raylax/rayx/actions/workflows/build.yml)
[![codecov](https://codecov.io/github/raylax/rayx/branch/master/graph/badge.svg?token=l8Fl0uumvb)](https://codecov.io/github/raylax/rayx)
[![Go Report Card](https://goreportcard.com/badge/github.com/raylax/rayx)](https://goreportcard.com/report/github.com/raylax/rayx)

> 仅供个人学习、研究使用，请勿用于非法用途

# 使用

[Releases](https://github.com/raylax/rayx/releases)页面包含最新下载地址
```shell
# 
curl -L https://github.com/raylax/rayx/releases/download/0.1.0/rayx_darwin_amd64 -o rayx
chmod +x rayx
./rayx -p _testdata/pocs -u http://localhost:1234

# zeroshell-cve-2019-12725-rce.yml - [H] √√√
# zimbra-cve-2019-9670-xxe.yml - [M]
# zzcms-zsmanage-sqli.yml - [E]

```
`[H]` 命中规则

`[M]` 未命中规则

`[E]` 执行错误


# 帮助 `./rayx -h`
```
Usage:
  rayx [flags]

Flags:
      --cookie        (eg cookie1=value1,cookie2=value2)
      --header        (eg header1=value1,header1=value2)
  -h, --help          help for rayx
  -p, --poc string    POC file/folder (eg. path/to/poc,rce.yaml)
  -s, --sleep int     sleep seconds
  -t, --threads int   threads (default 1)
      --timeout int   timeout seconds (default 30)
  -u, --url string    target url (eg. https://httpbin.org)
  -v, --version       version for rayx

```

# 工作原理
使用语法分析器分析`expression`构造语法树以及上下文并执行，执行的过程是惰性的


```
name: poc-yaml-demo
manual: true
transport: http
set:
    r1: randomLowercase(16)
    r2: randomLowercase(16)
rules:
    r0:
        request:
            cache: true
            path: /a/{{r1}}
        expression: response.status == 200 
    r1:
        request:
            cache: true
            path: /b/{{r1}}
        expression: response.status == 200
expression: r0() && r1()
detail:
    author: raylax(raylax@inurl.org)
```
1. 执行器首先解析最外层表达式`r0() && r1()`
2. 解析`rules.r0`，由于依赖`set.r1`，所以计算`set.r1`
3. 计算出所有依赖值后执行`rules.r0`，当`rules.r0.expression`为`false`时，整个POC验证失败，结束验证
4. 如果`rules.r0.expression`为`true`则继续重复上述步骤
5. `set`中的变量会缓存计算结果，也就是说如果`set.r1=randomLowercase(16)`重复使用`set.r1`的值是一样的
6. 如果`rules`配置了`cache: true`同样会缓存结果

# 问题

1. 由于暂未实现反连功能，所以使用反连的poc会报`error - 'newReverse' undefined` 错误
