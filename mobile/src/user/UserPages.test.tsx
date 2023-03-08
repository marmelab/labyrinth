import { describe, it, expect, afterEach, vi } from "vitest";
import { mock } from "vitest-mock-extended";

import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

import { RenderRouteContext } from "../shared/SharedTestUtils";
import { UserRepository } from "./UserRepository";
import { SignIn, SignOut } from "./UserPages";

expect.extend(matchers);
afterEach(() => {
  cleanup();
  vi.clearAllMocks();
});

describe("SignIn", () => {
  const testUserName = "test-user";

  it("Should sign in user if name is set", async () => {
    const userRepository = mock<UserRepository>();
    userRepository.me.mockResolvedValue(null);
    userRepository.signIn.mockResolvedValueOnce({ id: 1, name: testUserName });

    render(
      <RenderRouteContext userRepository={userRepository}>
        <SignIn />
      </RenderRouteContext>
    );

    const nameField = await screen.findByLabelText("Username");
    const button = await screen.findByRole("button");

    fireEvent.change(nameField, {
      target: { value: testUserName },
    });

    fireEvent.click(button);

    expect(userRepository.me).toBeCalledTimes(1);
    expect(userRepository.signIn).toBeCalledWith(testUserName);
  });

  it("Should do nothing if no username is provided", async () => {
    const userRepository = mock<UserRepository>();
    userRepository.me.mockResolvedValue(null);

    render(
      <RenderRouteContext userRepository={userRepository}>
        <SignIn />
      </RenderRouteContext>
    );

    const button = await screen.findByRole("button");

    fireEvent.click(button);

    expect(userRepository.me).toBeCalledTimes(1);
    expect(userRepository.signIn).not.toBeCalled();
  });
});

describe("SignOut", () => {
  const testUser = { id: 1, name: "test-user" };

  it("Should sign out user", async () => {
    const userRepository = mock<UserRepository>();
    userRepository.me.mockResolvedValue(testUser);
    userRepository.signOut.mockReturnValue(Promise.resolve());

    render(
      <RenderRouteContext userRepository={userRepository}>
        <SignOut />
      </RenderRouteContext>
    );

    expect(userRepository.me).toBeCalledTimes(1);
    expect(userRepository.signOut).toBeCalledTimes(1);
  });
});
