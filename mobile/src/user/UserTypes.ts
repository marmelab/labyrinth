export interface User {
  id: number;
  name: string;
}

export type NullableUser = User | null;

export interface UserRepository {
  signIn(name: string): Promise<User>;
  signOut(): Promise<void>;
  getIdentity(): Promise<NullableUser>;
}
