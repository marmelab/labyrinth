<?php

namespace App\Service;

interface DomainServiceInterface
{
    function newBoard(): array;

    function rotateRemainingTile(array $board, Rotation $rotation): array;

    function insertTile(array $board, Direction $direction, int $index): array;
}