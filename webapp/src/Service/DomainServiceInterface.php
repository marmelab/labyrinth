<?php

namespace App\Service;

interface DomainServiceInterface
{
    function newBoard(int $playerCount): array;

    function rotateRemainingTile(array $board, Rotation $rotation): array;

    function insertTile(array $board, Direction $direction, int $index): array;

    function movePlayer(array $board, int $line, int $row): array;

    function getPlaceTileHint(array $board): array;
}
