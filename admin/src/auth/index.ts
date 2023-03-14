import { AuthProvider } from "react-admin";

const localStorageUserKey = "user";

type User = {
  id: number;
  username: string;
  email: string;
  roles: string[];
};
type UserResponse = {
  error?: string;
  data?: User | null;
  token?: string;
};

async function fetchUser(url: string, body?: string) {
  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body,
  });

  const responseContent: UserResponse = await response.json();
  if (response.status != 200) {
    throw { message: responseContent.error! };
  }

  localStorage.setItem(
    localStorageUserKey,
    JSON.stringify(responseContent.data!)
  );

  return responseContent.data!;
}

export const authProvider: AuthProvider = {
  // authentication
  login: async ({ username, password }) =>
    fetchUser("/api/v1/auth/sign-in", JSON.stringify({ username, password })),
  checkError: (error) => {
    const status = error.status;
    if (status !== 200) {
      localStorage.removeItem(localStorageUserKey);
      return Promise.reject();
    }
    return Promise.resolve();
  },
  checkAuth: async (params) => {
    const user = await fetchUser("/api/v1/auth/check");
    if (!user) {
      localStorage.removeItem(localStorageUserKey);
      throw new Error("User not logged in");
    }
  },
  logout: async () => {
    await fetchUser("/api/v1/auth/sign-out");
    localStorage.removeItem(localStorageUserKey);
  },
  getIdentity: () => {
    try {
      const localStorageUser = localStorage.getItem(localStorageUserKey);
      if (!localStorageUser) {
        throw new Error("User not logged in");
      }

      const user = JSON.parse(localStorageUser);
      return Promise.resolve({
        id: user.id,
        fullName: user.username,
      });
    } catch (error) {
      return Promise.reject(error);
    }
  },
  getPermissions: () => {
    try {
      const localStorageUser = localStorage.getItem(localStorageUserKey);
      if (!localStorageUser) {
        throw new Error("User not logged in");
      }

      const user = JSON.parse(localStorageUser);
      return Promise.resolve(user.roles);
    } catch (error) {
      return Promise.reject(error);
    }
  },
};
