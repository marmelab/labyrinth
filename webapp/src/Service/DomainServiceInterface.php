<?php

namespace App\Service;

interface DomainServiceInterface
{
    function newBoard(): array;

    function rotateRemainingTile(array $board, Rotation $rotation): array;
}