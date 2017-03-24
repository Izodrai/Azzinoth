-- rambler up

UPDATE regular_days SET minimal_start_time = "00:00:00" WHERE minimal_start_time IS NULL;

-- rambler up

UPDATE regular_days SET maximal_start_time = "23:59:59" WHERE maximal_start_time IS NULL;

-- rambler up

ALTER TABLE regular_days MODIFY minimal_start_time TIME DEFAULT "00:00:00";

-- rambler down

ALTER TABLE regular_days MODIFY minimal_start_time TIME DEFAULT NULL,

-- rambler up

ALTER TABLE regular_days MODIFY maximal_start_time TIME DEFAULT "23:59:59";

-- rambler down

ALTER TABLE regular_days MODIFY maximal_start_time TIME DEFAULT NULL,
