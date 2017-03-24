-- rambler up

CREATE TABLE IF NOT EXISTS servers (
    `reference` VARCHAR(50)           NOT NULL,
    `dns`       VARCHAR(255)          NOT NULL DEFAULT ".kdata.fr",
    `type`      ENUM('calc','data')   NOT NULL,
    PRIMARY KEY (`reference`)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE servers;

-- rambler up

INSERT IGNORE INTO servers (`reference`,`type`)
VALUES  ("blink","calc"),
        ("freya","calc"),
        ("widow","calc"),
        ("whiplash","calc"),
        ("flash","calc"),
        ("chimera","calc"),
        
        ("dust","data"),
        ("magneto","data"),
        ("winter","data"),
        ("elixir","data"),
        ("fury","data"),
        ("ares","data"),
        ("angel","data"),
        ("daredevil","data"),
        ("brute","data");

-- rambler down

DELETE FROM services_types;