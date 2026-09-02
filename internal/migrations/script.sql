CREATE TABLE categorias (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    descricao VARCHAR(150) NOT NULL UNIQUE,
    preco NUMERIC(10,2) NOT NULL,
    quantidade INTEGER NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    categoria_id INTEGER NOT NULL,
    CONSTRAINT fk_categoria
        FOREIGN KEY (categoria_id)
        REFERENCES categorias(id) ON UPDATE NO ACTION
        ON DELETE CASCADE
);