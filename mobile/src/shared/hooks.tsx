import type { Dispatch, SetStateAction } from "react";

import { useOutletContext } from "react-router-dom";

import type { User } from "../user/entity";
import type { UserRepository } from "../user/repository";

type NullableUser = User | null;

type NullableUserStateType = [
  NullableUser,
  Dispatch<SetStateAction<NullableUser>>
];

interface NullableUserContextType {
  user: NullableUser;
  setUser: Dispatch<SetStateAction<NullableUser>>;
}

export function useUser(): NullableUserStateType {
  const { user, setUser } = useOutletContext<NullableUserContextType>();
  return [user, setUser];
}

interface RemoteUserRepositoryContextType {
  remoteUserRepository: UserRepository;
}
export function useRemoteUserRepository(): UserRepository {
  const { remoteUserRepository } =
    useOutletContext<RemoteUserRepositoryContextType>();
  return remoteUserRepository;
}
