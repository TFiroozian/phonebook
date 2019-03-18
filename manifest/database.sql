\c app;    
CREATE SCHEMA app;    
CREATE TABLE IF NOT EXISTS app.contact (    
  id BIGSERIAL PRIMARY KEY,    
  first_name VARCHAR(128) NOT NULL,    
  last_name VARCHAR(128) NOT NULL,    
  phone_number VARCHAR(11) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE
);

