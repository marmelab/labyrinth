import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { mock } from "vitest-mock-extended";

import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

expect.extend(matchers);

let setUser: any;
let userRepository: any;
let navigate: any;

vi.mock("../shared/SharedHooks", () => {
  return {
    useUser: function () {
      return [null, setUser];
    },
    useUserRepository: function () {
      return userRepository;
    },
  };
});

vi.mock("react-router-dom", () => {
  return {
    useNavigate: function () {
      return navigate;
    },
  };
});

beforeEach(() => {
  setUser = vi.fn();
  userRepository = mock<UserRepository>();
  navigate = vi.fn();
});

afterEach(() => {
  cleanup();
  vi.clearAllMocks();
});

import { SignIn, SignOut } from "./UserPages";
import { UserRepository } from "./UserRepository";

describe("SignIn", () => {
  const testName = "test-user";
  const testUser = { id: 1, name: testName };

  it("Should sign in user if name is set", async () => {
    userRepository.signIn.mockResolvedValueOnce(testUser);

    render(<SignIn />);

    const nameField = await screen.findByLabelText("Username");
    const button = await screen.findByRole("button");

    fireEvent.change(nameField, {
      target: { value: testName },
    });

    fireEvent.click(button);

    await screen.findAllByText("");

    expect(userRepository.signIn).toBeCalledWith(testName);
    expect(setUser).toBeCalledWith(testUser);
    expect(navigate).toBeCalledWith("/");
  });

  it("Should do nothing if no username is provided", async () => {
    const userRepository = mock<UserRepository>();

    render(<SignIn />);

    const button = await screen.findByRole("button");

    fireEvent.click(button);

    await screen.findAllByText("");

    expect(userRepository.signIn).not.toBeCalled();
    expect(setUser).not.toBeCalled();
    expect(navigate).not.toBeCalled();
  });
});

describe("SignOut", () => {
  it("Should sign out user", async () => {
    userRepository.signOut.mockReturnValue(Promise.resolve());

    render(<SignOut />);

    await screen.findAllByText("");

    expect(userRepository.signOut).toBeCalledTimes(1);
    expect(setUser).toBeCalledWith(null);
    expect(navigate).toBeCalledWith("/");
  });
});
