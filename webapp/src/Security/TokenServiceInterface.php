<?php

namespace App\Security;

use Symfony\Component\HttpFoundation\Response;

use App\Entity\User;

interface TokenServiceInterface
{
    function setToken(Response $response, User $user): Response;
}
