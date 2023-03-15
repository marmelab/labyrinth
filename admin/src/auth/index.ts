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

async function storeUser(user: User | null) {
  if (!user) {
    localStorage.removeItem(localStorageUserKey);
    throw { status: 401, message: "User not logged in" };
  } else {
    localStorage.setItem(localStorageUserKey, JSON.stringify(user));
  }
}

export const authProvider: AuthProvider = {
  // authentication
  login: async ({ username, password }) => {
    const response = await fetch("/api/v1/auth/sign-in", {
      headers: {
        "Content-Type": "application/json",
      },
      method: "POST",
      body: JSON.stringify({ username, password }),
    });

    const responseContent: UserResponse = await response.json();
    console.log(responseContent);

    if (response.status != 200) {
      throw { message: responseContent.error! };
    }

    storeUser(responseContent.data!);
  },
  checkError: (error) => {
    const status = error.status;
    console.log(error);
    if (status !== 200) {
      localStorage.removeItem(localStorageUserKey);
      return Promise.reject();
    }
    return Promise.resolve();
  },
  checkAuth: async () => {
    const response = await fetch("/api/v1/auth/check");

    const responseContent: UserResponse = await response.json();
    if (response.status != 200) {
      throw { message: responseContent.error! };
    }

    storeUser(responseContent.data!);
  },
  logout: async () => {
    try {
      await fetch("/api/v1/auth/sign-out", {
        method: "POST",
      });
    } finally {
      localStorage.removeItem(localStorageUserKey);
    }
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
      return Promise.reject({ status: 500, message: error });
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
