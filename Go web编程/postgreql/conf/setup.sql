CREATE TABLE `USERS` (
    `id` serial primary key,
    `uuid` varchar(64) NOT NULL unique,
    `name` varchar(255) DEFAULT NULL,
    `email` varchar(255) DEFAULT NULL unique,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL
);

CREATE TABLE `sessions` (
    `id` serial primary key,
    `uuid` varchar(64) NOT NULL unique,
    `email` varchar(255) DEFAULT NULL unique,
    `user_id` integer references users(id),
    `created_at` timestamp NOT NULL
)

CREATE TABLE `threads` (
    `id` serial primary key,
    `uuid` varchar(64) NOT NULL unique,
    `topic` text,
    `user_id` integer references users(id),
    `created_at` timestamp NOT NULL
)

CREATE TABLE `posts` (
    `id` serial primary key,
    `uuid` varchar(64) NOT NULL unique,
    `body` text,
    `user_id` integer references users(id),
    `user_id` integer references threads(id),
    `created_at` timestamp NOT NULL
)
