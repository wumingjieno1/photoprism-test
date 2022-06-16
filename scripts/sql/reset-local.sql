DROP DATABASE IF EXISTS `local`;
CREATE DATABASE IF NOT EXISTS `local`;
CREATE USER IF NOT EXISTS 'local'@'%' IDENTIFIED BY 'local';
GRANT ALL PRIVILEGES ON `local`.* TO 'local'@'%';

FLUSH PRIVILEGES;
