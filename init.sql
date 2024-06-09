CREATE TABLE IF NOT EXISTS trending (
    id              SERIAL PRIMARY KEY,
    query           VARCHAR(255) NOT NULL,
    query_count     INT NOT NULL,
    service_area_id VARCHAR(2) NOT NULL,
    CONSTRAINT no_duplicate_query_in_same_location UNIQUE (query, service_area_id)
);

