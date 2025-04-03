create table alarm_config
(
    id              bigint auto_increment
        primary key,
    alarm_code      varchar(50)                          not null comment '告警场景编码',
    alarm_name      varchar(256)                         not null comment '告警类型名称',
    business_code   varchar(128)                         not null comment '业务场景编码',
    business_name   varchar(256)                         not null comment '业务场景名称',
    priority        varchar(50)                          not null comment '告警优先级',
    message_title   varchar(128)                         null comment '消息标题',
    message_content longtext                             null comment '消息内容',
    content_type    varchar(50)                          null comment '消息方式',
    description     varchar(256)                         not null comment '消息描述',
    status          tinyint(1) default 0                 not null comment '状态',
    is_delete       tinyint(1) default 1                 not null comment '删除状态',
    create_time     timestamp  default CURRENT_TIMESTAMP not null comment '创建时间时间戳',
    update_time     timestamp  default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间时间戳'
)
    comment '告警配置表';