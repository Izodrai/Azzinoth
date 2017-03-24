-- rambler up

CREATE TABLE IF NOT EXISTS regular_need (
    `regular_id`    INT     NOT NULL,
    `need_id`       INT     NOT NULL,
    
    PRIMARY KEY (regular_id, need_id),
    CONSTRAINT regular_need_fk_regular_id FOREIGN KEY regular_id(regular_id) REFERENCES regular_timetable(id),
    CONSTRAINT regular_need_fk_need_id FOREIGN KEY need_id(need_id) REFERENCES regular_timetable(id)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE regular_need;