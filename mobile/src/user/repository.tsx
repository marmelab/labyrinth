import type { User } from "./entity";

export interface UserRepository {
  signIn(name: string): Promise<User>;
  signOut(): Promise<void>;
  me(): Promise<User | null>;
}

export class RemoteUserRepository {
  async signIn(name: string): Promise<User> {
    const response = await fetch(`/api/v1/auth/sign-in`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name,
      }),
    });

    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: User } = await response.json();
    return responseContent.data;
  }
  async signOut(): Promise<void> {
    const response = await fetch(`/api/v1/auth/sign-out`, {
      method: "POST",
    });

    if (response.status != 200) {
      throw response;
    }
  }

  async me(): Promise<User | null> {
    const response = await fetch(`/api/v1/auth/me`);

    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: User | null } = await response.json();
    return responseContent.data;
  }
}
