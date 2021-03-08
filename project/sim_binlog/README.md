# 接收Binlog 流程

* 握手协议 HandShake协议

* show master status 返回binlogFileName, Position

* show global variables like 'binlog_checksum' 返回checksum 类型

* 发送 binlog dump 命令

* 循环读取binlog event 流
    * DumpBinlog包

# 需求
```markdown
有两个文件：
一个是SQL文件，内容是建表语句，格式由SHOW CREATE TABLE语句获得；
一个是CSV文件，有不定行数、N+1列，N是那个SQL文件的表的列数目。CSV第1-N列与表的N列对应，N+1列可能是“I”或者“D”，分别表示该行是插入或者删除
这两个文件可以表示MySQL表内容的变更日志。
1：编写一个程序，正确地将变更日志同步到一个MySQL实例中（连接MySQL实例的参数可以硬编码）
2：如果仅在同步完成后要求数据一致，同步过程中每个时刻不要求“数据与上游对应时刻的数据一致”，有什么提升速度的方法。大体描述方法并设计测试
扩展 3a：实现 2
3b：SQL文件支持其他合法的、但格式与SHOW CREATE TABLE结果不同的建表语句
3c：实现简单的断点续传
```

* client 客户端
    * 负责发送两个文件
    
* 主服务端
    * 启动方式 go run main.go
    * 功能点：模拟中继日志，缓存数据。断点续传接收文件等。
