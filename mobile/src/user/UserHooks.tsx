import { useMutation, useQuery } from "react-query";

import { type User, type NullableUser } from "./UserTypes";

type UserResponse = { error?: string; data?: NullableUser };
export const useGetIdentityQuery = ({
  onSuccess,
}: {
  onSuccess?: (user: NullableUser) => void;
}) =>
  useQuery<NullableUser, string>(
    "identity",
    async () => {
      const response = await fetch(`/api/v1/auth/identity`);

      const responseContent: UserResponse = await response.json();
      if (response.status != 200) {
        throw responseContent.error!;
      }

      return responseContent.data!;
    },
    {
      onSuccess,
    }
  );

type SignUpVariables = { email: string; username: string; password: string };

export const useSignUpMutation = () =>
  useMutation<User, string, SignUpVariables, void>(
    async ({ email, username, password }) => {
      const response = await fetch(`/api/v1/auth/sign-up`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, username, plainPassword: password }),
      });

      const responseContent: UserResponse = await response.json();
      if (response.status != 200) {
        throw responseContent.error!;
      }

      return responseContent.data!;
    }
  );

type SignInVariables = { username: string; password: string };

export const useSignInMutation = () =>
  useMutation<User, string, SignInVariables, void>(
    async ({ username, password }) => {
      const response = await fetch(`/api/v1/auth/sign-in`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      const responseContent: UserResponse = await response.json();
      if (response.status != 200) {
        throw responseContent.error!;
      }

      return responseContent.data!;
    }
  );

export const useSignOutMutation = () =>
  useMutation<User, string, void, void>(async () => {
    const response = await fetch(`/api/v1/auth/sign-out`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({}),
    });

    const responseContent: UserResponse = await response.json();
    if (response.status != 200) {
      throw responseContent.error!;
    }

    return responseContent.data!;
  });
