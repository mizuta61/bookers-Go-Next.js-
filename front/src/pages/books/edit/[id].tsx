import Link from "next/link";
import { useRouter } from "next/router";
import UpdateForm from "../../../components/UpdateForm";
import React, { useEffect, useState } from "react";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import { NextPage } from "next";

const Edit: NextPage = () => {
  const router = useRouter();
  const bookId = router.query.id;
  const [book, setBook] = useState(null);
  const API_URL = process.env.NEXT_PUBLIC_API_URL;
  useEffect(() => {
    if (!bookId) return;
    fetch(`${API_URL}/books/${bookId}`)
      .then((res) => res.json())
      .then((book) => setBook(book));
  }, [bookId]);

  return (
    <>
      <Typography
        component="h4"
        variant="h4"
        align="center"
        margin="15px 0"
        fontWeight="fontWeightBold"
      >
        EditingBook
      </Typography>
      <Link href="/books">
        <Typography
          component="h5"
          variant="h5"
          align="center"
          marginBottom="15px"
          fontWeight="fontWeightBold"
        >
          <Box component="a">Back</Box>
        </Typography>
      </Link>
      {book && <UpdateForm book={book} />}
    </>
  );
};
export default Edit;
