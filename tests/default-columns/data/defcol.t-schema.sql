CREATE TABLE t (
    pk INT PRIMARY KEY AUTO_INCREMENT,
    x INT NULL,
    y INT NOT NULL DEFAULT 123,
    z DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
