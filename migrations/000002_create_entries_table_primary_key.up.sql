--Filename: migrations/000002_create_entries_table_primary_key.sql

ALTER TABLE entries 
    id bigserial NOT NULL;