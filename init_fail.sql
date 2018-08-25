CREATE TABLE IF NOT EXISTS region (
    id                   serial  NOT NULL,
	created_at           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ,
	updated_at           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ,
	CONSTRAINT region_id PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS shop (
    id                   serial  NOT NULL,
    -- region_id            bigint(20) unsigned NOT NULL,
    name                 varchar(255) NOT NULL,
	created_at           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ,
	updated_at           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ,
	CONSTRAINT shop_id PRIMARY KEY ( id ),
    -- CONSTRAINT shop_id_region_id_foreign FOREIGN KEY ( region_id )
    CONSTRAINT shop_id_region_id_foreign FOREIGN KEY ( id )
        REFERENCES region( id )
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS book (
    id                   serial  NOT NULL,
    shop_id              bigint(20) unsigned NOT NULL,
    name                 varchar(255) NOT NULL,
    price                int NOT NULL,
	created_at           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ,
	updated_at           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  ,
	CONSTRAINT book_id PRIMARY KEY ( id ),
    CONSTRAINT book_shop_id_shop_id_foreign FOREIGN KEY ( shop_id )
        REFERENCES shop( id )
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);