CREATE TABLE `tools`
(
    `tools_id`  bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '节点ID',
    `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户id',
    PRIMARY KEY (`tools_id`)
)USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='工具';



INSERT INTO `test_data_1`.`tools` (`tools_id`, `tenant_id`) VALUES (1, 1, 'ccc.test_db_1');
INSERT INTO `test_data_1`.`tools` (`tools_id`, `tenant_id`) VALUES (2, 2, 'vvv.ccc.test_db_1');
INSERT INTO `test_data_1`.`tools` (`tools_id`, `tenant_id`) VALUES (3, 1, 'tool2.ccc.test_db_1');
