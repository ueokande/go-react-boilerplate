import React from "react";
import Grid from "@mui/material/Grid";
import AddIcon from "@mui/icons-material/Add";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardHeader from "@mui/material/CardHeader";
import CircularProgress from "@mui/material/CircularProgress";
import PersonIcon from "@mui/icons-material/Person";
import Typography from "@mui/material/Typography";
import useAxios from "axios-hooks";
import { Link } from "react-router-dom";

type ArticleSummary = {
  id: number;
  title: string;
  author: string;
  summary: string;
  created_at: string;
};

const ArticleIndexPage: React.FC = () => {
  const [{ data, loading, error }] = useAxios("/api/articles");

  if (loading) {
    return <CircularProgress />;
  }
  if (error) {
    return <p>{error}</p>;
  }

  return (
    <>
      <Box sx={{ my: 2, mx: 4 }}>
        <Grid container justifyContent="flex-end">
          <Button
            variant="contained"
            startIcon={<AddIcon />}
            component={Link}
            to={`/articles/new`}
          >
            New Article
          </Button>
        </Grid>

        {data.article_summaries.map((a: ArticleSummary) => {
          return (
            <Card key={a.id} sx={{ my: 1 }}>
              <CardHeader
                avatar={
                  <Avatar>
                    <PersonIcon />
                  </Avatar>
                }
                title={a.author}
                subheader={new Date(a.created_at).toDateString()}
              />
              <CardContent>
                <Typography
                  variant="h5"
                  color="primary"
                  component={Link}
                  to={`/article/${a.id}`}
                >
                  {a.title}
                </Typography>
                <Typography
                  sx={{ mt: 2.0, whiteSpace: "pre-wrap" }}
                  color="text.secondary"
                >
                  {a.summary}...
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small" component={Link} to={`/article/${a.id}`}>
                  Read More
                </Button>
              </CardActions>
            </Card>
          );
        })}
      </Box>
    </>
  );
};

export default ArticleIndexPage;
