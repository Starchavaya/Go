CREATE TABLE Mother (
    ID serial primary key ,
    Firstname  varchar(255),
    Lastname  varchar(255),
    Patronymic varchar(255)
);

CREATE TABLE Child (
    ID serial primary key ,
    Firstname  varchar(255),
    Lastname  varchar(255),
    Patronymic varchar(255),
    IDMother integer REFERENCES Mother (ID) ON DELETE CASCADE
);


INSERT INTO Mother(Firstname,Lastname,Patronymic) values('Starchavaya','Aliaksandra','Dmitrievna');
INSERT INTO Mother(Firstname,Lastname,Patronymic) values('Volochkova','Alina','Vladimirovna');
INSERT INTO Mother(Firstname,Lastname,Patronymic) values('Koshevaya','Zinaida','Ivanovna');

INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Starchavaya','Maria','Leonidovna',1);
INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Starchavaya','Sofia','Leonidovna',1);
INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Volochkova','Anna','Dmitrievna',2);
INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Koshevaya','Alesya','Dmitrievna',3);
INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Koshevaya','Diana','Dmitrievna',3);
INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Koshevaya','Violetta','Dmitrievna',3);
INSERT INTO Child(Firstname,Lastname,Patronymic,IDMother) values('Koshevaya','Karina','Dmitrievna',3);