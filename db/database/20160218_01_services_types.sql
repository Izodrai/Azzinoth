-- rambler up

CREATE TABLE IF NOT EXISTS services_types (
    `reference` VARCHAR(50)     NOT NULL,
    `name`      VARCHAR(255)    NOT NULL,
    PRIMARY KEY (`reference`)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE services_types;

-- rambler up

INSERT IGNORE INTO services_types (`reference`,`name`)
VALUES  ("ftu","format-to-universal"),
        ("dmc","data-mapping-component"),
        ("optin","optin-maker"),
        ("specific","specific"),
        ("sampler","sampler"),
        ("grouper","grouper"),
        ("pression","pression"),
        ("broadcaster","broadcaster");

-- rambler down

DELETE FROM services_types;