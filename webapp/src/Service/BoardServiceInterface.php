<?php

namespace App\Service;

use App\Entity\Board;

interface BoardServiceInterface
{
    function findByCurrentUser(int $page = 1): array;
}
