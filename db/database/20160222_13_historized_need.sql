-- rambler up

CREATE TABLE IF NOT EXISTS historized_need (
    `historized_id`  INT     NOT NULL,
    `need_id`       INT     NOT NULL,
    
    PRIMARY KEY (historized_id, need_id),
    CONSTRAINT historized_need_fk_historized_id FOREIGN KEY historized_id(historized_id) REFERENCES historized_timetable(id),
    CONSTRAINT historized_need_fk_need_id FOREIGN KEY need_id(need_id) REFERENCES historized_timetable(id)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE historized_need;