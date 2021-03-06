CREATE DATABASE customers;
CREATE DATABASE products;
CREATE DATABASE orders;

\c customers;

CREATE TABLE public.customers (
    id serial NOT NULL,
    name varchar NOT NULL,
    email varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp,
    CONSTRAINT customers_pk PRIMARY KEY (id),
    CONSTRAINT customers_email_uq UNIQUE (email)
);

\c products;

CREATE TABLE public.products (
    id serial NOT NULL,
    title varchar NOT NULL,
    price integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp,
    CONSTRAINT products_pk PRIMARY KEY (id),
    CONSTRAINT products_title_uq UNIQUE (title)
);

\c orders;

CREATE TABLE public.orders (
    id serial NOT NULL,
    customer_id integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    deleted_at timestamp,
    CONSTRAINT orders_pk PRIMARY KEY (id)
);

CREATE TABLE public.order_items (
    id serial NOT NULL,
    order_id integer NOT NULL,
    product_id integer NOT NULL,
    qtd integer NOT NULL,
    CONSTRAINT order_items_pk PRIMARY KEY (id)
);

ALTER TABLE public.order_items ADD CONSTRAINT order_items_order_id_fk FOREIGN KEY (order_id)
REFERENCES public.orders (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;