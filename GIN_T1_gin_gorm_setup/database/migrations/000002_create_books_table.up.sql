CREATE TABLE books (
   id SERIAL PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   author VARCHAR(100) NOT NULL,
   isbn VARCHAR(13) UNIQUE NOT NULL,
   publication_date DATE,
   page_count INT,
   genre VARCHAR(50),
   description TEXT,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_books_modtime
    BEFORE UPDATE ON books
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();