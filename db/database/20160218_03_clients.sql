-- rambler up

CREATE TABLE IF NOT EXISTS clients (
    `reference` VARCHAR(50)     NOT NULL,
    `name`      VARCHAR(255)    NOT NULL,
    PRIMARY KEY (`reference`)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE clients;

-- rambler up

INSERT IGNORE INTO clients (`reference`,`name`)
VALUES  ("vp_fr","voyage prive fr"),
        ("vp_uk","voyage prive uk"),
        ("vp_es","voyage prive es"),
        ("vp_it","voyage prive it"),
        ("th","thalasseo"),
        ("pa","parc asterix"),
        ("odv","officiel des vacances"),
        ("clv","croquons la vie"),
        ("andre","andre"),
        ("vc","vacances campings");

-- rambler down

DELETE FROM clients;