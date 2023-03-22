<?php

namespace App\ViewModel;

use App\Service\Direction;

class LastInsertionViewModel
{

    public function __construct(
        private Direction $direction,
        private int $index,
    ) {
    }

    public function getDirection(): Direction
    {
        return $this->direction;
    }

    public function getIndex(): int
    {
        return $this->index;
    }
}
