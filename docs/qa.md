## Q & A

- 关于之前区块信息的导入
- Redis / Kafka 的持久化
- PubSub 队列的去重
- 检索
- 定时备份原始数据

## 关于之前区块信息的导入优化

1. 通过 Handler 去解析 RPC 命令
2. 

### Redis / Kafka 的持久化

Redis AOF

### PubSub 队列的去重

1.kafka 队列去重

幂等producr

EOS

2.Redis 发布订阅去重

利用 Set

BloomFilter 去重

### 检索

1.ElasticSearch 等第三方工具？