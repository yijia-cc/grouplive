-- +migrate Up
CREATE TABLE resource_type
(
    id   VARCHAR(3) PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE resource
(
    id      VARCHAR(3) PRIMARY KEY,
    name    VARCHAR(100),
    type VARCHAR(3)
);

ALTER TABLE resource
    ADD CONSTRAINT fk_resource_type_id
        FOREIGN KEY (type) REFERENCES resource_type(id);

ALTER TABLE permission_binding
    ADD resource_type VARCHAR(3) NOT NULL,

    DROP FOREIGN KEY fk_permission_binding_permission_id,
    DROP FOREIGN KEY fk_permission_binding_user_id,
    DROP PRIMARY KEY;

ALTER TABLE permission_binding
    ADD CONSTRAINT fk_permission_binding_permission_id
        FOREIGN KEY (permission_id) REFERENCES permission(id),
    ADD CONSTRAINT fk_permission_binding_user_id
        FOREIGN KEY (user_id) REFERENCES user(id),
    ADD PRIMARY KEY (permission_id, user_id, resource_type, resource_id),

    ADD CONSTRAINT fk_permission_binding_resource_type
        FOREIGN KEY (resource_type) REFERENCES resource(type),
    ADD CONSTRAINT fk_permission_binding_resource_id
        FOREIGN KEY (resource_id) REFERENCES resource(id);

-- +migrate Down
ALTER TABLE permission_binding
    DROP COLUMN resource_type,

    DROP FOREIGN KEY fk_permission_binding_permission_id,
    DROP FOREIGN KEY fk_permission_binding_user_id,
    DROP FOREIGN KEY fk_permission_binding_resource_type,
    DROP FOREIGN KEY fk_permission_binding_resource_id,
    DROP PRIMARY KEY;

ALTER TABLE permission_binding
    ADD PRIMARY KEY (permission_id, user_id, resource_id),
    ADD CONSTRAINT fk_permission_binding_permission_id
        FOREIGN KEY (permission_id) REFERENCES permission(id),
    ADD CONSTRAINT fk_permission_binding_user_id
        FOREIGN KEY (user_id) REFERENCES user(id);

ALTER TABLE resource
    DROP CONSTRAINT fk_resource_type_id;

DROP TABLE resource;
DROP TABLE resource_type;

