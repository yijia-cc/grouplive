-- +migrate Up
CREATE TABLE user
(
    id                 VARCHAR(3) PRIMARY KEY,
    name               VARCHAR(100),
    username           VARCHAR(50),
    email              VARCHAR(200),
    encrypted_password VARCHAR(72),
    last_signed_in_at  DATETIME
);

CREATE TABLE permission
(
    id   VARCHAR(3) PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE permission_binding
(
    permission_id VARCHAR(3),
    user_id       VARCHAR(3),
    resource_id   VARCHAR(6)
);

ALTER TABLE permission_binding
    ADD CONSTRAINT fk_permission_binding_permission_id
        FOREIGN KEY (permission_id) REFERENCES permission (id),
    ADD CONSTRAINT fk_permission_binding_user_id
        FOREIGN KEY (user_id) REFERENCES user (id),
    ADD CONSTRAINT fk_permission_binding
        PRIMARY KEY (permission_id, user_id, resource_id);

-- +migrate Down
ALTER TABLE permission_binding
    DROP FOREIGN KEY fk_permission_binding_permission_id,
    DROP FOREIGN KEY fk_permission_binding_user_id;

DROP TABLE permission_binding;
DROP TABLE permission;
DROP TABLE user;

