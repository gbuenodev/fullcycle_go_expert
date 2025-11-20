import type { ReactNode } from "react";
import { Box, CircularProgress } from "@mui/material";

interface LoadingProps {
  loading: boolean;
  children: ReactNode;
}

export default function Loading({ loading, children }: LoadingProps) {
  if (loading) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          minHeight: "200px",
        }}
      >
        <CircularProgress />
      </Box>
    );
  }

  return <>{children}</>;
}
