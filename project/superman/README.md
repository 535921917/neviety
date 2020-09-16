superman --  基于Redis的极简服务发现系统

#简介
1. 通讯协议是HTTP
2. 每个服务的实例存储在zset中,value=ip:port，score是「now+TTL」
3. 提供global version 和 service verison ,global version 便于客户端快速检测 是否有变动
4. watcher 每秒检测没有保持心跳的实例，直接清理

#API
- GET http://localhost:8088/api/service/register?PSM=sample&IP=localhost&port=6000&ttl=30
- GET http://localhost:8088/api/service/cancel?PSM=sample&IP=localhost&port=6000
- GET http://localhost:8088/api/service/list?PSM=sample1&PSM=sample2
- GET http://localhost:8088/api/service/version?PSM=sample


