import type { User, NullableUser } from "./UserTypes";

export const userRepository = {
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
  },
  async signOut(): Promise<void> {
    const response = await fetch(`/api/v1/auth/sign-out`, {
      method: "POST",
    });

    if (response.status != 200) {
      throw response;
    }
  },
  async getIdentity(): Promise<NullableUser> {
    const response = await fetch(`/api/v1/auth/me`);

    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: NullableUser } = await response.json();

    return responseContent.data;
  },
};
