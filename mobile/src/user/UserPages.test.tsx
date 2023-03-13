import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { mock } from "vitest-mock-extended";

import { cleanup, render, screen, fireEvent } from "@testing-library/react";

import userEvent from "@testing-library/user-event";
import matchers from "@testing-library/jest-dom/matchers";

expect.extend(matchers);

import { MemoryRouter } from "react-router-dom";

import type { BoardRepository } from "../board/BoardTypes";

let boardRepository = mock<BoardRepository>();

const useGetIdentityQuery = vi.fn();
const useSignInMutation = vi.fn();
const useSignOutMutation = vi.fn();
const useUserContext = vi.fn();

vi.mock("../board/BoardRepository", () => {
  return {
    boardRepository,
  };
});
vi.mock("./UserHooks", () => {
  return {
    useGetIdentityQuery,
    useSignInMutation,
    useSignOutMutation,
  };
});

afterEach(() => {
  cleanup();
  vi.resetAllMocks();
});

import App from "../App";

describe("SignIn", () => {
  const user = userEvent.setup();

  const testName = "test-user";
  const testEmail = "test@example.org";
  const testUser = {
    id: 1,
    username: testName,
    email: testEmail,
    roles: ["ROLE_USER"],
  };

  async function goToSignIn() {
    const signInLink = await screen.findByText("Sign In");

    await user.click(signInLink);

    const signInEmailField = await screen.findByLabelText("Email");
    const signInPasswordField = await screen.findByLabelText("Password");
    const signInButton = await screen.findByRole("button", {
      name: "Sign In",
    });

    return [signInEmailField, signInPasswordField, signInButton];
  }

  beforeEach(() => {
    boardRepository.list.mockResolvedValue([]);
  });

  it("Should sign in user if email and password are set", async () => {
    const mutateSignInAsync = vi.fn();
    mutateSignInAsync.mockResolvedValue(testUser);
    useSignInMutation.mockReturnValue({
      isError: false,
      mutateAsync: mutateSignInAsync,
    });

    render(
      <MemoryRouter>
        <App />
      </MemoryRouter>
    );

    const [emailField, passwordField, button] = await goToSignIn();

    fireEvent.change(emailField, {
      target: { value: testName },
    });

    fireEvent.change(passwordField, {
      target: { value: "password" },
    });

    await user.click(button);

    await screen.findByText(`${testName}`);
  });

  it("Should do nothing if no username is provided", async () => {
    render(
      <MemoryRouter>
        <App />
      </MemoryRouter>
    );

    useSignInMutation.mockReturnValue(() => ({
      isError: false,
    }));

    const [_, button] = await goToSignIn();

    fireEvent.click(button);

    const signInText = await screen.findAllByText("Sign In");

    expect(signInText).toHaveLength(2);
  });
});
