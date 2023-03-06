<?php

namespace App\Service;

use Doctrine\Persistence\ManagerRegistry;
use Doctrine\Persistence\ObjectManager;

use App\Entity\Board;
use App\Entity\Player;

class BoardService implements BoardServiceInterface
{
    private ObjectManager $entityManager;

    public function __construct(
        private DomainServiceInterface $domainService,
        private UserServiceInterface $userService,
        private ManagerRegistry $doctrine
    ) {
        $this->entityManager = $doctrine->getManager();
    }

    public function findByCurrentUser(int $page = 1): array
    {
        if (null === ($user = $this->userService->getCurrentUser())) {
            return [];
        }

        $boardRepository = $this->entityManager->getRepository(Board::class);
        return $boardRepository->findByUser($user, $page);
    }
}
