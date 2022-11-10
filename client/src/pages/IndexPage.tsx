import React, { ReactElement, useEffect } from 'react';
import { subscribeToPublishedEvent } from '@services/utilities/event';

const IndexPage = (): ReactElement => {
  // TODO: Remove when finished debugging
  useEffect(() => {
    const mercureJobListener = subscribeToPublishedEvent('jobs/1234');
    mercureJobListener.onmessage = event => {
      const eventData = JSON.parse(event.data);
      console.log(eventData);
    };
  }, []);

  return <h1>Hello, World!</h1>;
};

export default IndexPage;
