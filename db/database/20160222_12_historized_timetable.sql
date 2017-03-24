-- rambler up

CREATE TABLE IF NOT EXISTS historized_timetable (
    `id`                    INT         NOT NULL AUTO_INCREMENT,
    `regular_id`            INT         NOT NULL,
    `process_id`            INT,
    `scheduled_date`        DATE        NOT NULL,
    `minimal_start_time`    TIME,
    `start_time`            DATETIME    DEFAULT NULL,
    `maximal_start_time`    TIME,
    `end_time`              DATETIME    DEFAULT NULL,
    `status`                VARCHAR(50) DEFAULT "scheduled",
    `infos`                 TEXT,
    
    PRIMARY KEY (`id`),
    
    CONSTRAINT historized_timetable_fk_regular_id FOREIGN KEY regular_id(regular_id) REFERENCES regular_timetable(id),
    CONSTRAINT historized_timetable_fk_status_reference FOREIGN KEY status(status) REFERENCES status_types(reference)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE historized_timetable;