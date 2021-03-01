#断点续传


参考资料
* socket:https://blog.csdn.net/luckytanggu/article/details/79830493

* http: https://mojotv.cn/go/go-range-download


```markdown
模拟binlog同步

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
