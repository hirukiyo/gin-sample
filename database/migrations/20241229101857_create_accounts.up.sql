CREATE TABLE accounts (
    id bigint UNSIGNED AUTO_INCREMENT COMMENT 'アカウントID',
    name varchar(255) NOT NULL DEFAULT '' COMMENT '名前',
    email varchar(255) NOT NULL DEFAULT '' COMMENT 'メールアドレス',
    password varchar(255) NOT NULL DEFAULT '' COMMENT 'パスワード',
    created_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=10000;
