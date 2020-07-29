import axios from "redaxios";
import type { Book } from "../@types";

const BASE_URL = "http://localhost:8080/api/books";

const getAllBooks = (): Promise<Book[]> => {
  return axios.get(BASE_URL).then((response) => response.data);
};

export const BookService = {
  getAllBooks,
};
