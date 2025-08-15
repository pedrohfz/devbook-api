    CREATE DATABASE IF NOT EXISTS devbook;
    USE devbook;

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


insert into usuarios(nome, nick, email, senha)
values
("Pedrinho", "pedrinho", "pedrinho@gmail.com", "$2a$10$yB0Ap1HM/Bl3RkrpzmncRe6MzeVKf5csYhFEe5b4VBum9vqI38I3K"),
("Julia", "julia", "julia@gmail.com", "$2a$10$2wBcmZZKdUU.ryhve04uMuKuUns.qrFb5nA0V13Pb4k8u4gpgmHTG"),
("Mock User", "mock", "mock@gmail.com", "$2a$10$mtxH87HOLtYcZkiSJFk7Re4PwJaOXiqBUuc9b.jmt.1.tHSWHq9ny");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);