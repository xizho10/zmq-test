# zmq-test
zeromq PULL PUSH

1.前端 –>后端：PUSH/PULL模式，前端是Bind，每个后端连接到所有前端。后端先PULL先得，能负载均衡。

2.后端-->前端：PUB/SUB模式，后端是Bind，后端可以PUB到指定前端。

3.Redis缓存: 保存client和device连接地址。key：username，value：FrontendName.  前端负责写，后端负责读。

--------------------------------
-  haproxy + keepalived        -
-                              -
-                              -
-   frontend + backend         -
-                              -
-    mongodb                   -
-                              -
-                              -
--------------------------------

参考 https://github.com/paulobsf/go-zeromq-tests
