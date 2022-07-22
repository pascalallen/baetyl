import React, { ReactElement } from 'react';
import { createRoot } from 'react-dom/client';
import { Helmet, HelmetProvider } from 'react-helmet-async';
import Router from '@routes/Router';
import '@assets/scss/app.scss';

const App = (): ReactElement => {
  return (
    <React.StrictMode>
      <HelmetProvider>
        <Helmet>
          <title>Baetyl</title>
          <meta name="robots" content="index, follow" />
        </Helmet>
        <Router />
      </HelmetProvider>
    </React.StrictMode>
  );
};

const container = document.getElementById('root');
const root = createRoot(container!);
root.render(<App />);
