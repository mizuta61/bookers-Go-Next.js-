import { useSession, signIn, signOut } from "next-auth/react";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Link from "next/link";
import Box from "@mui/material/Box";
import { NextPage } from "next";

const Component: NextPage = () => {
  const { data: session } = useSession();
  if (session) {
    return (
      <>
        <Typography
          component="h4"
          variant="h4"
          align="center"
          fontWeight="fontWeightBold"
          marginTop="30px"
        >
          Signed in as {session.user?.name} <br />
          <Button onClick={() => signOut()} variant="outlined" color="error">
            Sign out
          </Button>
        </Typography>

        <Link href="/books">
          <Box component="a">
            <Typography
              component="h4"
              variant="h4"
              align="center"
              marginTop="30px"
              fontWeight="fontWeightBold"
            >
              Click Here
            </Typography>
          </Box>
        </Link>
      </>
    );
  }
  return (
    <>
      <Typography
        component="h4"
        variant="h4"
        align="center"
        fontWeight="fontWeightBold"
        marginTop="30px"
      >
        Not signed in <br />
        <Button onClick={() => signIn()} variant="outlined" color="primary">
          Sign in
        </Button>
      </Typography>
    </>
  );
};
export default Component;
