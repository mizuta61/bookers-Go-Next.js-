import React, { useState, useEffect, FC } from "react";
import { DateTime } from "luxon";
import Link from "next/link";
import axios from "axios";
import Button from "@mui/material/Button";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import VisibilityIcon from "@mui/icons-material/Visibility";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import { Book } from "../types/Book";
import Box from "@mui/material/Box";
import Modal from "@mui/material/Modal";

const IndexTable: FC = () => {
  const API_URL = process.env.NEXT_PUBLIC_API_URL;
  const [books, setBooks] = useState<Book[]>([]);
  const [selectedBookId, setSelectedBookId] = useState<number | null>(null);
  const selectedBook = books.find((book) => book.id === selectedBookId);

  const deleteBook = async (id: number) => {
    await axios
      .delete(`${API_URL}/books/${id}`)
      .catch(() => console.log("削除失敗"));
    console.log("削除成功");
    setBooks((books) => books.filter((book: Book) => book.id !== id));
  };

  useEffect(() => {
    fetch(`${API_URL}/books`)
      .then((res) => res.json())
      .then((books) => setBooks(books));
  }, []);

  const handleShowDetails = (id?: number) => setSelectedBookId(id || null);

  const style = {
    position: "absolute" as "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 400,
    bgcolor: "lightblue",
    p: 4,
    borderRadius: "0.5em",
  };

  return (
    <>
      <TableContainer>
        <Table sx={{ maxWidth: 650 }} align="center">
          <TableHead>
            <TableRow>
              <TableCell>Title</TableCell>
              <TableCell>Body</TableCell>
              <TableCell colSpan={3}></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {books.map((book) => {
              return (
                <TableRow key={book.id}>
                  <TableCell>{book.title}</TableCell>
                  <TableCell>{book.body}</TableCell>
                  <TableCell>
                    <Button
                      variant="contained"
                      color="primary"
                      size="small"
                      startIcon={<VisibilityIcon />}
                      onClick={() => handleShowDetails(book.id)}
                    >
                      SHOW
                    </Button>
                  </TableCell>
                  <TableCell>
                    <Button
                      variant="contained"
                      color="success"
                      size="small"
                      startIcon={<EditIcon />}
                    >
                      <Link href={`/books/edit/${book.id}`}>
                        <Box component="a">EDIT</Box>
                      </Link>
                    </Button>
                  </TableCell>
                  <TableCell>
                    <Button
                      variant="contained"
                      color="error"
                      size="small"
                      startIcon={<DeleteIcon />}
                      onClick={() => deleteBook(book.id)}
                    >
                      DESTROY
                    </Button>
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>

      {selectedBook && (
        <Modal open>
          <Box sx={style}>
            <Button onClick={() => handleShowDetails()} variant="contained">
              Close ✖️
            </Button>
            <Box component="p">ID: {selectedBook.id}</Box>
            <Box component="p">Title: {selectedBook.title}</Box>
            <Box component="p">Body: {selectedBook.body}</Box>
            <Box>
              Created At:{" "}
              {DateTime.fromISO(selectedBook.created_at, {
                zone: "Asia/Tokyo",
              }).toLocaleString(DateTime.DATE_MED)}
            </Box>
          </Box>
        </Modal>
      )}
    </>
  );
};
export default IndexTable;
