import { Permission } from '@domain/data/Permission';

export type Role = {
  id: string;
  name: string;
  permissions?: Permission[];
  created_at: string;
  modified_at: string;
};
