-- rambler up

CREATE TABLE IF NOT EXISTS status_types (
    `reference` VARCHAR(50)     NOT NULL,
    PRIMARY KEY (`reference`)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE status_types;

-- rambler up

INSERT IGNORE INTO status_types (`reference`)
VALUES  ("scheduled"),
        ("completed"),
        ("failed"),
        ("canceled"),
        ("running"),
        ("waiting");

-- rambler down

DELETE FROM status_types;