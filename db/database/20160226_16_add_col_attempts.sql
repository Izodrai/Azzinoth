-- rambler up

ALTER TABLE regular_days ADD COLUMN `attemps` INT DEFAULT 5;

-- rambler down

ALTER TABLE regular_days DROP COLUMN `attemps`;
