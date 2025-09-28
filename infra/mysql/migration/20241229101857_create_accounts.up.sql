CREATE TABLE accounts (
    id bigint UNSIGNED AUTO_INCREMENT COMMENT 'アカウントID',
    name varchar(255) NOT NULL DEFAULT '' COMMENT '名前',
    email varchar(512) NOT NULL DEFAULT '' COMMENT 'メールアドレス',
    password varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'パスワード',
    status integer NOT NULL DEFAULT 0 COMMENT 'ステータス(1: 有効 / 他: 無効)',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=10001 COMMENT='アカウント';

INSERT INTO accounts (name, email, password, status) VALUES ("Test User1", "test1@example.jp", "password1", 1);
INSERT INTO accounts (name, email, password, status) VALUES ("Test User2", "test2@example.jp", "password2", 1);
