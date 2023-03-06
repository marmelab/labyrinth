<?php

namespace App\Service;

use Doctrine\Persistence\ManagerRegistry;
use Doctrine\Persistence\ObjectManager;
use Symfony\Component\HttpFoundation\RequestStack;

use App\Entity\Player;

class UserService implements UserServiceInterface
{
    private ObjectManager $entityManager;

    public function __construct(
        private ManagerRegistry $doctrine,
        private RequestStack $requestStack
    ) {
        $this->entityManager = $doctrine->getManager();
    }

    public function getCurrentUser(): ?Player
    {
        $request = $this->requestStack->getCurrentRequest();
        $user = $request->getSession()->get(static::SESSION_PLAYER_KEY);
        if ($user == NULL) {
            return NULL;
        }

        $playerRepository = $this->entityManager->getRepository(Player::class);
        return $playerRepository->find($user->getId());
    }
}
