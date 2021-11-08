import React from "react";
import ArticleIndexPage from "./ArticleIndexPage";
import ArticlePage from "./ArticlePage";
import NewArticlePage from "./NewArticlePage";
import NotFoundPage from "./NotFoundPage";
import { Routes, Route } from "react-router-dom";

const App: React.FC = () => {
  return (
    <Routes>
      <Route index element={<ArticleIndexPage />} />
      <Route path="/article/:article_id" element={<ArticlePage />} />
      <Route path="/articles/new" element={<NewArticlePage />} />
      <Route path="*" element={<NotFoundPage />} />
    </Routes>
  );
};

export default App;
