import {
  useContext,
  createContext,
  type Dispatch,
  type SetStateAction,
} from "react";

import type { NullableUser } from "./UserTypes";

export type UserContextType = [
  NullableUser,
  Dispatch<SetStateAction<NullableUser>>
];

export const UserContext = createContext<UserContextType>([null, () => {}]);

export function useUserContext() {
  return useContext(UserContext);
}
