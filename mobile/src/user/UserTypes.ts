export interface User {
  id: number;
  username: string;
  email: string;
  roles: string[];
}

export type NullableUser = User | null;
