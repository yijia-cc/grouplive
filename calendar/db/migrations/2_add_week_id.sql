-- +migrate Up
ALTER TABLE reservation
    ADD week_id VARCHAR(6) NOT NULL;

SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE reservation
    ADD CONSTRAINT fk_reservation_week_id
        FOREIGN KEY (week_id) REFERENCES week(id);


-- +migrate Down
ALTER TABLE reservation
    DROP FOREIGN KEY fk_reservation_week_id;
ALTER TABLE reservation
    DROP COLUMN week_id;
