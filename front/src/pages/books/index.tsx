import CreateForm from "../../components/CreateForm";
import IndexTable from "../../components/IndexTable";
import Typography from "@mui/material/Typography";
import { useSession, signIn, signOut } from "next-auth/react";
import Button from "@mui/material/Button";

const BookList = () => {
  const { data: session } = useSession();
  return (
    <>
      {session ? (
        <Typography
          component="h5"
          variant="h5"
          align="center"
          fontWeight="fontWeightBold"
          marginTop="30px"
        >
          Signed in as {session.user?.name} <br />
          <Button onClick={() => signOut()} variant="outlined" color="error">
            Sign out
          </Button>
        </Typography>
      ) : (
        <Typography
          component="h5"
          variant="h5"
          align="center"
          fontWeight="fontWeightBold"
          marginTop="30px"
        >
          Not signed in <br />
          <Button onClick={() => signIn()} variant="outlined" color="primary">
            Sign in
          </Button>
        </Typography>
      )}
      <Typography
        component="h4"
        variant="h4"
        align="center"
        marginTop="15px"
        fontWeight="fontWeightBold"
      >
        CreateBook
      </Typography>

      <CreateForm />

      <Typography
        component="h4"
        variant="h4"
        align="center"
        fontWeight="fontWeightBold"
      >
        BookList
      </Typography>
      {session && <IndexTable />}
    </>
  );
};
export default BookList;