<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\Board;
use App\Service\BoardServiceInterface;

#[Route('/api/v1/board', name: 'board_api_')]
class BoardApiController extends AbstractController
{
    public function __construct(
        private BoardServiceInterface $boardService
    ) {
    }

    #[Route('', name: 'find', methods: 'GET')]
    public function find(Request $request): JsonResponse
    {
        $page = intval($request->query->get('page', 1));
        if ($page < 1) {
            $page = 1;
        }
        return $this->json($this->boardService->findByCurrentUser($page));
    }

    #[Route('/{id}', name: 'find_by_id', methods: 'GET')]
    public function findById(Board $board): JsonResponse
    {
        return $this->json($board);
    }
}
