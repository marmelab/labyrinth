<?php

namespace App\Controller;

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;

use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Validator\Validator\ValidatorInterface;
use Symfony\Component\Serializer\SerializerInterface;

use App\Entity\Player;

#[Route('/api/v1/auth', name: 'auth_api_')]
class AuthApiController extends AuthBaseController
{
    public function __construct(
        protected ManagerRegistry $doctrine,
    ) {
        parent::__construct($doctrine->getManager());
    }

    private function getCurrentUser(Request $request): ?Player
    {
        $player = $request->getSession()->get(static::SESSION_PLAYER_KEY);
        if ($player == NULL) {
            return NULL;
        }

        $playerRepository = $this->entityManager->getRepository(Player::class);
        return $playerRepository->find($player->getId());
    }

    #[Route('/me', name: 'me', methods: 'GET')]
    public function getMe(Request $request): JsonResponse
    {
        $user = $this->getCurrentUser($request);
        return $this->json([
            'data' => $user,
        ]);
    }

    #[Route('/sign-in', name: 'sign_in', methods: 'POST')]
    public function postSignIn(Request $request, ValidatorInterface $validator): JsonResponse
    {
        $form = json_decode($request->getContent(), true);
        $player = new Player();
        $player->setName($form['name']);

        $errors = $validator->validate($player);
        if (count($errors) > 0) {
            return $this->json([
                'errors' => $errors,
            ], 400);
        }

        $user = $this->signInUser($request, $player->getName());
        return $this->json([
            'data' => $user,
        ]);
    }

    #[Route('/sign-out', name: 'sign_out', methods: 'POST')]
    public function postSignOut(Request $request): JsonResponse
    {
        $this->signOutUser($request);
        return $this->json([]);
    }
}
