<?php

namespace App\Service;

use App\Entity\Player;

interface UserServiceInterface
{
    const SESSION_PLAYER_KEY = 'player';

    function getCurrentUser(): ?Player;
}
