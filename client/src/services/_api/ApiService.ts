import axios, { AxiosError, AxiosInstance, AxiosRequestConfig } from 'axios';
import { StatusCodes } from 'http-status-codes';
import jwtDecode from 'jwt-decode';
import moment from 'moment';
import { Json } from '@domain/types/Json';
import SessionStore from '@stores/SessionStore';
import { SessionResponsePayload } from '@services/_api/api/auth';
import { listToMap } from '@services/utilities/collections';

export type SuccessResponseBody<T> = {
  status: 'success';
  data: T;
};

export type ApiResponse<T> = {
  statusCode: number;
  body: SuccessResponseBody<T>;
};

export type FailResponseBody = {
  status: 'fail';
  data: { [field: string]: string };
};

export type FailApiResponse = {
  statusCode: number;
  body: FailResponseBody;
};

export type ErrorResponseBody = {
  status: 'error';
  message: string;
  code?: string;
};

export type ErrorApiResponse = {
  statusCode: number;
  body: ErrorResponseBody;
};

type RequestBody = Json;

type AppError = {
  message: string;
};

type Options = {
  auth: boolean;
  sessionStore?: SessionStore;
};

const addAuthorizationHeader = (api: AxiosInstance, sessionStore?: SessionStore): void => {
  const session = sessionStore?.getSession();
  const token = session?.token;

  if (token) {
    api.interceptors.request.use(config => {
      if (config.headers) {
        config.headers.Authorization = `Bearer ${token}`;

        return config;
      }

      return config;
    });
  }
};

export class ApiService {
  private readonly api: AxiosInstance;
  private readonly sessionStore?: SessionStore;
  private readonly readConfig: AxiosRequestConfig = { headers: { Accept: 'application/json' } };
  private readonly writeConfig: AxiosRequestConfig = {
    headers: { 'Accept': 'application/json', 'Content-Type': 'application/json' }
  };

  public constructor(options: Options) {
    this.api = axios.create();
    this.sessionStore = options.sessionStore;

    if (options.auth) {
      if (!this.sessionStore) {
        throw makeAppError('Must pass session store when auth needed');
      }

      if (this.isTokenExpired()) {
        this.refreshToken().then(() => addAuthorizationHeader(this.api, this.sessionStore));

        return;
      }

      addAuthorizationHeader(this.api, this.sessionStore);
    }
  }

  public async get<T>(url: string): Promise<ApiResponse<T>> {
    const axiosResponse = await this.api.get(url, this.readConfig);

    return {
      statusCode: axiosResponse.status,
      body: axiosResponse.data
    };
  }

  public async post<T>(url: string, body: RequestBody = {}): Promise<ApiResponse<T>> {
    const axiosResponse = await this.api.post(url, body, this.writeConfig);

    return {
      statusCode: axiosResponse.status,
      body: axiosResponse.data
    };
  }

  public async put<T>(url: string, body: RequestBody = {}): Promise<ApiResponse<T>> {
    const axiosResponse = await this.api.put(url, body, this.writeConfig);

    return {
      statusCode: axiosResponse.status,
      body: axiosResponse.data
    };
  }

  public async patch<T>(url: string, body: RequestBody = {}): Promise<ApiResponse<T>> {
    const axiosResponse = await this.api.patch(url, body, this.writeConfig);

    return {
      statusCode: axiosResponse.status,
      body: axiosResponse.data
    };
  }

  public async delete<T>(url: string): Promise<ApiResponse<T>> {
    const axiosResponse = await this.api.delete(url, this.writeConfig);

    return {
      statusCode: axiosResponse.status,
      body: axiosResponse.data
    };
  }

  private async refreshToken(): Promise<void> {
    try {
      if (this.sessionStore) {
        const response = await this.patch<SessionResponsePayload>('/api/v1/auth/session');

        this.sessionStore.setSession({
          token: response.body.data?.token,
          data: jwtDecode(response.body.data?.token),
          clientIat: moment().unix(),
          roles: listToMap(response.body.data?.roles),
          permissions: listToMap(response.body.data?.permissions)
        });
      }
    } catch (err) {
      this.sessionStore?.clearSession();
    }
  }

  private isTokenExpired(): boolean {
    const session = this.sessionStore?.getSession();

    if (!session) {
      return true;
    }

    try {
      const clientIssuedAtTime = session?.clientIat ?? 0;
      const serverIssuedAtTime = session?.data?.iat ?? 0;
      const serverExpTimestamp = session?.data?.exp ?? 0;
      const clientNowTimestamp = moment().unix();
      const clientServerDifference = clientIssuedAtTime - serverIssuedAtTime || 0;
      const calculatedServerNowTimestamp = clientNowTimestamp - clientServerDifference;
      const serverTokenExpTime = moment.unix(serverExpTimestamp || 0).subtract(5, 'minutes');
      const serverNowTimestamp = moment.unix(calculatedServerNowTimestamp);
      return serverNowTimestamp.isSameOrAfter(serverTokenExpTime);
    } catch (err) {
      return true;
    }
  }
}

export const makeAppError = (errorMessage?: string): AppError => {
  return { message: errorMessage || 'An unexpected error occurred' };
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const makeApiErrorResponse = (error: AxiosError<any>): ErrorApiResponse => {
  return {
    statusCode: error.response?.status || StatusCodes.INTERNAL_SERVER_ERROR,
    body: {
      status: error.response?.data?.status || 'error',
      message: error.response?.data?.message || 'API ERROR',
      code: error.code
    }
  };
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const makeApiFailResponse = (error: AxiosError<any>): FailApiResponse => {
  return {
    statusCode: error.response?.status || StatusCodes.INTERNAL_SERVER_ERROR,
    body: {
      status: error.response?.data?.status || 'fail',
      data: error.response?.data
    }
  };
};

export const makeApiService = (options: Options): ApiService => {
  return new ApiService(options);
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const isApiFail = (error: any | unknown): error is FailApiResponse => {
  return (<FailApiResponse>error).body !== undefined && (<FailApiResponse>error).body.status === 'fail';
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const isApiError = (error: any | unknown): error is ErrorApiResponse => {
  return (<ErrorApiResponse>error).body !== undefined && (<ErrorApiResponse>error).body.status === 'error';
};

export default ApiService;
