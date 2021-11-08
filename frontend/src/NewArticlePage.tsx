import React from "react";
import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardHeader from "@mui/material/CardHeader";
import DeleteIcon from "@mui/icons-material/Delete";
import SendIcon from "@mui/icons-material/Send";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

const NewArticlePage: React.FC = () => {
  const navigate = useNavigate();
  const [submitLabel, setSubmitLabel] = React.useState("Publish");
  const [submitDisabled, setSubmitDisabled] = React.useState(false);
  const [submissionError, setSubmissionError] = React.useState<string>();
  const [title, setTitle] = React.useState("");
  const [author, setAuthor] = React.useState("");
  const [content, setContent] = React.useState("");
  const handleTitleChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };
  const handleAuthorChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAuthor(e.target.value);
  };
  const handleContentChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setContent(e.target.value);
  };
  const publishArticle = () => {
    setSubmitLabel("Publishing...");
    setSubmitDisabled(true);
    setSubmissionError("");
    axios
      .post("/api/articles", {
        title,
        content,
        author,
      })
      .then((response) => {
        navigate(`/article/${response.data.id}`);
      })
      .catch((err) => {
        if (err?.response?.data?.message) {
          setSubmissionError(err?.response?.data?.message);
        } else {
          setSubmissionError(err.message);
        }
      })
      .finally(() => {
        setSubmitDisabled(false);
        setSubmitLabel("Publish");
      });
  };

  return (
    <Box sx={{ my: 2, mx: 4 }}>
      <Card sx={{ my: 1 }}>
        <CardHeader
          title={
            <TextField
              label="Title"
              placeholder="Title"
              variant="standard"
              value={title}
              onChange={handleTitleChanged}
              required
              fullWidth
            />
          }
          subheader={
            <TextField
              label="Author"
              variant="standard"
              value={author}
              onChange={handleAuthorChanged}
              fullWidth
              required
              sx={{ mt: 2 }}
            />
          }
        />
        <CardContent>
          <TextField
            variant="outlined"
            placeholder="Write your article content here"
            minRows="12"
            value={content}
            onChange={handleContentChanged}
            fullWidth
            required
            multiline
          />
          <Stack spacing={2} direction="row" sx={{ mt: 4 }}>
            <Button
              variant="contained"
              endIcon={<SendIcon />}
              onClick={publishArticle}
              disabled={submitDisabled}
            >
              {submitLabel}
            </Button>
            <Button
              variant="outlined"
              startIcon={<DeleteIcon />}
              component={Link}
              to="/"
            >
              Discard
            </Button>
          </Stack>
          {submissionError ? (
            <Alert sx={{ mt: 4 }} severity="error">
              {submissionError}
            </Alert>
          ) : null}
        </CardContent>
      </Card>
    </Box>
  );
};

export default NewArticlePage;
