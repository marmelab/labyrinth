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
use App\Service\Direction;
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

    #[Route('', name: 'create', methods: 'PUT')]
    public function create(Request $request): JsonResponse
    {
        $user = $this->getUser();
        if (!$user) {
            return $this->json([
                'data' => ['message' => 'You must be signed in to create a board.'],
            ], 401);
        }

        $form = json_decode($request->getContent(), true);
        $playerCount = $form['playerCount'];
        if ($playerCount < 1 || $playerCount > 4) {
            return $this->json([
                'data' => ['message' => 'Player count must be between 1 and 4.'],
            ], 400);
        }

        $board = $this->newBoard($user, $playerCount);
        return $this->json([
            'data' => $this->createBoardViewModel($user, $board),
        ]);
    }


    #[Route('', name: 'find', methods: 'GET')]
    public function find(Request $request): JsonResponse
    {
        $page = intval($request->query->get('page', 1));
        if ($page < 1) {
            $page = 1;
        }

        $boardRepository = $this->entityManager->getRepository(Board::class);

        $user = $this->getUser();
        if ($user == null) {
            return $this->json([
                'data' => $boardRepository->findByAnonymous($page),
            ]);
        }

        return $this->json([
            'data' => $boardRepository->findByUser($user, $page),
        ]);
    }

    #[Route('/{id}', name: 'find_by_id', methods: 'GET')]
    public function findById(Board $board): JsonResponse
    {
        $user = $this->getUser();
        return $this->json([
            'data' => $this->createBoardViewModel($user, $board),
        ]);
    }

    #[Route('/{id}/rotate-remaining', name: 'rotate_remaining', methods: 'POST')]
    public function postRotateRemaining(Request $request, Board $board): JsonResponse
    {
        $user = $this->getUser();
        if (!$this->canUserPlay($user, $board)) {
            return $this->json([
                'data' => ['message' => 'This is not your turn', 'severity' => 'warning'],
            ], 403);
        }

        $this->rotateRemaining($board, Rotation::CLOCKWISE);

        return $this->json([
            'data' => null,
        ]);
    }

    #[Route('/{id}/insert-tile', name: 'insert_tile', methods: 'POST')]
    public function postInsertTile(Request $request, Board $board): JsonResponse
    {
        $user = $this->getUser();
        if (!$this->canUserPlay($user, $board)) {
            return $this->json([
                'data' => ['message' => 'This is not your turn.', 'severity' => 'warning'],
            ], 403);
        }

        $form = json_decode($request->getContent(), true);

        $direction = Direction::tryFrom($form['direction']);
        if (!$direction) {
            return $this->json([
                'data' => ['message' => 'Invalid direction.', 'severity' => 'warning'],
            ], 400);
        }

        if (!in_array($form['index'], [1, 3, 5])) {
            return $this->json([
                'data' => ['message' => 'Invalid index, expected one of [1, 3, 5].', 'severity' => 'warning'],
            ], 400);
        }

        $this->insertTile($board, $direction, $form['index']);
        return $this->json([
            'data' => null,
        ]);
    }

    #[Route('/{id}/move-player', name: 'move_player', methods: 'POST')]
    public function postMovePawn(Request $request, Board $board): JsonResponse
    {
        $user = $this->getUser();
        if (!$this->canUserPlay($user, $board)) {
            return $this->json([
                'data' => ['message' => 'This is not your turn.', 'severity' => 'warning'],
            ], 403);
        }

        $form = json_decode($request->getContent(), true);

        if ($form['line'] < 0 || $form['line'] > 6) {
            return $this->json([
                'data' => ['message' => 'Invalid line', 'severity' => 'warning'],
            ], 400);
        }

        if ($form['row'] < 0 || $form['row'] > 6) {
            return $this->json([
                'data' => ['message' => 'Invalid row', 'severity' => 'warning'],
            ], 400);
        }

        $this->movePlayer($board, $form['line'], $form['row']);
        return $this->json([
            'data' => null,
        ]);
    }

    #[Route('/{id}/join', name: 'join', methods: 'POST')]
    public function postJoin(Board $board): JsonResponse
    {
        $user = $this->getUser();
        if (!$this->joinBoard($user, $board)) {
            return $this->json([
                'data' => ['message' => 'You cannot join this board'],
            ], 400);
        }
        return $this->json([
            'data' => null,
        ]);
    }

    #[Route('/{id}/place-tile-hint', name: 'place_tile_hint', methods: 'PUT')]
    public function getPlaceTileHint(Board $board): JsonResponse
    {
        $user = $this->getUser();
        if (!$this->canUserPlay($user, $board)) {
            return $this->json([
                'data' => ['message' => 'This is not your turn.', 'severity' => 'warning'],
            ], 403);
        }

        if ($board->getGameState() != static::GAME_STATE_PLACE_TILE) {
            return $this->json([
                'data' => ['message' => 'You cannot place a tile now.', 'severity' => 'warning'],
            ], 400);
        }

        $placeTileHint = $this->domainService->getPlaceTileHint($board->getState());

        if ($placeTileHint['hint'] != null && count($placeTileHint['actions']) > 0) {
            $initialState = $board->getState();
            $initialState['remainingTile']['rotation'] = $placeTileHint['hint']['remainingTileRotation'];

            $this->updateBoard($board, $initialState);

            $this->entityManager->flush();
            $this->publishUpdate($board, $placeTileHint['actions']);

            return $this->json([
                'data' => [
                    'direction' => $placeTileHint['hint']['insertDirection'],
                    'index' => $placeTileHint['hint']['insertIndex'],
                ],
            ], 200);
        }

        return $this->json([
            'data' => ['message' => 'No hint has been found.', 'severity' => 'info'],
        ], 409);
    }

    #[Route('/{id}/move-pawn-hint', name: 'move_pawn_hint', methods: 'GET')]
    public function getMovePawnHint(Board $board): JsonResponse
    {
        $user = $this->getUser();
        if (!$this->canUserPlay($user, $board)) {
            return $this->json([
                'data' => ['message' => 'This is not your turn.', 'severity' => 'warning'],
            ], 403);
        }

        if ($board->getGameState() != static::GAME_STATE_MOVE_PAWN) {
            return $this->json([
                'data' => ['message' => 'You cannot move pawn now.', 'severity' => 'warning'],
            ], 400);
        }

        $movePawnHint = $this->domainService->getMovePawnHint($board->getState());

        if ($movePawnHint['hint'] != null) {
            return $this->json([
                'data' => $movePawnHint['hint'],
            ], 200);
        }

        return $this->json([
            'data' => ['message' => 'No hint has been found.', 'severity' => 'info'],
        ], 409);
    }
}
