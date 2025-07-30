const db = connect('mongodb://root:root@localhost/test?authSource=admin');

db.book.insertOne({
  title: 'Dune',
  publicationYear: 1965,
  tags: ['sci-fi', 'desert', 'classic'],
  authors: ['Frank Hebert'],
  copies: [{
    status: 'available',
    branch: 'Downtown'
  },
  {
    status: 'checked_out',
    branch: 'Uptown'
  }]
});
