CREATE DATABASE IF NOT EXISTS trayDB;

USE trayDB;

CREATE TABLE IF NOT EXISTS trayTable (Id int NOT NULL AUTO_INCREMENT, Timestmp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, User_id varchar(50), Is_logged_in binary, Missing_tray_title text, Added_tray_title text, PRIMARY KEY (id));