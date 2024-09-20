# 活动管理系统

## 概述

该系统通过 Kafka 和 Go 语言实现了活动管理功能。包括生成和发送消息的生产者服务和处理消息的消费者服务。

## 运行生产者
```go
go run producer.go
```

### 运行消费者
```go
go run consumer.go
```

## 数据库模式

使用 MySQL 数据库。请运行以下 SQL 语句来创建表格：

```sql
CREATE TABLE activities (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            csv_path VARCHAR(255),
                            template TEXT,
                            scheduled_time DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='活动表';

CREATE TABLE recipients (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            phone VARCHAR(20),
                            name VARCHAR(100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='收件人表';


CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id INT,
    phone VARCHAR(20),
    message TEXT,
    status VARCHAR(20) comment 'success,failed',
    send_time DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='消息表';
```

### 创建topic 
```text
docker exec -it kafka-kafka-1 /bin/bash
kafka-topics.sh --create --topic message_topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

```
### 查看创建topic
```text
kafka-topics.sh --list --bootstrap-server localhost:9092
```
