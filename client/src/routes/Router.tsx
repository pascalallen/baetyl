import React, { ReactElement } from 'react';
import { Route } from 'react-router';
import { BrowserRouter, Routes } from 'react-router-dom';
import IndexPage from '@pages/IndexPage';

const Router = (): ReactElement => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/">
          <Route index element={<IndexPage />} />
          <Route path="test" element={<h1>Another Page!</h1>} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
