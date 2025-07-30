./inspect_db.sh "SELECT b.* FROM BOOK b INNER JOIN book_tag bt ON bt.book_id = b.id INNER JOIN tag t ON t.id = bt.tag_id WHERE t.name = 'sci-fi'"
./inspect_db.sh "SELECT count(*) FROM book b INNER JOIN book_copy bc ON bc.book_id = b.id WHERE b.title = 'Dune'"
./inspect_db.sh "SELECT br.name, count(*) FROM book b INNER JOIN book_copy bc ON bc.book_id = b.id INNER JOIN branch br ON br.id = bc.branch_id GROUP BY br.name"
./inspect_db.sh "SELECT a.name, count(*) FROM author a INNER JOIN author_book ba ON a.id = ba.author_id GROUP BY ba.book_id, a.name"
