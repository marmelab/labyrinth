<?php

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;
use Doctrine\Persistence\ObjectManager;
use Doctrine\ORM\EntityManager;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\Board;
use App\Entity\Player;

use App\Form\Type\InsertTileType;
use App\Form\Type\JoinBoardType;
use App\Form\Type\MovePlayerType;
use App\Form\Type\NewBoardType;
use App\Form\Type\RotateRemainingType;
use App\Service\Direction;
use App\Service\Rotation;
use App\Service\DomainServiceInterface;

class BoardController extends AbstractController
{
    const SESSION_PLAYER_KEY = PlayerController::SESSION_PLAYER_KEY;

    const TREASURE_EMOJIS = [
        '.' => ' ',
        'A' => '💌',
        'B' => '💣',
        'C' => '🛍',
        'D' => '📿',
        'E' => '🔭',
        'F' => '💎',
        'G' => '💰',
        'H' => '📜',
        'I' => '🗿',
        'J' => '🏺',
        'K' => '🔫',
        'L' => '🛡',
        'M' => '💈',
        'N' => '🛎',
        'O' => '⌛',
        'P' => '🌡',
        'Q' => '⛱',
        'R' => '🎈',
        'S' => '🎎',
        'T' => '🎁',
        'U' => '🔮',
        'V' => '📷',
        'W' => '🕯',
        'X' => '🥦',
    ];

    const PLAYER_COLORS = [
        "Blue",
        "Green",
        "Red",
        "Yellow",
    ];

    public function __construct(
        private DomainServiceInterface $domainService
    ) {
    }

    private function getPlayer(Request $request, ObjectManager $entityManager): ?Player
    {
        $player = $request->getSession()->get(static::SESSION_PLAYER_KEY);
        if ($player == NULL) {
            return NULL;
        }

        $playerRepository = $entityManager->getRepository(Player::class);
        return $playerRepository->find($player->getId());
    }

    private function getCurrentPlayer(array $boardState): array
    {
        $currentPlayer = $boardState['players'][0];
        if (count($boardState['remainingPlayers']) > 0) {
            $currentPlayerIndex = $boardState['remainingPlayers'][$boardState['currentPlayerIndex']];
            $currentPlayer = $boardState['players'][$currentPlayerIndex];
        }
        return $currentPlayer;
    }

    private function getPlayerTarget(Board $board, ?Player $player): string
    {
        if ($player == null) {
            return '';
        }

        foreach ($board->getPlayers() as $index => $gamePlayer) {
            if ($gamePlayer->getId() == $player->getId()) {
                $targets = $board->getState()['players'][$index]['targets'];
                if (count($targets) > 0) {
                    return $targets[0];
                }
            }
        }

        return '';
    }

    private function canPlayerPlay(Board $board, ?Player $player): bool
    {
        if ($player == null) {
            return false;
        }

        $boardState = $board->getState();
        if ($boardState['gameState'] == 2) {
            return false;
        }

        $stateCurrentPlayer = $this->getCurrentPlayer($boardState);
        $currentPlayer = $board->getPlayers()[$stateCurrentPlayer['color']];
        return $currentPlayer->getId() == $player->getId();
    }

    private function canPlayerJoin(?Player $player, Board $board): bool
    {
        if ($player == null) {
            return false;
        }

        foreach ($board->getPlayers() as $gamePlayer) {
            if ($gamePlayer->getId() == $player->getId()) {
                return false;
            }
        }
        return true;
    }

    #[Route('/board/new', name: 'board_new', methods: 'POST')]
    public function getNew(Request $request, ManagerRegistry $doctrine)
    {
        $entityManager = $doctrine->getManager();
        $player = $this->getPlayer($request, $entityManager);
        if ($player == NULL) {
            return $this->redirectToRoute('home');
        }

        $form = $this->createForm(NewBoardType::class, null, [
            'action' => $this->generateUrl('board_new'),
        ]);

        $form->handleRequest($request);
        if (!$form->isSubmitted() || !$form->isValid()) {
            return $this->render('board/new.html.twig', [
                'form' => $form,
            ]);
        }

        $playerCount = $form->getData()['player_count'];
        $boardState = $this->domainService->newBoard(intval($playerCount));

        $board = new Board();
        $board->setState($boardState);
        $board->addPlayer($player);
        $board->setRemainingSeats($playerCount - 1);

        $entityManager->persist($board);
        $entityManager->flush();

        return $this->redirectToRoute('board_view', [
            'id' => $board->getId(),
        ]);
    }

