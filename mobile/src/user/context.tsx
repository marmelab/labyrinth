import { createContext, useState } from "react";

import { User } from "./entity";
import { UserRepository } from "./repository";

export const UserContext = createContext<User | null>(null);
