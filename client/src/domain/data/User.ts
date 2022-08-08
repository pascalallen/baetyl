import { Role } from '@domain/data/Role';

export type User = {
  id: string;
  first_name: string;
  last_name: string;
  email_address: string;
  roles?: Role[];
  created_at: string;
  modified_at: string;
};
