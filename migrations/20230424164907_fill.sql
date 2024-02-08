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

INSERT INTO airline_provider (airline_id, provider_id) VALUES ('AA', 'FZ');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('AA', 'JB');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('AA', 'SJ');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('IF', 'SU');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('IF', 'S7');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('IF', 'FZ');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('IF', 'N4');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('IF', 'JB');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('IF', 'WZ');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('RS', 'SU');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('RS', 'S7');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('RS', 'KV');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('RS', 'U6');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('RS', 'UT');
INSERT INTO airline_provider (airline_id, provider_id) VALUES ('RS', '5N');