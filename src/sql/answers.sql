DROP TABLE IF EXISTS answers;
CREATE TABLE answers (
  answerid              INT AUTO_INCREMENT NOT NULL,
  questionid            INT NOT NULL,
  question_answer       VARCHAR(512) NOT NULL,
  Valid                 VARCHAR(128) NOT NULL,
  PRIMARY KEY (`answerid`),
  FOREIGN KEY (`questionid`) REFERENCES questions(questionid)
);

INSERT INTO answers
  (questionid, question_answer, Valid)
VALUES
  (1, 'Léonard de Vinci', "Vrai"),
  (1, 'Vincent van Gogh', "Faux"),
  (1, 'Pablo Picasso', "Faux"),
  (1, 'Claude Monet', "Faux"),
  (2, 'Le Sahara', "Faux"),
  (2, 'Le désert de Atacama', "Faux"),
  (2, 'Le Dasht-e Lut', "Vrai"),
  (2, 'Le désert de Gobi', "Faux"),
  (3, 'Pablo Neruda', "Faux"),
  (3, 'Jorge Luis Borges', "Faux"),
  (3, 'Mario Vargas Llosa', "Faux"),
  (3, 'Gabriel García Márquez', "Vrai");
  

ALTER TABLE answers
ADD FOREIGN KEY (questionid) REFERENCES questions(questionid);