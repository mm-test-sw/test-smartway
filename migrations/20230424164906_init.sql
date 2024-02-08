-- +goose Up

--
-- Структура таблицы `airlines`
--
CREATE TABLE airlines (
                           code varchar(2) PRIMARY KEY,
                           name text NULL
);

--
-- Структура таблицы `providers`
--
CREATE TABLE providers (
                            id varchar(2) PRIMARY KEY,
                            name text NULL
);

--
-- Структура таблицы `airlineProvider`
--
CREATE TABLE airline_provider (
                                   airline_id varchar(2) NOT NULL,
                                   provider_id varchar(2) NOT NULL
);

--
-- Структура таблицы `schemas`
--
CREATE TABLE schemas (
                          id SERIAL PRIMARY KEY,
                          name text NOT NULL
);

--
-- Структура таблицы `schemaProvider`
--
CREATE TABLE schema_provider (
                                  schema_id int NOT NULL,
                                  provider_id varchar(2) NOT NULL
);

--
-- Структура таблицы `accounts`
--
CREATE TABLE accounts (
                           id serial PRIMARY KEY,
                           schema_id int NOT NULL
);