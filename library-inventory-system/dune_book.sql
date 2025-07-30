BEGIN; 

INSERT INTO tag (name) VALUES 
    (/*1*/'sci-fi'), 
    (/*2*/'desert'), 
    (/*3*/'classic');

INSERT INTO author (name) VALUES 
  (/*1*/'Frank Hebert');

INSERT INTO book (title, publication_year) VALUES 
  (/*1*/'Dune', 1965);

INSERT INTO author_book (book_id, author_id) VALUES 
  (1, 1);

INSERT INTO book_tag (book_id, tag_id) VALUES 
  (1, 1),
  (1, 2),
  (1, 3);

INSERT INTO status (name) VALUES 
  (/*1*/'available'),
  (/*2*/'checked_out');

INSERT INTO branch (name) VALUES
  (/*1*/'Downtown'),
  (/*2*/'Uptown');

INSERT INTO book_copy (book_id, branch_id, status_id) VALUES 
  (1, 1, 1),
  (1, 2, 2);

COMMIT;
