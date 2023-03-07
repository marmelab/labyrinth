import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import BoardView from "./index";

import board from "./index.test.json";

describe("Board", () => {
  it("Should display tiles", async () => {
    render(<BoardView board={board} />);

    const buttons = await screen.findAllByRole("button");

    expect(buttons).toHaveLength(50);
  });
});
