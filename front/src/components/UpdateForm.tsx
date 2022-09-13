import axios from "axios";
import { FC, useState } from "react";
import UpdateIcon from "@mui/icons-material/Update";
import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import { TextField } from "@mui/material";
import Box from "@mui/material/Box";
import { Book, BookCreate } from "../types/Book";

type Props = {
  book: Book;
};
const UpdateForm: FC<Props> = ({ book }) => {
  const { title, body } = book;
  const [value, setValue] = useState<BookCreate>({ title, body });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setValue((value) => {
      return { ...value, [e.target.name]: e.target.value };
    });
  };
  const  API_URL = process.env.NEXT_PUBLIC_API_URL;
  const updateBook = async () => {
    await axios.patch(`${API_URL}/books/${book.id}`, {
      book: value,
    });
  };

  return (
    <Grid container>
      <Grid item xs display="flex" justifyContent="center" alignItems="center">
        <Box component="form" sx={{ width: "30%" }} onSubmit={updateBook}>
          <TextField
            sx={{ marginBottom: "5px", width: "100%" }}
            type="text"
            value={value.title}
            name="title"
            onChange={handleChange}
            variant="standard"
            label="TitleUpdate"
          />
          <br />
          <TextField
            sx={{ marginBottom: "5px", width: "100%" }}
            type="text"
            value={value.body}
            name="body"
            onChange={handleChange}
            variant="standard"
            label="BodyUpdate"
          />
          <br />
          <Button
            sx={{ marginBottom: "15px", width: "100%" }}
            type="submit"
            variant="contained"
            color="secondary"
            startIcon={<UpdateIcon />}
          >
            Update
          </Button>
        </Box>
      </Grid>
    </Grid>
  );
};
export default UpdateForm;