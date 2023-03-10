import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import BoardView from "./BoardView";

import board from "../assets/board.test.json";
import { Direction } from "../BoardTypes";

describe("Board", () => {
  it("Should display tiles", async () => {
    const onRotateRemainingTile = vi.fn();
    const onInsertTile = vi.fn();
    render(
      <BoardView
        board={board}
        onRotateRemainingTile={onRotateRemainingTile}
        onInsertTile={onInsertTile}
      />
    );

    const buttons = await screen.findAllByRole("button");

    expect(buttons).toHaveLength(50);
  });

  it("Should support a player turn", async () => {
    const onRotateRemainingTile = vi.fn();
    const onInsertTile = vi.fn();

    onInsertTile.mockResolvedValueOnce({ ...board, gameState: 1 });

    render(
      <BoardView
        board={board}
        onRotateRemainingTile={onRotateRemainingTile}
        onInsertTile={onInsertTile}
      />
    );

    const buttons = await screen.findAllByRole("button");

    fireEvent.click(buttons[1]);
    expect(onInsertTile).toBeCalledWith(Direction.Top, 1, expect.anything());
  });

  describe("rotateRemainingTile", function () {
    it("Should call callback when user can play and clicks on the remaining tile", async () => {
      const onRotateRemainingTile = vi.fn();
      const onInsertTile = vi.fn();

      render(
        <BoardView
          board={board}
          onRotateRemainingTile={onRotateRemainingTile}
          onInsertTile={onInsertTile}
        />
      );

      const buttons = await screen.findAllByRole("button");

      fireEvent.click(buttons[49]);
      expect(onRotateRemainingTile).toBeCalledTimes(1);
    });

    it("Should not call callback when user cannot play and clicks on the remaining tile", async () => {
      const onRotateRemainingTile = vi.fn();
      const onInsertTile = vi.fn();

      render(
        <BoardView
          board={{ ...board, canPlay: false }}
          onRotateRemainingTile={onRotateRemainingTile}
          onInsertTile={onInsertTile}
        />
      );

      const buttons = await screen.findAllByRole("button");

      fireEvent.click(buttons[49]);
      expect(onRotateRemainingTile).toBeCalledTimes(0);
    });
  });
});
