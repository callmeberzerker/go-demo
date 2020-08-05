import axios from "redaxios";
import type { Book, BookCreate } from "../@types";

const BASE_URL = "http://localhost:8080/api/books";

const client = axios.create({});

const getAllBooks = async (): Promise<Book[]> => {
  const response = await client.get(BASE_URL);
  return response.data;
};

/**
 * Updates a book.
 *
 * @param book the book to be updated.
 */
const updateBook = async (book: Book): Promise<Book> => {
  const response = await client.put(`${BASE_URL}/${book.id}`, book);
  return response.data;
};

/**
 * Creates a book.
 *
 * @param book the book to be created.
 */
const createBook = async (book: BookCreate): Promise<Book> => {
  const response = await client.post(BASE_URL, book);
  return response.data;
};

export const BookService = {
  getAllBooks,
  updateBook,
  createBook,
};
