import axios from "axios";
import React, { useState, FC } from "react";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import { Book, BookCreate } from "../types/Book";

const CreateForm: FC = () => {
  const API_URL = process.env.NEXT_PUBLIC_API_URL;
  const initialBook: BookCreate = { title: "", body: "" }
  const [book, setBook] = useState(initialBook);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setBook({ ...book, [e.target.name]: e.target.value });
  };

  const createBook = async () => {
    await axios.post(`${API_URL}/books`, book );
  };

  return (
    <Grid container>
      <Grid
        item
        xs
        display="flex"
        justifyContent="center"
        alignItems="center"
      >
        <Box component="form" sx={{ width: "30%" }} onSubmit={createBook}	>
          <TextField
            sx={{ marginBottom: "5px", width: "100%" }}
            type="text"
            value={book.title}
            name="title"
            onChange={handleChange}
            variant="standard"
            label="Title"
          />
          <TextField
            sx={{ marginBottom: "5px", width: "100%" }}
            type="text"
            value={book.body}
            name="body"
            onChange={handleChange}
            variant="standard"
            label="Body"
          />
          <Button
            sx={{ marginBottom: "15px", width: "100%" }}
            type="submit"
            variant="contained"
            color="secondary"
            startIcon={<MenuBookIcon />}
            formMethod="POST"
          >
            Create
          </Button>
        </Box>
      </Grid>
    </Grid>
  );
};
export default CreateForm;
