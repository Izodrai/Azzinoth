-- rambler up

CREATE TABLE IF NOT EXISTS scheduled_need (
    `scheduled_id`  INT     NOT NULL,
    `need_id`       INT     NOT NULL,
    
    PRIMARY KEY (scheduled_id, need_id),
    CONSTRAINT scheduled_need_fk_scheduled_id FOREIGN KEY scheduled_id(scheduled_id) REFERENCES scheduled_timetable(id),
    CONSTRAINT scheduled_need_fk_need_id FOREIGN KEY need_id(need_id) REFERENCES scheduled_timetable(id)
) ENGINE=innodb DEFAULT CHARSET=utf8;

-- rambler down

DROP TABLE scheduled_need;