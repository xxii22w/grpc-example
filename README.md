### 构建简单的grpc服务
- module1 ~ module2
1. 定义 gRPC 服务协定
2. 实现 Server
3. 实现客户端
4. 错误处理
5. 练习 - 构建 TODO 应用程序
6. Exercise Solution
7. 正确运行 Server
### grpc流式处理
- module3
1. 实现服务器流式处理
2. 实现客户端流式处理
3. 实现双向流式处理
4. 练习 - 构建文件上传应用程序
5. Exercise Solution
### 身份验证
- module4
1. 什么是 SSL/TLS？
2. 在 gRPC 中实现服务器端 TLS
3. 在 gRPC 中实施 mTLS  (双向TLS（mTLS）是指在服务器端和客户端之间使用双向加密通道。)
4. 练习：将 mTLS 添加到您的文件上传应用程序
6. Exercise Solution
### 拦截器、元数据和授权
- module5
1. 拦截器 - 简介， 客户端拦截器， 服务器拦截器
2. 设置截止时间/超时
3. CallOptions & 元数据
4. 通过拦截器进行 API 密钥授权
5. 通过 CallCredentials 进行身份验证
6. 练习：添加侦听器以解析 JWT 信息
7. Exercise Solution
### 客户端服务配置和负载均衡
- module6
1. 客户端服务配置和超时
2. 客户端自动重试
3. 客户端负载均衡简介
4. 循环负载均衡
5. 创建自定义负载均衡策略

https://www.bytesizego.com/grpc-with-go