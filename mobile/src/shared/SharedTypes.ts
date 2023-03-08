import type { Dispatch, SetStateAction } from "react";

import type { User } from "../user/UserTypes";

export type NullableUser = User | null;

export type NullableUserStateType = [
  NullableUser,
  Dispatch<SetStateAction<NullableUser>>
];
