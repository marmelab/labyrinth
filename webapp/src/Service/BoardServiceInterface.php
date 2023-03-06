<?php

namespace App\Service;

use App\Entity\Board;

interface BoardServiceInterface
{
    const SESSION_PLAYER_KEY = UserService::SESSION_PLAYER_KEY;

    function findByCurrentUser(int $page = 1): array;
}
