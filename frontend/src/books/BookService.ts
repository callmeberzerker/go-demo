import axios from "redaxios";
import type { Book } from "../@types";

const BASE_URL = "http://localhost:8080/api/books";

const getAllBooks = async (): Promise<Book[]> => {
  const response = await axios.get(BASE_URL);
  return response.data;
};

export const BookService = {
  getAllBooks,
};
