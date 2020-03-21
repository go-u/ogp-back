CREATE TABLE user
(
    id         serial,
    uid        VARCHAR(128) CHARACTER SET ascii  NOT NULL,
    name       VARCHAR(50) CHARACTER SET utf8mb4 NOT NULL,
    created_at DATETIME                          NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (uid)
)
