CREATE database if not exists library;
use library;
CREATE TABLE if not exists book
(
  id              binary(16) NOT NULL,
  title           VARCHAR(150) NOT NULL,               
  author          VARCHAR(150) NOT NULL,               
  pages           INT NOT NULL,
  quantity        INT NOT NULL,  
  created_at      DATE NOT NULL,    
  updated_at      DATE,         
  PRIMARY KEY     (id)                            
);