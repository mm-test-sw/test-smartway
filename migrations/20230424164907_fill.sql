-- +goose Up

--
-- Заполнение таблицы `providers`
--

INSERT INTO providers (id, name) VALUES ('AA', 'AmericanAir');
INSERT INTO providers (id, name) VALUES ('IF', 'InternationFlights');
INSERT INTO providers (id, name) VALUES ('RS', 'RedStar');

--
-- Заполнение таблицы `airlines`
--

INSERT INTO airlines (code, name) VALUES ('SU', 'Аэрофлот');
INSERT INTO airlines (code, name) VALUES ('S7', 'S7');
INSERT INTO airlines (code, name) VALUES ('KV', 'КрасАвиа');
INSERT INTO airlines (code, name) VALUES ('U6', 'Уральские авиалинии');
INSERT INTO airlines (code, name) VALUES ('UT', 'ЮТэйр');
INSERT INTO airlines (code, name) VALUES ('FZ', 'Flydubai');
INSERT INTO airlines (code, name) VALUES ('JB', 'JetBlue');
INSERT INTO airlines (code, name) VALUES ('SJ', 'SuperJet');
INSERT INTO airlines (code, name) VALUES ('WZ', 'Wizz Air');
INSERT INTO airlines (code, name) VALUES ('N4', 'Nordwind Airlines');
INSERT INTO airlines (code, name) VALUES ('5N', 'SmartAvia');

--
-- Заполнение таблицы `schemas`
--

INSERT INTO schemas (id, name) VALUES (1, 'Основная');
INSERT INTO schemas (id, name) VALUES (2, 'Тестовая');

--
-- Заполнение таблицы `accounts`
--

INSERT INTO accounts (id, schema_id) VALUES (1, 2);
INSERT INTO accounts (id, schema_id) VALUES (2, 2);
INSERT INTO accounts (id, schema_id) VALUES (3, 1);

--
-- Заполнение таблицы `schemaProvider`
--

INSERT INTO schema_provider (schema_id, provider_id) VALUES (1, 'AA');
INSERT INTO schema_provider (schema_id, provider_id) VALUES (1, 'IF');
INSERT INTO schema_provider (schema_id, provider_id) VALUES (1, 'RS');
INSERT INTO schema_provider (schema_id, provider_id) VALUES (2, 'IF');
INSERT INTO schema_provider (schema_id, provider_id) VALUES (2, 'RS');

--
-- Заполнение таблицы `airlineProvider`
--

INSERT INTO airline_provider (airline_id, provider_id) VALUES ('FZ', 'AA');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('JB', 'AA');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('SJ', 'AA');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('SU', 'IF');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('S7', 'IF');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('FZ', 'IF');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('N4', 'IF');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('JB', 'IF');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('WZ', 'IF');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('SU', 'RS');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('S7', 'RS');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('KV', 'RS');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('U6', 'RS');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('UT', 'RS');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('5N', 'RS');