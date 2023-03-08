import type { Dispatch, SetStateAction } from "react";

import { useOutletContext } from "react-router-dom";

import type { NullableUser, NullableUserStateType } from "./SharedTypes";
import { RemoteUserRepository, UserRepository } from "../user/UserRepository";

interface NullableUserContextType {
  user: NullableUser;
  setUser: Dispatch<SetStateAction<NullableUser>>;
}

export function useUser(): NullableUserStateType {
  const { user, setUser } = useOutletContext<NullableUserContextType>();
  return [user, setUser];
}

export function useUserRepository(): UserRepository {
  return new RemoteUserRepository();
}
