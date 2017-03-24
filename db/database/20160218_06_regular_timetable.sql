-- rambler up

CREATE TABLE IF NOT EXISTS regular_timetable (
    `id`                    INT             NOT NULL AUTO_INCREMENT,
    `service_reference`     VARCHAR(50)     NOT NULL,
    `config_reference`      VARCHAR(50)     NOT NULL,
    `process_id`            INT             NOT NULL,
    
    PRIMARY KEY (`id`),
    CONSTRAINT regular_timetable_fk_service_reference FOREIGN KEY service_reference(service_reference) REFERENCES services(reference),
    CONSTRAINT regular_timetable_fk_config_reference FOREIGN KEY config_reference(config_reference) REFERENCES configurations(reference)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE regular_timetable;