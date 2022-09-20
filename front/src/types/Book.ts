export type Book = {
  id: number;
  title: string;
  body: string;
  created_at: string;
  updated_at: string;
};

export type BookCreate = Pick<Book, "title" | "body">