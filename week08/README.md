# 作业内容：

### 1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

redis benchmark 文档: https://redis.io/docs/reference/optimization/benchmarks/

```
操作系统: MacOS
redis版本: 7.0.2
环境配置: 8CPU, 16G
```

```shell
redis-benchmark -q -n 100000 -t set,get -d 10
SET: 198412.69 requests per second, p50=0.135 msec
GET: 210526.31 requests per second, p50=0.135 msec

redis-benchmark -q -n 100000 -t set,get -d 20
SET: 194931.77 requests per second, p50=0.135 msec
GET: 208333.34 requests per second, p50=0.135 msec

redis-benchmark -q -n 100000 -t set,get -d 50
SET: 187969.92 requests per second, p50=0.143 msec
GET: 204918.03 requests per second, p50=0.135 msec

redis-benchmark -q -n 100000 -t set,get -d 100
SET: 191938.56 requests per second, p50=0.135 msec
GET: 186219.73 requests per second, p50=0.143 msec

redis-benchmark -q -n 100000 -t set,get -d 200
SET: 194552.53 requests per second, p50=0.143 msec
GET: 190839.70 requests per second, p50=0.135 msec

redis-benchmark -q -n 100000 -t set,get -d 1024
SET: 193798.45 requests per second, p50=0.143 msec
GET: 199203.20 requests per second, p50=0.135 msec

redis-benchmark -q -n 100000 -t set,get -d 5120
SET: 182815.36 requests per second, p50=0.151 msec
GET: 192307.70 requests per second, p50=0.151 msec
```


### 2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

```shell

--------------------------------------------------------
Start Insert Values. ValueSize: 127 byte, Num: 10000
Before Memory: 1259920 byte
After Memory: 3536528 byte
Average Memory Usage for key: 227.660800 byte
--------------------------------------------------------
Start Insert Values. ValueSize: 512 byte, Num: 10000
Before Memory: 1259920 byte
After Memory: 7376528 byte
Average Memory Usage for key: 611.660800 byte
--------------------------------------------------------
Start Insert Values. ValueSize: 1.00 k, Num: 10000
Before Memory: 1259920 byte
After Memory: 17505168 byte
Average Memory Usage for key: 1624.524800 byte
--------------------------------------------------------
Start Insert Values. ValueSize: 5.00 k, Num: 10000
Before Memory: 1308560 byte
After Memory: 58450320 byte
Average Memory Usage for key: 5714.176000 byte
--------------------------------------------------------
--------------------------------------------------------
Start Insert Values. ValueSize: 127 byte, Num: 500000
Before Memory: 1293712 byte
After Memory: 109489552 byte
Average Memory Usage for key: 216.391680 byte
--------------------------------------------------------
Start Insert Values. ValueSize: 512 byte, Num: 500000
Before Memory: 1295472 byte
After Memory: 301489776 byte
Average Memory Usage for key: 600.388608 byte
--------------------------------------------------------
Start Insert Values. ValueSize: 1.00 k, Num: 500000
Before Memory: 1295664 byte
After Memory: 805489968 byte
Average Memory Usage for key: 1608.388608 byte
--------------------------------------------------------
Start Insert Values. ValueSize: 5.00 k, Num: 500000
Before Memory: 1295856 byte
After Memory: 2853490160 byte
Average Memory Usage for key: 5704.388608 byte


```