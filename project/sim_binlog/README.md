# 接收Binlog 流程

* 握手协议 HandShake协议

* show master status 返回binlogFileName, Position

* show global variables like 'binlog_checksum' 返回checksum 类型

* 发送 binlog dump 命令

* 循环读取binlog event 流
    * DumpBinlog包
    * 