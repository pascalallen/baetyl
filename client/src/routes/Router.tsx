import React, { ReactElement } from 'react';
import { Route } from 'react-router';
import { BrowserRouter, Routes } from 'react-router-dom';
import IndexPage from '@pages/IndexPage';
import RegisterPage from '@pages/register/RegisterPage';

const Router = (): ReactElement => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/">
          <Route index element={<IndexPage />} />
          <Route path="register" element={<RegisterPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
