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
