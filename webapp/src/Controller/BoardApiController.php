<?php

namespace App\Controller;

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;

use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Serializer\SerializerInterface;

use App\Entity\Board;
use App\Service\DomainServiceInterface;
use App\Service\Rotation;

#[Route('/api/v1/board', name: 'board_api_')]
class BoardApiController extends BoardBaseController
{
    public function __construct(
        protected DomainServiceInterface $domainService,
        protected ManagerRegistry $doctrine,
        protected HubInterface $hub,
        protected SerializerInterface $serializer,
    ) {
        parent::__construct($domainService, $doctrine->getManager(), $hub, $serializer);
    }

    #[Route('', name: 'find', methods: 'GET')]
    public function find(Request $request): JsonResponse
    {
        $page = intval($request->query->get('page', 1));
        if ($page < 1) {
            $page = 1;
        }

        $user = $this->getCurrentUser($request);
        $boardRepository = $this->entityManager->getRepository(Board::class);
        return $this->json([
            'data' => $boardRepository->findByUser($user, $page),
        ]);
    }

    #[Route('/{id}', name: 'find_by_id', methods: 'GET')]
    public function findById(Request $request, Board $board): JsonResponse
    {
        $user = $this->getCurrentUser($request);
        return $this->json([
            'data' => $this->createBoardViewModel($user, $board),
        ]);
    }
    #[Route('/{id}/rotate-remaining', name: 'rotate_remaining', methods: 'POST')]
    public function postRotateRemaining(Request $request, Board $board): JsonResponse
    {
        $user = $this->getCurrentUser($request);
        $this->rotateRemaining($user, $board, Rotation::CLOCKWISE);

        return $this->json([
            'data' => null,
        ]);
    }
}
