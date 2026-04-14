CREATE USER IF NOT EXISTS 'nassuser'@'localhost' IDENTIFIED BY 'nassuser';
GRANT ALL PRIVILEGES ON nass.* TO 'nassuser'@'localhost';
FLUSH PRIVILEGES;