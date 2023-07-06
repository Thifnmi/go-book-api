CREATE TABLE books (
    id INT NOT NULL AUTO_INCREMENT,
    uuid varchar(36) not null,
    name varchar(255),
    category_id varchar(36),
    price int,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    INDEX logs_id (uuid),
    PRIMARY KEY(id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;