
CREATE SCHEMA sample_schema;

SET search_path=sample_schema;

CREATE TABLE users 
(
  uid SERIAL NOT NULL,
  user_open_id VARCHAR(32),
  name VARCHAR(255),
  mail VARCHAR(255),
  PRIMARY KEY (uid)
);