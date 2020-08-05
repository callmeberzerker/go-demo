export interface Book {
  id: number | string;
  isbn: string;
  title: string;
  authorId: number | string;
}

export type BookCreate = Omit<Book, "id">;

export interface Author {
  id: number | string;
  firstName: string;
  lastName: string;
}
