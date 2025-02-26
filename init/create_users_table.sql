DROP TABLE IF EXISTS `testDB`.`users`;

CREATE TABLE `testDB`.`users`(
  `ramen_id` INT(5) NOT NULL UNIQUE AUTO_INCREMENT,
  `store` VARCHAR(50) NOT NULL,
  `menu` VARCHAR(30) NOT NULL,
  `address` VARCHAR(100) DEFAULT "記載ナシ",
  `station` VARCHAR(20) DEFAULT "記載ナシ",
  `taste_id` INT(2) NOT NULL,
  `price` INT(5) NOT NULL DEFAULT 99999,
  `evaluation` FLOAT NOT NULL DEFAULT 99.9,
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ramen_id`)
);

INSERT INTO users (store, menu, address, station, taste_id, price, evaluation) VALUES ('野方ホープ  目黒店', '元', '東京都目黒区目黒1-5-16 ABC MEGURO 1F', '目黒駅', 3, 750, 9999);
-- INSERT INTO user values (id, 'お店', 'メニュー', '住所', '駅', '味の種類', '価格', '評価', '現在時刻', 'メニュー');
--  (store, menu, address, station, taste_id, price, evaluation)
-- ("", "", "", "", , , )