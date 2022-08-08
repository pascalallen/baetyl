import { TokenData } from '@domain/data/TokenData';

export type Session = {
  token: string;
  data: TokenData;
  clientIat: number;
  roles: { [name: string]: string };
  permissions: { [name: string]: string };
};
