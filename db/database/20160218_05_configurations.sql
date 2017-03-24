-- rambler up

CREATE TABLE IF NOT EXISTS configurations (
    `reference`             VARCHAR(50)     NOT NULL,
    `client_reference`      VARCHAR(50)     NOT NULL,
    `path`                  VARCHAR(255)    NOT NULL,
    `stream_name`           VARCHAR(255)    DEFAULT NULL,
    
    PRIMARY KEY (`reference`),
    CONSTRAINT configurations_fk_client_reference FOREIGN KEY client_reference(client_reference) REFERENCES clients(reference)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE configurations;