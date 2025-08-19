CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
);

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    primary key (usuario_id, seguidor_id)
);

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp
);

INSERT INTO usuarios(nome, nick, email, senha) VALUES
("Pedro", "pedro", "pedro@gmail.com", "$2a$10$yB0Ap1HM/Bl3RkrpzmncRe6MzeVKf5csYhFEe5b4VBum9vqI38I3K"),
("Marcelo", "marcelo", "marcelo@gmail.com", "$2a$10$2wBcmZZKdUU.ryhve04uMuKuUns.qrFb5nA0V13Pb4k8u4gpgmHTG"),
("Gabriel", "gabriel", "gabriel@gmail.com", "$2a$10$mtxH87HOLtYcZkiSJFk7Re4PwJaOXiqBUuc9b.jmt.1.tHSWHq9ny"),
("Matheus", "matheus", "matheus@gmail.com", "$2a$10$GsXIMenXrfDzqBWG/o5NDOlrxVBA7iBE2./VJLo38z1u5tp6Sk1uG");

INSERT INTO seguidores(usuario_id, seguidor_id) VALUES
(1, 2),
(3, 1),
(1, 4),
(2, 3),
(3, 2),
(2, 4),
(4, 3),
(4, 1);

INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES
("Primeira Publicação do Blog!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", 1),
("Segunda Publicação do Blog!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", 2),
("Terceira Publicação do Blog!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", 3),
("Quarta Publicação do Blog!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", 4);