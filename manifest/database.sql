CREATE DATABASE app;

\c app;    
CREATE SCHEMA phone_book;    
CREATE TABLE IF NOT EXISTS phone_book.users (    
  id BIGSERIAL PRIMARY KEY,
  password VARCHAR(64), 
  username VARCHAR(128) NULL UNIQUE    
);

CREATE TABLE IF NOT EXISTS phone_book.contacts (    
  id BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL,
  first_name VARCHAR(128) NOT NULL,
  last_name VARCHAR(128) NOT NULL,   
  phone_number VARCHAR(11) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE,
  FOREIGN KEY (user_id) REFERENCES phone_book.users(id)
);


INSERT INTO phone_book.users (password, username) VALUES( 'Tara', '1234567890');
INSERT INTO phone_book.contacts (user_id, first_name, last_name, phone_number, email) VALUES(1, 'Tara', 'Firoozian', '12345678901', 'thr.firoozian@g,ail.com');
