<?php

namespace App\Controller;

use Lexik\Bundle\JWTAuthenticationBundle\Services\JWTTokenManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Cookie;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\User;
use DateInterval;
use DateTime;

#[Route('/api/v1/auth')]
class SecurityApiController extends AbstractController
{
    const JWT_TOKEN_COOKIE_NAME = '_JWT';

    public function __construct(private JWTTokenManagerInterface $JWTManager)
    {
    }

    private function setToken(JsonResponse $response, User $user): JsonResponse
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

    #[Route(path: '/sign-in', name: 'api_login')]
    public function login(): JsonResponse
    {
        /** @var User */
        $user = $this->getUser();
        if ($user) {
            $reponse = $this->json([
                'data' => $user,
            ]);
            return $this->setToken($reponse, $user);
        }

        return $this->json([
            'data' => null,
        ]);
    }

    #[Route(path: '/check', methods: 'POST')]
    public function check(): JsonResponse
    {
        /** @var User */
        $user = $this->getUser();
        if ($user) {
            $reponse = $this->json([
                'data' => $user,
            ]);
            return $this->setToken($reponse, $user);
        }

        return $this->json([
            'data' => null,
        ]);
    }
}
