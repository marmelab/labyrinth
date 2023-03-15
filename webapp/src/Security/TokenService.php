<?php

namespace App\Security;

use DateTime;
use DateInterval;

use Lexik\Bundle\JWTAuthenticationBundle\Services\JWTTokenManagerInterface;
use Symfony\Component\HttpFoundation\Cookie;
use Symfony\Component\HttpFoundation\Response;

use App\Entity\User;

class TokenService implements TokenServiceInterface
{
    const JWT_TOKEN_COOKIE_NAME = '_JWT';

    public function __construct(
        protected JWTTokenManagerInterface $JWTManager,
    ) {
    }

    function setToken(Response $response, User $user): Response
    {
        $token = $this->JWTManager->createFromPayload($user, [
            'role' => strtolower($user->getRole()),
        ]);
        $expire  = new DateTime('now');
        $expire->add(new DateInterval('PT1H'));

        $response->headers->setCookie(new Cookie(
            static::JWT_TOKEN_COOKIE_NAME,
            $token,
            $expire,
            '/',
            null,
            true
        ));

        return $response;
    }
}
