<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\User;

#[Route('/api/v1/auth')]
class SecurityApiController extends AbstractController
{
    #[Route(path: '/sign-in', name: 'api_login')]
    public function login(): ?JsonResponse
    {
        return $this->json([
            'data' => $this->getUser(),
        ]);
    }
}
