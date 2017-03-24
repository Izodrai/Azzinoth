-- rambler up

CREATE TABLE IF NOT EXISTS regular_days (
    `regular_id`            INT NOT NULL,
    `day`                   ENUM('monday','tuesday','wednesday','thursday','friday','saturday','sunday')   NOT NULL,
    `minimal_start_time`    TIME DEFAULT NULL,
    `maximal_start_time`      TIME DEFAULT NULL,
    
    PRIMARY KEY (`regular_id`, `day`),
    CONSTRAINT regular_weeks_fk_regular_id FOREIGN KEY regular_id(regular_id) REFERENCES regular_timetable(id)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE regular_days;