CREATE TABLE `user`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT "",
    `age` INT UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `user`(`name`,`age`) VALUES
("one1","18"),("one2","19"),("one3","20"),("one4","21");