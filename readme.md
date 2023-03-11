#### 实现一个单机版本的对象存储

1. 支持get和put操作
2. 展示如下：
   1. 启动服务器：LISTEN_ADDRESS=:12345 STORAGE_ROOT=/tmp go run server.go
   2. 通过curl模拟客户端操作：
      1. get操作：curl 10.10.103.176:12345/objects/test -v

         ![1678542919981](image/readme/1678542919981.png)
      2. put操作：curl -v 10.10.103.176:12345/objects/test -XPUT -d"this is a test objects"

         ![1678542968295](image/readme/1678542968295.png)
      3. 观察存储位置：/tmp/objects/

         ![1678543033019](image/readme/1678543033019.png)
      4. 再一次get

         ![1678543086314](image/readme/1678543086314.png)
