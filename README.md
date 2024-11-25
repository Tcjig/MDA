# **MDA**

## **/cmd**
程序的主入口，存放了我的main.go文件

## **/package**
被调用的包
### **/tcp_connection**
负责服务端和客户端的连接

server：监听和接收客户端的请求

client：向服务端请求连接
### **/file_function**
文件的传输和接收功能
#### /file_delivery
文件的传输，被客户端调用

delivary_PathCheck：传输路径的输入和检查

file_delivary：负责文件的传输
#### /file_storage
文件的接收，被服务端调用

storage_PathCehck：存储路径的输入和检查

file_storage：负责文件的接收
