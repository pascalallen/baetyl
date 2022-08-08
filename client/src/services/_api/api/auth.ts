import HttpMethod from '@domain/constants/HttpMethod';
import { User } from '@domain/data/User';
import { ApiResponse } from '@services/_api/ApiService';
import request from '@services/_api/request';
import { RegisterFormValues } from '@pages/register/RegisterPage';

export type SessionResponsePayload = {
  token: string;
  user: User;
  roles: string[];
  permissions: string[];
};

const register = async (params: RegisterFormValues): Promise<ApiResponse<SessionResponsePayload>> => {
  return await request.send<SessionResponsePayload>({
    method: HttpMethod.POST,
    uri: '/api/v1/auth/register',
    body: params,
    options: { auth: false }
  });
};

export default Object.freeze({
  register
});
