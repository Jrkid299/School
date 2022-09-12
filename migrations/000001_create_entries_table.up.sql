--Filename: migrations/000001_create_entries_table.up.sql

CREATE TABLE IF NOT EXISTS entries (
    english word text NOT NULL,
    kriol word text NOT NULL,
);