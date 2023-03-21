<?php

namespace App\ViewModel;

class AccessibleTilesViewModel
{
    public function __construct(
        private bool $isShortestPath,
        private array $coordinates,
    ) {
    }

    public function getIsShortestPath(): bool
    {
        return $this->isShortestPath;
    }

    public function getCoordinates(): array
    {
        return $this->coordinates;
    }
}
