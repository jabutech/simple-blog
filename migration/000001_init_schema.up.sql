CREATE TABLE `users` (
  `id` varchar(255) PRIMARY KEY,
  `fullname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `isAdmin` tinyint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (current_timestamp),
  `updated_at` timestamp NOT NULL DEFAULT (current_timestamp)
);

CREATE TABLE `posts` (
  `id` varchar(255) PRIMARY KEY,
  `user_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (current_timestamp),
  `updated_at` timestamp NOT NULL DEFAULT (current_timestamp)
);

CREATE INDEX `users_index_0` ON `users` (`id`);

CREATE INDEX `users_index_1` ON `users` (`email`);

CREATE INDEX `users_index_2` ON `users` (`id`, `email`);

CREATE INDEX `posts_index_3` ON `posts` (`id`);

CREATE INDEX `posts_index_4` ON `posts` (`user_id`);

CREATE INDEX `posts_index_5` ON `posts` (`title`);

CREATE INDEX `posts_index_6` ON `posts` (`id`, `user_id`, `title`);

ALTER TABLE `posts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
