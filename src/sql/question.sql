DROP TABLE IF EXISTS questions;
CREATE TABLE questions (
  questionid             INT AUTO_INCREMENT NOT NULL,
  question       VARCHAR(512) NOT NULL,
  difficulty     VARCHAR(128) NOT NULL,
  theme          VARCHAR(128) NOT NULL,
  remanence      INT, 
  PRIMARY KEY (`questionid`)
);

INSERT INTO questions
  (question, difficulty, theme, remanence)
VALUES
  ('Qui a peint la Joconde, un célèbre tableau exposé au Louvre à Paris ?', 'Facile', 'Art et Culture', 1),
  ('Quel est le plus grand désert chaud du monde ?', 'Moyen', 'Géographie', 1),
  ('Qui a écrit le roman "Cent ans de solitude", considéré comme une des œuvres majeures de la littérature du XXe siècle ?', 'Difficile', 'Littérature', 1);