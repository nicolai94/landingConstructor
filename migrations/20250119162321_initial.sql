-- +goose Up
-- Создаем таблицы без внешних ключей
CREATE TABLE pre_landings (
                              id UUID PRIMARY KEY NOT NULL,
                              pwa_id UUID NOT NULL,
                              design INTEGER,
                              name VARCHAR(255),
                              icon_store TEXT,
                              icon_app TEXT,
                              developer VARCHAR(255),
                              subtitle VARCHAR(255),
                              rating INTEGER,
                              number_of_reviews INTEGER,
                              created_at TIMESTAMP NOT NULL DEFAULT now(),
                              updated_at TIMESTAMP NOT NULL DEFAULT now(),
                              deleted_at TIMESTAMP
);

CREATE TABLE pwas (
                      id UUID PRIMARY KEY NOT NULL,
                      name VARCHAR(255) UNIQUE NOT NULL,
                      type INTEGER,
                      icon TEXT,
                      pre_landing_id UUID,
                      created_at TIMESTAMP NOT NULL DEFAULT now(),
                      updated_at TIMESTAMP NOT NULL DEFAULT now(),
                      deleted_at TIMESTAMP
);

-- Добавляем внешние ключи после создания таблиц
ALTER TABLE pre_landings
    ADD CONSTRAINT fk_pre_landings_pwa_id
        FOREIGN KEY (pwa_id) REFERENCES pwas(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE pwas
    ADD CONSTRAINT fk_pwas_pre_landing_id
        FOREIGN KEY (pre_landing_id) REFERENCES pre_landings(id) ON UPDATE CASCADE ON DELETE CASCADE;

-- +goose Down
-- Удаляем таблицы в обратном порядке
DROP TABLE pwas;
DROP TABLE pre_landings;