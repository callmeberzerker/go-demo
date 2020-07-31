import axios from "redaxios";
import type { Author } from "../@types";

const BASE_URL = "http://localhost:8080/api/authors";

const client = axios.create({});

/**
 * Gets all authors.
 */
const getAllAuthors = async (): Promise<Author[]> => {
  const response = await client.get(BASE_URL);
  return response.data;
};

/**
 * Updates author.
 *
 * @param author the author to be updated.
 */
const updateAuthor = async (author: Author): Promise<Author> => {
  const response = await client.put(`${BASE_URL}/${author.id}`, author);
  return response.data;
};

/**
 * Creates author.
 *
 * @param author the author to be created.
 */
const createAuthor = async (author: Author): Promise<Author> => {
  const response = await client.post(BASE_URL, author);
  return response.data;
};

export const AuthorService = {
  getAllAuthors,
  updateAuthor,
  createAuthor,
};
