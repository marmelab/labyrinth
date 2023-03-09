import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import TileView from "./TileView";
import { BoardTile, Rotation, Shape } from "../BoardTypes";

describe("Tile", () => {
  const boardTile: BoardTile = {
    tile: { treasure: "A", shape: Shape.ShapeT },
    rotation: Rotation.Rotation90,
  };

  describe("props.onRotateRemainingTile", () => {
    it("should should call on click handler", async () => {
      const mock = vi.fn();

      render(
        <TileView
          line={0}
          row={0}
          boardTile={boardTile}
          onRotateRemainingTile={mock}
        />
      );

      await screen.findByRole("button");
      const button = screen.getByRole("button");
      expect(button).not.toBeDisabled();

      fireEvent.click(button);
      expect(mock).toHaveBeenCalledTimes(1);
    });
  });

  describe("props.onInsertTile", () => {
    it("should be disabled if not on an odd index", async () => {
      const mock = vi.fn();

      render(
        <TileView line={0} row={0} boardTile={boardTile} onInsertTile={mock} />
      );

      await screen.findByRole("button");
      const button = screen.getByRole("button");
      expect(button).toBeDisabled();
    });

    it("should should call on click handler", async () => {
      const mock = vi.fn();

      render(
        <TileView line={0} row={1} boardTile={boardTile} onInsertTile={mock} />
      );

      await screen.findByRole("button");
      const button = screen.getByRole("button");
      expect(button).not.toBeDisabled();

      fireEvent.click(button);
      expect(mock).toHaveBeenCalledTimes(1);
    });
  });
});
