import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import BoardView from "./BoardView";

import board from "../assets/board.test.json";

describe("Board", () => {
  it("Should display tiles", async () => {
    render(<BoardView board={board} />);

    const buttons = await screen.findAllByRole("button");

    expect(buttons).toHaveLength(50);
  });

  describe("rotateRemainingTile", function () {
    it("Should call callback when user can play and clicks on the remaining tile", async () => {
      const rotateRemainingTile = vi.fn();

      render(
        <BoardView board={board} onRotateRemainingTile={rotateRemainingTile} />
      );

      const buttons = await screen.findAllByRole("button");

      fireEvent.click(buttons[49]);
      expect(rotateRemainingTile).toBeCalledTimes(1);
    });

    it("Should not call callback when user cannot play and clicks on the remaining tile", async () => {
      const rotateRemainingTile = vi.fn();

      render(
        <BoardView
          board={{ ...board, canPlay: false }}
          onRotateRemainingTile={rotateRemainingTile}
        />
      );

      const buttons = await screen.findAllByRole("button");

      fireEvent.click(buttons[49]);
      expect(rotateRemainingTile).toBeCalledTimes(0);
    });
  });
});
