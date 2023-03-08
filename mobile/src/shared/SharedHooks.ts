import type { Dispatch, SetStateAction } from "react";

import { useOutletContext } from "react-router-dom";

import type { NullableUser, NullableUserStateType } from "./SharedTypes";
import type { UserRepository } from "../user/UserRepository";

interface NullableUserContextType {
  user: NullableUser;
  setUser: Dispatch<SetStateAction<NullableUser>>;
}

export function useUser(): NullableUserStateType {
  const { user, setUser } = useOutletContext<NullableUserContextType>();
  return [user, setUser];
}

interface UserRepositoryContextType {
  userRepository: UserRepository;
}
export function useUserRepository(): UserRepository {
  const { userRepository } = useOutletContext<UserRepositoryContextType>();
  return userRepository;
}
