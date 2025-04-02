create table alarm_config
(
    id              bigint auto_increment
        primary key,
    alarm_code      varchar(50)                          not null comment '告警场景code，举例：POValid',
    alarm_name      varchar(256)                         not null comment '告警类型名称，举例：PO校验失败',
    business_code   varchar(128)                         not null comment '业务场景Code，举例：PO',
    business_name   varchar(256)                         not null comment '业务描述，举例：采购单',
    priority        varchar(50)                          not null comment '告警优先级：High、Middle、Low',
    message_title   varchar(128)                         null comment '告警消息标题（如果alarm_record传值了，则以alarm_record为准），需要支持传参business_id',
    message_content longtext                             null comment '消息内容（如果alarm_record传值了，则以alarm_record为准），需要支持传参business_id',
    content_type    varchar(50)                          null comment '默认 text_plain、html',
    description     varchar(256)                         not null comment '告警类型描述',
    status          tinyint(1) default 0                 not null comment '状态 0启用 1未启用',
    is_delete       tinyint(1) default 1                 not null comment '删除状态 0:已删除 1:未删除',
    create_time     timestamp  default CURRENT_TIMESTAMP not null comment '创建时间时间戳',
    update_time     timestamp  default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间时间戳',
    rsvst1          varchar(50)                          null comment '备用字段',
    rsvst2          varchar(50)                          null comment '备用字段',
    rsvst3          varchar(100)                         null comment '备用字段',
    rsvst4          varchar(100)                         null comment '备用字段',
    rsvst5          varchar(50)                          null comment '备用字段',
    rsvdc1          int                                  null comment '备用字段',
    rsvdc2          int                                  null comment '备用字段',
    rsvdc3          decimal(14, 2)                       null comment '备用字段',
    rsvdc4          decimal(14, 2)                       null comment '备用字段',
    rsvdc5          decimal(14, 2)                       null comment '备用字段'
)
    comment '告警配置表';