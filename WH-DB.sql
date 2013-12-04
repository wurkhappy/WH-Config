DROP TABLE IF EXISTS wh_user;
DROP TABLE IF EXISTS agreement;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS status;
DROP TABLE IF EXISTS balanced_user;
DROP TABLE IF EXISTS transaction;
DROP TABLE IF EXISTS tag;

CREATE TABLE wh_user (
    id uuid PRIMARY KEY,
    password bytea,
    data json
);

CREATE TABLE agreement (
    id uuid PRIMARY KEY,
    data json
);

CREATE TABLE comment (
    id uuid PRIMARY KEY,
    data json
);

CREATE TABLE status (
    id uuid PRIMARY KEY,
    data json
);

CREATE TABLE balanced_user (
    id uuid PRIMARY KEY,
    data json
);

CREATE TABLE transaction (
    id uuid PRIMARY KEY,
    data json
);

CREATE TABLE tag (
    id uuid PRIMARY KEY,
    data json
);

CREATE OR REPLACE FUNCTION upsert_user(uuid, bytea, json) RETURNS VOID AS
$$
BEGIN
    LOOP
        -- first try to update the key
        UPDATE wh_user SET password = $2, data = $3 WHERE id = $1;
        IF found THEN
            RETURN;
        END IF;
        -- not there, so try to insert the key
        -- if someone else inserts the same key concurrently,
        -- we could get a unique-key failure
        BEGIN
            INSERT INTO wh_user(id, password, data) VALUES ($1, $2, $3);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- Do nothing, and loop to try the UPDATE again.
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION upsert_agreement(uuid, json) RETURNS VOID AS
$$
BEGIN
    LOOP
        -- first try to update the key
        UPDATE agreement SET data = $2 WHERE id = $1;
        IF found THEN
            RETURN;
        END IF;
        -- not there, so try to insert the key
        -- if someone else inserts the same key concurrently,
        -- we could get a unique-key failure
        BEGIN
            INSERT INTO agreement(id, data) VALUES ($1, $2);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- Do nothing, and loop to try the UPDATE again.
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION upsert_status(uuid, json) RETURNS VOID AS
$$
BEGIN
    LOOP
        -- first try to update the key
        UPDATE status SET data = $2 WHERE id = $1;
        IF found THEN
            RETURN;
        END IF;
        -- not there, so try to insert the key
        -- if someone else inserts the same key concurrently,
        -- we could get a unique-key failure
        BEGIN
            INSERT INTO status(id, data) VALUES ($1, $2);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- Do nothing, and loop to try the UPDATE again.
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION upsert_comment(uuid, json) RETURNS VOID AS
$$
BEGIN
    LOOP
        -- first try to update the key
        UPDATE comment SET data = $2 WHERE id = $1;
        IF found THEN
            RETURN;
        END IF;
        -- not there, so try to insert the key
        -- if someone else inserts the same key concurrently,
        -- we could get a unique-key failure
        BEGIN
            INSERT INTO comment(id, data) VALUES ($1, $2);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- Do nothing, and loop to try the UPDATE again.
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION upsert_balanceduser(uuid, json) RETURNS VOID AS
$$
BEGIN
    LOOP
        -- first try to update the key
        UPDATE balanced_user SET data = $2 WHERE id = $1;
        IF found THEN
            RETURN;
        END IF;
        -- not there, so try to insert the key
        -- if someone else inserts the same key concurrently,
        -- we could get a unique-key failure
        BEGIN
            INSERT INTO balanced_user(id, data) VALUES ($1, $2);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- Do nothing, and loop to try the UPDATE again.
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION upsert_transaction(uuid, json) RETURNS VOID AS
$$
BEGIN
    LOOP
        -- first try to update the key
        UPDATE transaction SET data = $2 WHERE id = $1;
        IF found THEN
            RETURN;
        END IF;
        -- not there, so try to insert the key
        -- if someone else inserts the same key concurrently,
        -- we could get a unique-key failure
        BEGIN
            INSERT INTO transaction(id, data) VALUES ($1, $2);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- Do nothing, and loop to try the UPDATE again.
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;

