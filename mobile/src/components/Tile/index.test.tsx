import { describe, it, expect, afterEach, vi } from "vitest";
import { cleanup, render, screen, fireEvent } from "@testing-library/react";
import matchers from "@testing-library/jest-dom/matchers";

// Setup test environment.
expect.extend(matchers);
afterEach(cleanup);

import Tile from "./index";
import { Rotation, Shape } from "../../model/Tile";
import { BoardTile } from "../../model/Board";

describe("Tile", () => {
  const boardTile: BoardTile = {
    tile: { treasure: "A", shape: Shape.ShapeT },
    rotation: Rotation.Rotation90,
  };

  describe("props.disabled", () => {
    const data = [
      [
        {
          boardTile,
          disabled: false,
          callback: (button: HTMLElement) => {
            expect(button).not.toBeDisabled();
          },
        },
      ],
      [
        {
          boardTile,
          disabled: true,
          callback: (button: HTMLElement) => {
            expect(button).toBeDisabled();
          },
        },
      ],
    ];

    it.each(data)(
      "should be disabled according to disabled props",
      async ({ boardTile, disabled, callback }) => {
        render(
          <Tile boardTile={boardTile} disabled={disabled} onClick={() => {}} />
        );

        await screen.findByRole("button");
        callback(screen.getByRole("button"));
      }
    );
  });

  describe("click", () => {
    it("should should call on click handler", async () => {
      const mock = vi.fn(() => {});

      render(<Tile boardTile={boardTile} onClick={mock} />);

      await screen.findByRole("button");
      const button = screen.getByRole("button");

      fireEvent.click(button);
      expect(mock).toHaveBeenCalledTimes(1);
    });
  });
});
