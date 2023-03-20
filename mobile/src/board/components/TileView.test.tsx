import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import { TileView } from "./TileView";
import { BoardTile, GameState, Rotation, Shape } from "../BoardTypes";

describe("Tile", () => {
  const boardTile: BoardTile = {
    tile: { treasure: "A", shape: Shape.ShapeT },
    rotation: Rotation.Rotation90,
  };

  describe("props.onInsertTile", () => {
    it("should be disabled if not on an odd index", async () => {
      const onInsertTile = vi.fn();
      const onMovePlayer = vi.fn();

      render(
        <TileView
          boardTile={boardTile}
          canPlay={true}
          gameState={GameState.PlaceTile}
          coordinates={{ line: 0, row: 0 }}
          onInsertTile={onInsertTile}
          onMovePlayer={onMovePlayer}
          isAccessible={true}
          hint={false}
        />
      );

      await screen.findByRole("button");
      const button = screen.getByRole("button");
      expect(button).toBeDisabled();
    });

    it("should should call on click handler", async () => {
      const onInsertTile = vi.fn();
      const onMovePlayer = vi.fn();

      render(
        <TileView
          boardTile={boardTile}
          canPlay={true}
          gameState={GameState.PlaceTile}
          coordinates={{ line: 0, row: 1 }}
          onInsertTile={onInsertTile}
          onMovePlayer={onMovePlayer}
          isAccessible={true}
          hint={false}
        />
      );

      await screen.findByRole("button");
      const button = screen.getByRole("button");
      expect(button).not.toBeDisabled();

      fireEvent.click(button);
      expect(onInsertTile).toHaveBeenCalledTimes(1);
    });
  });

  describe("props.onMovePlayer", () => {
    it("should should call on click handler", async () => {
      const onInsertTile = vi.fn();
      const onMovePlayer = vi.fn();

      render(
        <TileView
          boardTile={boardTile}
          canPlay={true}
          gameState={GameState.MovePawn}
          coordinates={{ line: 0, row: 1 }}
          onInsertTile={onInsertTile}
          onMovePlayer={onMovePlayer}
          isAccessible={true}
          hint={false}
        />
      );

      await screen.findByRole("button");
      const button = screen.getByRole("button");
      expect(button).not.toBeDisabled();

      fireEvent.click(button);
      expect(onMovePlayer).toHaveBeenCalledWith(0, 1, expect.anything());
    });
  });
});
