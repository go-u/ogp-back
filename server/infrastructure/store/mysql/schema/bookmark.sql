CREATE TABLE bookmark
(
    user_id    BIGINT UNSIGNED                  NOT NULL,
    fqdn       VARCHAR(100) CHARACTER SET ascii NOT NULL,
    created_at DATETIME                         NOT NULL,

    PRIMARY KEY (user_id, fqdn)
);
