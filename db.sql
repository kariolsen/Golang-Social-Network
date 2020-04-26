DROP USER 'netadmin_s96lu'@'%';
CREATE USER 'netadmin_s96lu'@'%' IDENTIFIED BY 'netadmin_s96lu';
GRANT ALL PRIVILEGES ON socialnet.* TO 'netadmin_s96lu'@'%';

DROP TABLE IF EXISTS `Users`;
CREATE TABLE `Users` (
    `user_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(32) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `avatar` VARCHAR(255) NOT NULL DEFAULT 'avatar.png',
    PRIMARY KEY(user_id),
    UNIQUE(username),
    UNIQUE(password),
    UNIQUE(email)
);

DROP TABLE IF EXISTS `Profile`;
CREATE TABLE `Profile`(
    `profile_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `allow_unfollowed_views` BOOLEAN DEFAULT true,
    `job` VARCHAR(32) NOT NULL,
    `quote` VARCHAR(255) NOT NULL DEFAULT 'Nice to meet you',
    `followers_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `following_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `posts_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `views` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(profile_id),
    UNIQUE(user_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Follow`;
CREATE TABLE `Follow` (
    `follow_by` INT UNSIGNED NOT NULL,
    `follow_to` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(follow_by, follow_to),
    FOREIGN KEY(follow_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(follow_to) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Blacklist`;
CREATE TABLE `Blacklist` (
    `black_by` INT UNSIGNED NOT NULL,
    `black_to` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(black_by, black_to),
    FOREIGN KEY(black_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(black_to) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Posts`;
CREATE TABLE `Posts` (
    `post_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `likes` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_by` INT UNSIGNED NOT NULL NOT NULL,
    `created_date` DATETIME NOT NULL,
    `allow_comments` BOOLEAN DEFAULT true,
    `comments_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `images_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `title` VARCHAR(255) NOT NULL,
    `content` TEXT NOT NULL,
    PRIMARY KEY(post_id),
    FOREIGN KEY(created_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Comments`;
CREATE TABLE `Comments`(
    `comment_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `content` TEXT NOT NULL,
    `likes` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(comment_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Likes`;
CREATE TABLE `Likes` (
    `post_id` INT UNSIGNED NOT NULL,
    `like_by` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, like_by),
    FOREIGN KEY(like_by) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Hashtags`;
CREATE TABLE `Hashtags`(
    `hashtag_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `hashtag_name` VARCHAR(255) NOT NULL,
    `followers_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `posts_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(hashtag_id),
    UNIQUE(hashtag_name)
);

DROP TABLE IF EXISTS `Images`;
CREATE TABLE `Images`(
    `image_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `image_size` INT UNSIGNED NOT NULL DEFAULT 0,
    `image_name` VARCHAR(255) NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(image_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Posts_Hashtags`;
CREATE TABLE `Posts_Hashtags` (
    `post_id` INT UNSIGNED NOT NULL,
    `hashtag_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, hashtag_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(hashtag_id) REFERENCES Hashtags(hashtag_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Users_Hashtags`;
CREATE TABLE `Users_Hashtags` (
    `user_id` INT UNSIGNED NOT NULL,
    `hashtag_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(user_id, hashtag_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(hashtag_id) REFERENCES Hashtags(hashtag_id) ON DELETE CASCADE ON UPDATE CASCADE
);

/*DROP TABLE IF EXISTS `Topics`;
CREATE TABLE `Topics` (
    `topic_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `topic_name` VARCHAR(255) NOT NULL,
    `followers_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `hashtags_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `posts_num` INT UNSIGNED NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(topic_id),
    UNIQUE(topic_name)
);

DROP TABLE IF EXISTS `Posts_Topics`;
CREATE TABLE `Posts_Topics` (
    `post_id` INT UNSIGNED NOT NULL,
    `topic_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, topic_id),
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Users_Topics`;
CREATE TABLE `Users_Topics` (
    `user_id` INT UNSIGNED NOT NULL,
    `topic_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(user_id, topic_id),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE IF EXISTS `Hashtags_Topics`;
CREATE TABLE `Hashtags_Topics` (
    `hashtag_id` INT UNSIGNED NOT NULL,
    `topic_id` INT UNSIGNED NOT NULL,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(hashtag_id, topic_id),
    FOREIGN KEY(hashtag_id) REFERENCES Hashtags(hashtag_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(topic_id) REFERENCES Topics(topic_id) ON DELETE CASCADE ON UPDATE CASCADE
);*/

DROP TABLE IF EXISTS `Mentions`;
CREATE TABLE `Mentions`(
    `user_id` INT UNSIGNED NOT NULL,
    `post_id` INT UNSIGNED NOT NULL,
    `viewed` BOOLEAN DEFAULT false,
    `type` INT NOT NULL DEFAULT 0,
    `created_date` DATETIME NOT NULL,
    PRIMARY KEY(post_id, user_id, type),
    FOREIGN KEY(user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(post_id) REFERENCES Posts(post_id) ON DELETE CASCADE ON UPDATE CASCADE
);

DELIMITER $$
DROP TRIGGER IF EXISTS `create_user_profile`;
CREATE TRIGGER `create_user_profile` AFTER INSERT ON `Users` FOR EACH ROW
BEGIN
    INSERT INTO Profile (user_id, created_date) VALUES(NEW.user_id, NOW());
END$$

DROP TRIGGER IF EXISTS `new_follows_time`;
CREATE TRIGGER `new_follows_time` BEFORE INSERT ON `Follow` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_follows`;
CREATE TRIGGER `new_follows` AFTER INSERT ON `Follow` FOR EACH ROW
BEGIN
    UPDATE Profile SET followers_num = followers_num + 1 WHERE user_id = NEW.follow_to;
    UPDATE Profile SET following_num = following_num + 1 WHERE user_id = NEW.follow_by;
END$$

DROP TRIGGER IF EXISTS `remove_follows`;
CREATE TRIGGER `remove_follows` AFTER DELETE ON `Follow` FOR EACH ROW
BEGIN   
    UPDATE Profile SET followers_num = followers_num - 1 WHERE user_id = OLD.follow_to;
    UPDATE Profile SET following_num = following_num - 1 WHERE user_id = OLD.follow_by;
END$$

DROP TRIGGER IF EXISTS `new_blacks_time`;
CREATE TRIGGER `new_blacks_time` BEFORE INSERT ON `Blacklist` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$
DELIMITER ;


DELIMITER $$
DROP TRIGGER IF EXISTS `new_likes_time`;
CREATE TRIGGER `new_likes_time` BEFORE INSERT ON `Likes` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_likes`;
CREATE TRIGGER `new_likes` AFTER INSERT ON `Likes` FOR EACH ROW
BEGIN
    UPDATE Posts SET likes = likes + 1 WHERE post_id = NEW.post_id;
END$$

DROP TRIGGER IF EXISTS `remove_likes`;
CREATE TRIGGER `remove_likes` AFTER DELETE ON `Likes` FOR EACH ROW
BEGIN
    UPDATE Posts SET likes = likes - 1 WHERE post_id = OLD.post_id;
END$$

DROP TRIGGER IF EXISTS `new_posts`;
CREATE TRIGGER `new_posts` BEFORE INSERT ON `Posts` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
    UPDATE Profile SET posts_num = posts_num + 1 WHERE user_id = NEW.created_by;
END$$

DROP TRIGGER IF EXISTS `remove_posts`;
CREATE TRIGGER `remove_posts` AFTER DELETE ON `Posts` FOR EACH ROW
BEGIN
    UPDATE Profile SET posts_num = posts_num - 1 WHERE user_id = OLD.created_by;
END$$

DROP TRIGGER IF EXISTS `new_comments`;
CREATE TRIGGER `new_comments` BEFORE INSERT ON `Comments` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_post_comments`;
CREATE TRIGGER `new_post_comments` AFTER INSERT ON `Comments` FOR EACH ROW
BEGIN
    UPDATE Posts SET comments_num = comments_num + 1 WHERE post_id = NEW.post_id;
END$$

DROP TRIGGER IF EXISTS `remove_post_comments`;
CREATE TRIGGER `remove_post_comments` AFTER DELETE ON `Comments` FOR EACH ROW
BEGIN
    UPDATE Posts SET comments_num = comments_num - 1 WHERE post_id = OLD.post_id;
END$$

DROP TRIGGER IF EXISTS `new_images`;
CREATE TRIGGER `new_images` BEFORE INSERT ON `Images` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$
DELIMITER ;


DELIMITER $$
DROP TRIGGER IF EXISTS `new_hashtags`;
CREATE TRIGGER `new_hashtags` BEFORE INSERT ON `Hashtags` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_hashtags_time`;
CREATE TRIGGER `new_users_hashtags_time` BEFORE INSERT ON `Users_Hashtags` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_hashtags`;
CREATE TRIGGER `new_users_hashtags` AFTER INSERT ON `Users_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET followers_num = followers_num + 1 WHERE hashtag_id = NEW.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `remove_users_hashtags`;
CREATE TRIGGER `remove_users_hashtags` AFTER DELETE ON `Users_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET followers_num = followers_num - 1 WHERE hashtag_id = OLD.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `new_posts_hashtags_time`;
CREATE TRIGGER `new_posts_hashtags_time` BEFORE INSERT ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_posts_hashtags`;
CREATE TRIGGER `new_posts_hashtags` AFTER INSERT ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num + 1 WHERE hashtag_id = NEW.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `remove_posts_hashtags`;
CREATE TRIGGER `remove_posts_hashtags` AFTER DELETE ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num - 1 WHERE hashtag_id = OLD.hashtag_id;
END$$
DELIMITER ;

DELIMITER $$
DROP TRIGGER IF EXISTS `new_mentions`;
CREATE TRIGGER `new_mentions` BEFORE INSERT ON `Mentions` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$
DELIMITER ;

DELIMITER $$
CREATE PROCEDURE Follow_Relations (IN targetID INT UNSIGNED, IN userID INT UNSIGNED, OUT followers_num INT, OUT followings_num INT, OUT following_bool BOOL) 
BEGIN 
    SELECT COUNT(*) INTO followings_num FROM Follow WHERE follow_by = targetID; 
    SELECT COUNT(*) INTO followers_num FROM Follow WHERE follow_to = targetID; 
    SELECT COUNT(*) INTO following_bool FROM Follow WHERE follow_by = userID AND follow_to = targetID; 
END$$
DELIMITER ;
/*DROP TRIGGER IF EXISTS `new_posts_hashtags`;
CREATE TRIGGER `new_posts_hashtags` AFTER INSERT ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num + 1 WHERE hashtag_id = NEW.hashtag_id;
END$$

DROP TRIGGER IF EXISTS `remove_posts_hashtags`;
CREATE TRIGGER `remove_posts_hashtags` AFTER DELETE ON `Posts_Hashtags` FOR EACH ROW
BEGIN
    UPDATE Hashtags SET posts_num = posts_num - 1 WHERE hashtag_id = OLD.hashtag_id;
END$$
DELIMITER ;*/

/*
DELIMITER $$
DROP TRIGGER IF EXISTS `new_topics`;
CREATE TRIGGER `new_topics` BEFORE INSERT ON `Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_topics_time`;
CREATE TRIGGER `new_users_topics_time` BEFORE INSERT ON `Users_Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_users_topics`;
CREATE TRIGGER `new_users_topics` AFTER INSERT ON `Users_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET followers_num = followers_num + 1 WHERE topic_id = NEW.topic_id;
END$$

DROP TRIGGER IF EXISTS `remove_users_topics`;
CREATE TRIGGER `remove_users_topics` AFTER DELETE ON `Users_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET followers_num = followers_num - 1 WHERE topic_id = OLD.topic_id;
END$$

DROP TRIGGER IF EXISTS `new_hashtags_topics_time`;
CREATE TRIGGER `new_hashtags_topics_time` BEFORE INSERT ON `Hashtags_Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_hashtags_topics`;
CREATE TRIGGER `new_hashtags_topics` AFTER INSERT ON `Hashtags_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET hashtags_num = hashtags_num + 1 WHERE topic_id = NEW.topic_id;
END$$

DROP TRIGGER IF EXISTS `remove_hashtags_topics`;
CREATE TRIGGER `remove_hashtags_topics` AFTER DELETE ON `Hashtags_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET hashtags_num = hashtags_num - 1 WHERE topic_id = OLD.topic_id;
END$$

DROP TRIGGER IF EXISTS `new_posts_topics_time`;
CREATE TRIGGER `new_posts_topics_time` BEFORE INSERT ON `Posts_Topics` FOR EACH ROW
BEGIN
    SET NEW.created_date = NOW();
END$$

DROP TRIGGER IF EXISTS `new_posts_topics`;
CREATE TRIGGER `new_posts_topics` AFTER INSERT ON `Posts_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET posts_num = posts_num + 1 WHERE topic_id = NEW.topic_id;
END$$

DROP TRIGGER IF EXISTS `remove_posts_topics`;
CREATE TRIGGER `remove_posts_topics` AFTER DELETE ON `Posts_Topics` FOR EACH ROW
BEGIN
    UPDATE Topics SET posts_num = posts_num - 1 WHERE topic_id = OLD.topic_id;
END$$
DELIMITER ;
*/

/*
INSERT INTO `Topics` (`topic_name`) VALUES
('Shoes'),
('Bags'),
('Coats'),
('Computers'),
('Wall Street'),
('Cars');
*/

/*Passwords are encoded by API, but password is the same as username*/
INSERT INTO `Users` (`username`, `password`, `email`) VALUES
('sijia', '$2a$10$eXcpaPHH6tYg8Ie.bhvuZ.PSIykhBdIVpts0BnL0cXl/b3F9XyOKa', 'sijia@gmail.com'),
('takkar', '$2a$10$ttnsVDOPgMlA5vvDE33eneqVO3BHE/zif/axxI5AwNpOuRetkxFk6', 'takkar@gmail.com'),
('faiyaz', '$2a$10$.Wx2jBjYPiMFgWGCW.USze.qFMwrgN1TWOf50CQgqHDBzpcYV2uSa', 'faiyaz@gmail.com'),
('ghalib', '$2a$10$ziw6cqTgpSBIvASZOjTheey8sQYf1iW3HW4N.Xjq4GX6faKqzIrE.', 'ghalib@gmail.com'),
('user1', '$2a$10$PrsYgrp62NJkjOy1FOrL9uaRpzSfuiMv3oL6Xj5Hl90ZQtTlRmfZq', 'user1@gmail.com'),
('user2', '$2a$10$judwDLrzupULLqb8gxRQveGx.knR3LP2qJ/zaPH8YmYzoEdkr.tue', 'user2@gmail.com'),
('user3', '$2a$10$fIt2Gsfntg..wRPgY11yTugAt3HEeJPsbajftyFT4mRKIkJyjuBtS', 'user3@gmail.com'),
('hero1', '$2a$10$vOqEDCYP2ji9MZEp0lg.Jei.uOijw6viV4T5hbmt8/S3.Wi3WpOXS', 'hero1@gmail.com'),
('hero2', '$2a$10$C3XrawSnJIm74IhaVJ7m6upxl8ZWHKp6p.1GtPy6PTV9gMl0qAdr6', 'hero2@gmail.com'),
('hero3', '$2a$10$kLG3iRB1ULBTK.Jnhk.R0.LHuV6sXK1Djcs7X4xI7L2Ap8k9YYMXS', 'hero3@gmail.com'),
('nature', '$2a$10$nBi64BlbJMlzuSJfOhPlXevwdCgHOXKLZQUbJQ1q2Y7Ltbpaf1Woa', 'nature@gmail.com');


INSERT INTO `Follow` (`follow_by`, `follow_to`) VALUES
(2, 1),
(3, 1),
(4, 1),
(5, 1),
(6, 1),
(7, 1),
(8, 1),
(9, 1),
(1, 2),
(1, 3),
(1, 4),
(3, 2),
(4, 2),
(3, 6),
(5, 3),
(7, 6),
(5, 7),
(7, 5),
(2, 3);


INSERT INTO `Blacklist` (`black_by`, `black_to`) VALUES
(1, 2),
(5, 2);


INSERT INTO `Posts` (`title`, `content`, `created_by`, `images_num`) VALUES
('Welcome', '#Welcome# Welcome to the community, guys', 1, 2),
('my title..', 'my content...', 1, 1),
('first,', 'first_content', 2, 1),
('second', 'second_content', 2, 1),
('third', 'third content..', 2, 1),
('Awesome platform', '#Welcome# I love this platform', 5, 2),
('FirstPost', '#FirstPost# This is my first post !', 4, 2),
('ghalib''s first title..', 'and this is content!!!', 3, 1),
('Wow', '#Welcome# It has been a month now, still loving it', 1, 1),
('Number 8', 'Number 8', 8, 1),
('ID 4', 'ID 4', 4, 1),
('Hey guys', '#FirstPost# First day here', 7, 1),
('Number 9', '#Number9# I am number 9', 9, 1),
('Good game', 'Good game', 8, 1),
('Bad game', 'Bad game', 4, 1),
('Soso game', '#FirstPost# Soso game', 7, 1),
('I love this game', '#Number9# I love this game', 9, 1),
('Hello', 'World!!', 6, 1),
('Good day', 'Good day', 6, 1),
('Bad day', 'Bad day', 10, 1),
('Soso day', '#FirstPost# Soso day', 11, 1),
('Resident Evil 3 Released', '#RE3# #RE2# As an RE fan, @takkar I pre-purchased it. Gonna love it', 1, 9);


INSERT INTO `Hashtags` (`hashtag_name`) VALUES
('Welcome'),
('FirstPost'),
('Number9'),
('RE3'),
('RE2');


INSERT INTO `Posts_Hashtags` (`post_id`, `hashtag_id`) VALUES
(1,1),
(6,1),
(7,2),
(9,1),
(12,2),
(13,3),
(16,2),
(17,3),
(21,2),
(22,4),
(22,5);


INSERT INTO `Users_Hashtags` (`user_id`, `hashtag_id`) VALUES
(1,4),
(1,5),
(2,4),
(2,5),
(3,4),
(3,5),
(4,4),
(4,5),
(5,4),
(5,5),
(9,4),
(10,5),
(6,1),
(8,3),
(8,1),
(7,2),
(7,5),
(6,3),
(1,2),
(1,3);


INSERT INTO `Images` (`post_id`, `user_id`, `image_name`, `image_size`) VALUES
(1, 1, 'image1.jpeg', 11111),
(1, 1, 'image2.jpg', 11112),
(2, 1, 'image3.jpg', 11113),
(3, 2, 'image4.jpg', 11114),
(4, 2, 'image5.jpg', 11115),
(5, 2, 'image6.jpg', 11120),
(6, 5, 'image7.jpg', 11121),
(6, 5, 'image8.jpeg', 11122),
(7, 4, 'image9.jpg', 11123),
(7, 4, 'image10.jpg', 11124),
(8, 3, 'image11.png', 11125),
(9, 1, 'image12.jpg', 11130),
(10, 8, 'image13.jpg', 11131),
(11, 4, 'image14.png', 11132),
(12, 7, 'image15.jpg', 11133),
(13, 9, 'image16.png', 11134),
(14, 8, 'image17.jpg', 11135),
(15, 4, 'image18.jpg', 11140),
(16, 7, 'image19.jpg', 111141),
(17, 9, 'image20.jpg', 11142),
(18, 6, 'image21.jpg', 11143),
(19, 6, 'image22.jpg', 11144),
(20, 10, 'image23.jpg', 11145),
(21, 11, 'image24.jpg', 11150),
(22, 1, 'res1.jpeg', 14000),
(22, 1, 'res2.jpg', 15000),
(22, 1, 'res3.jpg', 16000),
(22, 1, 'res4.jpg', 17000),
(22, 1, 'res5.jpg', 18000),
(22, 1, 'res6.jpg', 19000),
(22, 1, 'res8.jpg', 20000),
(22, 1, 'res9.jpg', 21000),
(22, 1, 'res10.jpg', 22000);


INSERT INTO `Likes` (`post_id`, `like_by`) VALUES
(1, 1),
(1, 2),
(1, 3),
(2, 4),
(2, 5),
(2, 6),
(3, 7),
(3, 8),
(3, 9),
(4, 10),
(4, 1),
(4, 2),
(5, 3),
(5, 4),
(5, 5),
(6, 6),
(6, 7),
(6, 8),
(7, 9),
(7, 10),
(7, 1),
(8, 1),
(9, 1),
(3, 1),
(9, 2),
(10, 3),
(11, 4),
(12, 5),
(13, 6),
(14, 7),
(15, 8),
(15, 9),
(15, 10),
(16, 1),
(16, 2),
(16, 3),
(17, 4),
(18, 5),
(19, 6),
(20, 7),
(20, 8),
(20, 9),
(21, 10),
(22, 1),
(22, 2),
(22, 3),
(22, 4),
(22, 5),
(22, 6),
(22, 7),
(22, 8),
(22, 9),
(22, 10);

INSERT INTO `Comments` (`post_id`, `user_id`, `content`) VALUES 
(22, 1, "Love it"),
(22, 2, "Love RE"),
(22, 3, "Capcom is doint good"),
(22, 4, "Awesome"),
(22, 5, "Cool"),
(1, 3, "Sounds good"),
(2, 4, "Nice"),
(3, 5, "Like it"),
(4, 6, "Wow"),
(5, 7, "Super cool"),
(6, 1, "Enjoy it"),
(7, 8, "Look good"),
(8, 9, "HAHA"),
(1, 1, "Really?"),
(9, 3, "Not bad"),
(1, 6, "Enjoy it"),
(10, 8, "Look good"),
(11, 9, "HAHA"),
(12, 3, "Really?"),
(13, 5, "Not bad"),
(14, 10, "Thumbs up"),
(15, 9, "I am here"),
(16, 8, "Love it, too"),
(17, 7, "Playing it now"),
(18, 4, "Good game, bro");

INSERT INTO `Mentions` (`post_id`, `user_id`, `type`) VALUES
(22, 1, 0),
(1, 1, 1),
(1, 2, 1),
(1, 3, 1),
(2, 4, 1),
(2, 5, 1),
(2, 6, 1),
(3, 7, 1),
(3, 8, 1),
(3, 9, 1),
(4, 10, 1),
(4, 1, 1),
(4, 2, 1),
(5, 3, 1),
(5, 4, 1),
(5, 5, 1),
(6, 6, 1),
(6, 7, 1),
(6, 8, 1),
(7, 9, 1),
(7, 10, 1),
(7, 1, 1),
(8, 1, 1),
(9, 1, 1),
(3, 1, 1),
(9, 2, 1),
(10, 3, 1),
(11, 4, 1),
(12, 5, 1),
(13, 6, 1),
(14, 7, 1),
(15, 8, 1),
(15, 9, 1),
(15, 10, 1),
(16, 1, 1),
(16, 2, 1),
(16, 3, 1),
(17, 4, 1),
(18, 5, 1),
(19, 6, 1),
(20, 7, 1),
(20, 8, 1),
(20, 9, 1),
(21, 10, 1),
(22, 1, 1),
(22, 2, 1),
(22, 3, 1),
(22, 4, 1),
(22, 5, 1),
(22, 6, 1),
(22, 7, 1),
(22, 8, 1),
(22, 9, 1),
(22, 10, 1),
(22, 1, 2),
(22, 2, 2),
(22, 3, 2),
(22, 4, 2),
(22, 5, 2),
(1, 3, 2),
(2, 4, 2),
(3, 5, 2),
(4, 6, 2),
(5, 7, 2),
(6, 1, 2),
(7, 8, 2),
(8, 9, 2),
(1, 1, 2),
(9, 3, 2),
(1, 6, 2),
(10, 8, 2),
(11, 9, 2),
(12, 3, 2),
(13, 5, 2),
(14, 10, 2),
(15, 9, 2),
(16, 8, 2),
(17, 7, 2),
(18, 4, 2);

