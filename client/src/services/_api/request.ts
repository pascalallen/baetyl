import HttpMethod from '@domain/constants/HttpMethod';
import { Json } from '@domain/types/Json';
import SessionStore from '@stores/SessionStore';
import {
  ApiResponse,
  makeApiErrorResponse,
  makeApiFailResponse,
  makeApiService,
  makeAppError
} from '@services/_api/ApiService';

export type RequestProps = {
  method: string;
  uri: string;
  body?: Json;
  options: {
    auth: boolean;
    sessionStore?: SessionStore;
  };
};

const send = async <T>(props: RequestProps): Promise<ApiResponse<T>> => {
  const { method, uri, body, options } = props;

  try {
    const api = makeApiService(options);

    switch (method) {
      case HttpMethod.GET: {
        return await api.get(uri);
      }
      case HttpMethod.POST: {
        return await api.post(uri, body);
      }
      case HttpMethod.PUT: {
        return await api.put(uri, body);
      }
      case HttpMethod.PATCH: {
        return await api.patch(uri, body);
      }
      case HttpMethod.DELETE: {
        return await api.delete(props.uri);
      }
    }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (error: any) {
    if (error.response?.data?.status === 'error') {
      throw makeApiErrorResponse(error);
    }

    if (error.response?.data?.status === 'fail') {
      throw makeApiFailResponse(error);
    }
  }

  throw makeAppError('Unknown method');
};

export default Object.freeze({
  send
});
