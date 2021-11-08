import React from "react";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { Link } from "react-router-dom";

const NotFoundPage: React.FC = () => {
  return (
    <Box sx={{ my: 4, mx: 4 }}>
      <Typography variant="h3">Page not found</Typography>
      <Button component={Link} to="/">
        Go to top
      </Button>
    </Box>
  );
};

export default NotFoundPage;
