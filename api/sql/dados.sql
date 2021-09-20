insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$c/V67IlqPwnMt5MYmoQ.J.nTQi42rTKKUfOzZpbdBUcjssVtkicZi"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$c/V67IlqPwnMt5MYmoQ.J.nTQi42rTKKUfOzZpbdBUcjssVtkicZi"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$c/V67IlqPwnMt5MYmoQ.J.nTQi42rTKKUfOzZpbdBUcjssVtkicZi");


insert into seguidores(usuario_id, seguidor_id)
values
(1,2),
(3,1),
(1,3);


insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacao do usuario 1", "Essa é a publicaca do usuario 1", 1),
("Publicacao do usuario 2", "Essa é a publicaca do usuario 2", 2),
("Publicacao do usuario 3", "Essa é a publicaca do usuario 3", 3)