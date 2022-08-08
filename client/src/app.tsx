import React, { ReactElement } from 'react';
import { createRoot } from 'react-dom/client';
import { Helmet, HelmetProvider } from 'react-helmet-async';
import Router from '@routes/Router';
import { storesInstance, StoresProvider } from '@stores/Stores';
import '@assets/scss/app.scss';

const App = (): ReactElement => {
  return (
    <React.StrictMode>
      <StoresProvider value={storesInstance}>
        <HelmetProvider>
          <Helmet>
            <title>Baetyl</title>
            <meta name="robots" content="index, follow" />
          </Helmet>
          <Router />
        </HelmetProvider>
      </StoresProvider>
    </React.StrictMode>
  );
};

const container = document.getElementById('root');
if (container !== null) {
  const root = createRoot(container);
  root.render(<App />);
}
