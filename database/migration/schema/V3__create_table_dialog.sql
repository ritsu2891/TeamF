/**
 * @author Riku Nunokawa
 */

CREATE TABLE dialog(
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  content VARCHAR(255),
  title VARCHAR(255),
  author VARCHAR(255),
  source VARCHAR(255),
  link VARCHAR(255),
  style VARCHAR(255),
  user_id INT UNSIGNED NOT NULL,
  ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES user(`id`),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;