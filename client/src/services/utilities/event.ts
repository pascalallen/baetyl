import env, { EnvKey } from '@services/utilities/env';

export const subscribeToPublishedEvent = (topic: string): EventSource => {
  const url = new URL(env(EnvKey.MERCURE_PUBLIC_URL) as string);
  url.searchParams.append('topic', `${env(EnvKey.APP_BASE_URL)}/${topic}`);

  return new EventSource(url.toJSON());
};
