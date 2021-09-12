## 10 byte数据
```shell
====== SET ======
  100000 requests completed in 1.71 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

96.76% <= 1 milliseconds
99.98% <= 2 milliseconds
100.00% <= 2 milliseconds
58445.36 requests per second
====== GET ======
  100000 requests completed in 1.71 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

97.09% <= 1 milliseconds
100.00% <= 1 milliseconds
58411.21 requests per second
```
## 20 byte数据
```shell
====== SET ======
  100000 requests completed in 1.73 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

96.61% <= 1 milliseconds
99.90% <= 4 milliseconds
99.94% <= 5 milliseconds
99.95% <= 9 milliseconds
99.96% <= 10 milliseconds
100.00% <= 10 milliseconds
57971.02 requests per second

====== GET ======
  100000 requests completed in 1.57 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

98.35% <= 1 milliseconds
100.00% <= 1 milliseconds
63775.51 requests per second
```
## 50 byte数据
```shell
====== SET ======
  100000 requests completed in 1.71 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

97.97% <= 1 milliseconds
100.00% <= 1 milliseconds
58343.06 requests per second

====== GET ======
  100000 requests completed in 1.61 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

97.44% <= 1 milliseconds
100.00% <= 1 milliseconds
61996.28 requests per second
```
## 100 byte数据
```shell
====== SET ======
  100000 requests completed in 1.73 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

96.88% <= 1 milliseconds
100.00% <= 2 milliseconds
57971.02 requests per second

====== GET ======
  100000 requests completed in 1.54 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

97.92% <= 1 milliseconds
100.00% <= 1 milliseconds
64808.82 requests per second
```

## 500 byte数据
```shell
====== SET ======
  100000 requests completed in 1.71 seconds
  50 parallel clients
  500 bytes payload
  keep alive: 1

97.17% <= 1 milliseconds
100.00% <= 1 milliseconds
58343.06 requests per second

====== GET ======
  100000 requests completed in 1.57 seconds
  50 parallel clients
  500 bytes payload
  keep alive: 1

97.79% <= 1 milliseconds
99.98% <= 2 milliseconds
100.00% <= 2 milliseconds
63653.72 requests per second
```

## 1k byte数据
```shell
====== SET ======
  100000 requests completed in 1.77 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

96.62% <= 1 milliseconds
99.97% <= 2 milliseconds
100.00% <= 2 milliseconds
56497.18 requests per second

====== GET ======
  100000 requests completed in 1.76 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

96.58% <= 1 milliseconds
100.00% <= 1 milliseconds
56785.91 requests per second
```
## 5k byte 数据
```shell
====== SET ======
  100000 requests completed in 2.01 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

86.29% <= 1 milliseconds
99.89% <= 2 milliseconds
99.90% <= 4 milliseconds
99.91% <= 5 milliseconds
99.95% <= 11 milliseconds
100.00% <= 11 milliseconds
49825.61 requests per second

====== GET ======
  100000 requests completed in 1.92 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1

76.49% <= 1 milliseconds
100.00% <= 2 milliseconds
100.00% <= 2 milliseconds
52137.64 requests per second
```

## 关于写入数据分析数据占用大小的问题
这个之前有做过其他方面的测试，每个数据大小相关的更多是单个数据大小的问题；
比较小的数据是使用压缩列表，占用的空间比较小;
当数据变大后会使用哈希表