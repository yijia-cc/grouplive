-- +migrate Up
CREATE TABLE amenity_type (
    id VARCHAR(3) PRIMARY KEY,
    title VARCHAR(100),
    description TEXT,
    thumbnail_url VARCHAR(255)
);

CREATE TABLE amenity (
    id VARCHAR(4) PRIMARY KEY,
    name VARCHAR(100),
    type_id VARCHAR(3)
);

CREATE TABLE operational_hour (
    id VARCHAR(4) PRIMARY KEY,
    start DATETIME,
    end DATETIME,
    amenity_id VARCHAR(4)
);

CREATE TABLE week (
    id VARCHAR(6) PRIMARY KEY,
    week_start DATETIME
);

CREATE TABLE time_slot_type (
    id BIT(4) PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE time_slot (
    id VARCHAR(6) PRIMARY KEY,
    type_id BIT(4),
    start DATETIME,
    end DATETIME,
    week_id VARCHAR(6)
);

CREATE TABLE reservation (
    id VARCHAR(6) PRIMARY KEY,
    amenity_id VARCHAR(4),
    time_slot_id VARCHAR(6),
    owner_id VARCHAR(6)
);

ALTER TABLE amenity
    ADD CONSTRAINT fk_amenity_type_id
    FOREIGN KEY (type_id) REFERENCES amenity_type(id);
ALTER TABLE operational_hour
    ADD CONSTRAINT fk_operational_hour_amenity_id
    FOREIGN KEY (amenity_id) REFERENCES amenity(id);
ALTER TABLE time_slot
    ADD CONSTRAINT fk_time_slot_type_id
    FOREIGN KEY (type_id) REFERENCES time_slot_type(id),
    ADD CONSTRAINT fk_time_slot_week_id
    FOREIGN KEY (week_id) REFERENCES week(id);
ALTER TABLE reservation
    ADD CONSTRAINT fk_reservation_amenity_id
    FOREIGN KEY (amenity_id) REFERENCES amenity(id),
    ADD CONSTRAINT fk_reservation_time_slot_id
    FOREIGN KEY (time_slot_id) REFERENCES time_slot(id);

-- +migrate Down
ALTER TABLE amenity
DROP FOREIGN KEY fk_amenity_type_id;
ALTER TABLE operational_hour
DROP FOREIGN KEY fk_operational_hour_amenity_id;
ALTER TABLE time_slot
DROP FOREIGN KEY fk_time_slot_type_id,
    DROP FOREIGN KEY fk_time_slot_week_id;
ALTER TABLE reservation
DROP FOREIGN KEY fk_reservation_amenity_id,
    DROP FOREIGN KEY fk_reservation_time_slot_id;

DROP TABLE reservation;
DROP TABLE time_slot;
DROP TABLE time_slot_type;
DROP TABLE week;
DROP TABLE operational_hour;
DROP TABLE amenity;
DROP TABLE amenity_type;

