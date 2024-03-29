create database IF NOT EXISTS parks;

DROP TABLE IF EXISTS parks;

create table parks (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(100) NOT NULL,
  location VARCHAR(100) NOT NULL,
  created_at datetime DEFAULT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  
  PRIMARY KEY (id)
);
