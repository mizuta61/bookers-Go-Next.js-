import type { AppProps } from "next/app";
import { SessionProvider } from "next-auth/react";
import { CssBaseline } from "@mui/material";

function MyApp({ Component, pageProps: { session, ...pageProps } }: AppProps) {
  return (
    <SessionProvider session={session}>
      <CssBaseline />
      <Component {...pageProps} />
    </SessionProvider>
  );
}

export default MyApp;