    #[Route('/board/{id}/view', name: 'board_view', methods: 'GET')]
    public function getView(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        $entityManager = $doctrine->getManager();
        $player = $this->getPlayer($request, $entityManager);

        if ($board->getRemainingSeats() > 0) {

            $form = $this->createForm(JoinBoardType::class, null, [
                'action' => $this->generateUrl('board_join', ['id' => $board->getId()]),
            ]);

            return $this->render('board/lobby.html.twig', [
                'board' => $board,
                'form' => $form,
                'canJoin' => $this->canPlayerJoin($player, $board),
            ]);
        }


        $players = $board->getPlayers();
        $boardState = $board->getState();
        $currentPlayer = $this->getCurrentPlayer($boardState);

        return $this->render('board/view.html.twig', [
            'board' => $board,
            'boardState' => $board->getState(),
            'currentPlayer' => $currentPlayer,
            'players' => $players,
            'canPlay' => $this->canPlayerPlay($board, $player),
            'emojis' => self::TREASURE_EMOJIS,
            'colors' => self::PLAYER_COLORS,
            'target' => $this->getPlayerTarget($board, $player),
        ]);
    }

    #[Route('/board/{id}/join', name: 'board_join', methods: 'POST')]
    public function postJoin(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        /** @var EntityManager $entityManager */
        $entityManager = $doctrine->getManager();
        $player = $this->getPlayer($request, $entityManager);
        if (!$this->canPlayerJoin($player, $board)) {
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $conn = $entityManager->getConnection();
        $conn->beginTransaction();
        $conn->setAutoCommit(false);
        try {
            // This will ensure that no concurrent updates happen on the board.
            $boardRepository = $entityManager->getRepository(Board::class);

            /** @var Board $board */
            $board = $boardRepository->find($board->getId());

            $remainingSeats = $board->getRemainingSeats();
            if ($remainingSeats == 0) {
                $conn->rollBack();
                return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
            }

            $board->setRemainingSeats($remainingSeats - 1);
            $board->addPlayer($player);

            $entityManager->flush();
            $conn->commit();
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        } catch (\Exception $e) {
            $conn->rollBack();
            throw $e;
        }
    }

    private function rotateRemaining(Request $request, ManagerRegistry $doctrine, Board $board, Rotation $rotation): Response
    {
        $entityManager = $doctrine->getManager();
        $player = $this->getPlayer($request, $entityManager);
        if (!$this->canPlayerPlay($board, $player)) {
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $form = $this->createForm(RotateRemainingType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $updatedBoard = $this->domainService->rotateRemainingTile($board->getState(), $rotation);
            $board->setState($updatedBoard);

            $entityManager->flush();
        }
        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }

    #[Route('/board/{id}/rotate-remaining-clockwise', name: 'board_rotate_remaining_clockwise', methods: 'POST')]
    public function postRotateRemainingClockwise(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->rotateRemaining($request, $doctrine, $board, Rotation::CLOCKWISE);
    }

    #[Route('/board/{id}/rotate-remaining-anticlockwise', name: 'board_rotate_remaining_anticlockwise', methods: 'POST')]
    public function postRotateRemainingAnticlockwise(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->rotateRemaining($request, $doctrine, $board, Rotation::ANTICLOCKWISE);
    }

    private function insertTile(Request $request, ManagerRegistry $doctrine, Board $board, Direction $direction): Response
    {
        $entityManager = $doctrine->getManager();
        $player = $this->getPlayer($request, $entityManager);
        if (!$this->canPlayerPlay($board, $player)) {
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $form = $this->createForm(InsertTileType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $insertTile = $form->getData();

            $updatedBoard = $this->domainService->insertTile(
                $board->getState(),
                $direction,
                intval($insertTile['index'])
            );

            $board->setState($updatedBoard);

            $entityManager->flush();
        }
        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }

    #[Route('/board/{id}/insert-tile-top', name: 'board_insert_tile_top', methods: 'POST')]
    public function postInsertTileTop(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->insertTile($request, $doctrine, $board, Direction::TOP);
    }

    #[Route('/board/{id}/insert-tile-right', name: 'board_insert_tile_right', methods: 'POST')]
    public function postInsertTileRight(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->insertTile($request, $doctrine, $board, Direction::RIGHT);
    }

    #[Route('/board/{id}/insert-tile-bottom', name: 'board_insert_tile_bottom', methods: 'POST')]
    public function postInsertTileBottom(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->insertTile($request, $doctrine, $board, Direction::BOTTOM);
    }

    #[Route('/board/{id}/insert-tile-left', name: 'board_insert_tile_left', methods: 'POST')]
    public function postInsertTileLeft(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->insertTile($request, $doctrine, $board, Direction::LEFT);
    }

    #[Route('/board/{id}/move-player', name: 'board_move_player', methods: 'POST')]
    public function postMovePlayer(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        $entityManager = $doctrine->getManager();
        $player = $this->getPlayer($request, $entityManager);
        if (!$this->canPlayerPlay($board, $player)) {
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $form = $this->createForm(MovePlayerType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $movePlayer = $form->getData();

            $updatedBoard = $this->domainService->movePlayer(
                $board->getState(),
                intval($movePlayer['line']),
                intval($movePlayer['row']),
            );

            $board->setState($updatedBoard);

            $entityManager->flush();
        }
        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }
}
