CREATE TABLE IF NOT EXISTS users (
    id character varying(64) PRIMARY KEY NOT NULL,
    email character varying(120) NOT NULL,
    password_hash character varying(256) NOT NULL,
    "timestamp" timestamp(6) without time zone,
    username character varying(64)
);