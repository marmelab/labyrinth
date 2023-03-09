import { describe, it, expect, beforeEach, afterEach, vi } from "vitest";
import { mock } from "vitest-mock-extended";

import { cleanup, render, screen, fireEvent } from "@testing-library/react";

import userEvent from "@testing-library/user-event";
import matchers from "@testing-library/jest-dom/matchers";

expect.extend(matchers);

import { MemoryRouter } from "react-router-dom";
import type { UserRepository } from "./UserTypes";

let userRepository = mock<UserRepository>();

vi.mock("./UserRepository", () => {
  return {
    userRepository,
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
  const testUser = { id: 1, name: testName };

  async function goToSignIn() {
    const signInLink = await screen.findByText("Sign In / Sign Up");

    await user.click(signInLink);

    const signInNameField = await screen.findByLabelText("Username");
    const signInButton = await screen.findByRole("button");

    return [signInNameField, signInButton];
  }

  beforeEach(() => {
    userRepository.getIdentity.mockResolvedValue(null);
  });

  it("Should sign in user if name is set", async () => {
    userRepository.signIn.mockResolvedValueOnce(testUser);

    render(
      <MemoryRouter>
        <App />
      </MemoryRouter>
    );

    const [nameField, button] = await goToSignIn();

    fireEvent.change(nameField, {
      target: { value: testName },
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

    const [_, button] = await goToSignIn();

    fireEvent.click(button);

    const signInText = await screen.findAllByText("Sign In / Sign Up");

    expect(signInText).toHaveLength(2);
  });
});
