# 活动管理系统

## 概述

该系统通过 Kafka 和 Go 语言实现了活动管理功能。包括生成和发送消息的生产者服务和处理消息的消费者服务。

## 数据库模式

使用 MySQL 数据库。请运行以下 SQL 语句来创建表格：

```sql
CREATE TABLE activities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    csv_path VARCHAR(255),
    template TEXT,
    scheduled_time DATETIME,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='活动表';

CREATE TABLE recipients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    phone VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='收件人表';

CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    activity_id INT,
    phone VARCHAR(20),
    message TEXT,
    status VARCHAR(20) comment 'success,failed',
    send_time DATETIME,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='消息表';
