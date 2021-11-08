import React from "react";
import Alert from "@mui/material/Alert";
import Avatar from "@mui/material/Avatar";
import Stack from "@mui/material/Stack";
import SendIcon from "@mui/icons-material/Send";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import TextField from "@mui/material/TextField";
import CardContent from "@mui/material/CardContent";
import CardHeader from "@mui/material/CardHeader";
import CircularProgress from "@mui/material/CircularProgress";
import Divider from "@mui/material/Divider";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import ListItemText from "@mui/material/ListItemText";
import PersonIcon from "@mui/icons-material/Person";
import Typography from "@mui/material/Typography";
import useAxios from "axios-hooks";
import { useParams, Link } from "react-router-dom";
import axios from "axios";

type ArticleCardProps = {
  article: {
    id: number;
    title: string;
    author: string;
    content: string;
    created_at: string;
  };
};

const ArticleCard: React.FC<ArticleCardProps> = ({ article }) => {
  return (
    <Card sx={{ my: 1 }}>
      <CardHeader
        avatar={
          <Avatar>
            <PersonIcon />
          </Avatar>
        }
        title={article.author}
        subheader={new Date(article.created_at).toDateString()}
      />
      <CardContent>
        <Typography variant="h3">{article.title}</Typography>
        <Typography
          sx={{ mt: 2.0, whiteSpace: "pre-wrap" }}
          color="text.secondary"
        >
          {article.content}
        </Typography>
      </CardContent>
    </Card>
  );
};

type Comment = {
  id: number;
  article_id: number;
  author: number;
  content: string;
  created_at: string;
};

const CommentsCard: React.FC = () => {
  const [submitLabel, setSubmitLabel] = React.useState("Submit");
  const [submitDisabled, setSubmitDisabled] = React.useState(false);
  const [submissionError, setSubmissionError] = React.useState<string>();
  const { article_id: articleId } = useParams();
  const [{ data, loading, error }, referch] = useAxios(
    `/api/articles/${articleId}/comments`
  );
  const [author, setAuthor] = React.useState("");
  const [content, setContent] = React.useState("");
  const handleAuthorChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAuthor(e.target.value);
  };
  const handleContentChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setContent(e.target.value);
  };
  const submitComment = () => {
    setSubmitLabel("Submitting...");
    setSubmitDisabled(true);
    setSubmissionError("");
    axios
      .post(`/api/articles/${articleId}/comments`, {
        author,
        content,
      })
      .then(() => {
        setAuthor("");
        setContent("");
        referch();
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

  if (loading) {
    return <CircularProgress />;
  }
  if (error) {
    return <p>{error}</p>;
  }
  return (
    <Card sx={{ my: 1 }}>
      <List sx={{ bgcolor: "background.paper" }}>
        <ListItem>
          <Stack direction="column" sx={{ mt: 4, width: "100%" }}>
            <TextField
              variant="outlined"
              placeholder="Yout name"
              size="small"
              value={author}
              onChange={handleAuthorChanged}
              fullWidth
              required
            />
            <TextField
              variant="outlined"
              placeholder="Leave a comment"
              minRows="4"
              size="small"
              value={content}
              onChange={handleContentChanged}
              fullWidth
              required
              multiline
            />
            {submissionError ? (
              <Alert sx={{ mt: 4 }} severity="error">
                {submissionError}
              </Alert>
            ) : null}
            <Button
              variant="contained"
              endIcon={<SendIcon />}
              sx={{ width: 100 }}
              onClick={submitComment}
              disabled={submitDisabled}
            >
              {submitLabel}
            </Button>
          </Stack>
        </ListItem>

        {data.comments.map((comment: Comment) => {
          return (
            <div key={comment.id}>
              <Divider />
              <ListItem alignItems="flex-start">
                <ListItemAvatar>
                  <Avatar>
                    <PersonIcon />
                  </Avatar>
                </ListItemAvatar>
                <ListItemText
                  primary={
                    <>
                      {comment.author}
                      <Typography
                        sx={{ display: "inline" }}
                        component="span"
                        variant="body2"
                        color="text.secondary"
                      >
                        {` — ${new Date(comment.created_at).toDateString()}`}
                      </Typography>
                    </>
                  }
                  secondary={comment.content}
                />
              </ListItem>
            </div>
          );
        })}
      </List>
    </Card>
  );
};

const ArticlePage: React.FC = () => {
  const { article_id: articleId } = useParams();
  const [{ data, loading, error }] = useAxios(`/api/articles/${articleId}`);

  if (loading) {
    return <CircularProgress />;
  }
  if (error) {
    return (
      <Box sx={{ my: 2, mx: 4 }}>
        <Typography variant="h3">Article not found</Typography>
        <Button component={Link} to="/">
          Go to top
        </Button>
      </Box>
    );
  }

  return (
    <Box sx={{ my: 2, mx: 4 }}>
      <Button component={Link} to="/">
        Go to top
      </Button>
      <ArticleCard article={data} />
      <CommentsCard />
    </Box>
  );
};

export default ArticlePage;
