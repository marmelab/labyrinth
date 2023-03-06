import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import { Board as State, GameState } from "../../model/Board";
import Board from "./index";

import state from "./index.test.json";

describe("Board", () => {
  it("Should display tiles", async () => {
    render(<Board state={state} />);

    const buttons = await screen.findAllByRole("button");

    expect(buttons).toHaveLength(50);
  });
});
