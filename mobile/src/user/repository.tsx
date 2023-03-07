import type { User } from "./entity";

export class UserRepository {
  static async signIn(username: string): Promise<User> {
    const response = await fetch(`/api/v1/auth/sign-in`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
      }),
    });

    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: User } = await response.json();
    return responseContent.data;
  }

  static async me(): Promise<User> {
    const response = await fetch(`/api/v1/auth/me`);

    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: User } = await response.json();
    return responseContent.data;
  }
}
