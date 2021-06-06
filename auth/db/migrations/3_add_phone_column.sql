-- +migrate Up
ALTER TABLE user
    ADD phone VARCHAR(13);

-- +migrate Down
ALTER TABLE user
    DROP COLUMN phone;

