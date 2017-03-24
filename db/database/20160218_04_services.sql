-- rambler up

CREATE TABLE IF NOT EXISTS services (
    `reference`             VARCHAR(50) NOT NULL,
    `service_type_reference`     VARCHAR(50) NOT NULL,
    `calc_server_reference` VARCHAR(50) NOT NULL,
    `last_activity`         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`reference`),
    CONSTRAINT services_fk_service_type_reference FOREIGN KEY service_type_reference(service_type_reference) REFERENCES services_types(reference),
    CONSTRAINT services_fk_calc_server_reference FOREIGN KEY calc_server_reference(calc_server_reference) REFERENCES servers(reference)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE services;