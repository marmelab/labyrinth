<?php

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;
use Doctrine\ORM\EntityManager;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Serializer\SerializerInterface;

use App\Entity\Board;

use App\Form\Type\InsertTileType;
use App\Form\Type\JoinBoardType;
use App\Form\Type\MovePlayerType;
use App\Form\Type\NewBoardType;
use App\Form\Type\RotateRemainingType;
use App\Service\Direction;
use App\Service\Rotation;
use App\Service\DomainServiceInterface;

class BoardController extends BoardBaseController
{
    const TREASURE_EMOJIS = [
        '' => ' ',
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
        protected DomainServiceInterface $domainService,
        protected ManagerRegistry $doctrine,
        protected HubInterface $hub,
        protected SerializerInterface $serializer,
    ) {
        parent::__construct($domainService, $doctrine->getManager(), $hub, $serializer);
    }

    #[Route('/board/new', name: 'board_new', methods: 'POST')]
    public function getNew(Request $request, ManagerRegistry $doctrine)
    {
        $user = $this->getCurrentUser($request);
        if ($user == NULL) {
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

        $playerCount = intval($form->getData()['player_count']);
        $board = $this->newBoard($user, $playerCount);

        return $this->redirectToRoute('board_view', [
            'id' => $board->getId(),
        ]);
    }

    #[Route('/board/{id}/view', name: 'board_view', methods: 'GET')]
    public function getView(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        $user = $this->getCurrentUser($request);
        if ($board->getRemainingSeats() > 0) {
            $form = $this->createForm(JoinBoardType::class, null, [
                'action' => $this->generateUrl('board_join', ['id' => $board->getId()]),
            ]);

            return $this->render('board/lobby.html.twig', [
                'board' => $board,
                'form' => $form,
                'canJoin' => $this->canUserJoin($user, $board),
            ]);
        }

        return $this->render('board/view.html.twig', [
            'board' => $this->createBoardViewModel($user, $board),
            'emojis' => static::TREASURE_EMOJIS,
            'colors' => static::PLAYER_COLORS,
        ]);
    }

    #[Route('/board/{id}/join', name: 'board_join', methods: 'POST')]
    public function postJoin(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        $this->joinBoard($this->getCurrentUser($request), $board);
        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }

    private function postRotateRemaining(Request $request, Board $board, Rotation $rotation): Response
    {
        $user = $this->getCurrentUser($request);
        if (!$this->canUserPlay($user, $board)) {
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $form = $this->createForm(RotateRemainingType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $this->rotateRemaining($board, $rotation);
        }

        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }

    #[Route('/board/{id}/rotate-remaining-clockwise', name: 'board_rotate_remaining_clockwise', methods: 'POST')]
    public function postRotateRemainingClockwise(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->postRotateRemaining($request, $board, Rotation::CLOCKWISE);
    }

    #[Route('/board/{id}/rotate-remaining-anticlockwise', name: 'board_rotate_remaining_anticlockwise', methods: 'POST')]
    public function postRotateRemainingAnticlockwise(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->postRotateRemaining($request, $board, Rotation::ANTICLOCKWISE);
    }

    private function postInsertTile(Request $request, Board $board, Direction $direction): Response
    {
        $user = $this->getCurrentUser($request);
        if (!$this->canUserPlay($user, $board)) {
            $this->addFlash('errors', 'You cannot perform this action.');
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $form = $this->createForm(InsertTileType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $insertTile = $form->getData();

            $this->insertTile($board, $direction, intval($insertTile['index']));
        }

        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }

    #[Route('/board/{id}/insert-tile-top', name: 'board_insert_tile_top', methods: 'POST')]
    public function postInsertTileTop(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->postInsertTile($request, $board, Direction::TOP);
    }

    #[Route('/board/{id}/insert-tile-right', name: 'board_insert_tile_right', methods: 'POST')]
    public function postInsertTileRight(Request $request, ManagerRegistry $doctrine, Board $board): Response
    {
        return $this->postInsertTile($request, $board, Direction::RIGHT);
    }

    #[Route('/board/{id}/insert-tile-bottom', name: 'board_insert_tile_bottom', methods: 'POST')]
    public function postInsertTileBottom(Request $request, Board $board): Response
    {
        return $this->postInsertTile($request, $board, Direction::BOTTOM);
    }

    #[Route('/board/{id}/insert-tile-left', name: 'board_insert_tile_left', methods: 'POST')]
    public function postInsertTileLeft(Request $request, Board $board): Response
    {
        return $this->postInsertTile($request, $board, Direction::LEFT);
    }

    #[Route('/board/{id}/move-player', name: 'board_move_player', methods: 'POST')]
    public function postMovePlayer(Request $request, Board $board): Response
    {
        $user = $this->getCurrentUser($request);
        if (!$this->canUserPlay($user, $board)) {
            $this->addFlash('errors', 'You cannot perform this action.');
            return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
        }

        $form = $this->createForm(MovePlayerType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $movePlayer = $form->getData();
            $this->movePlayer(
                $board,
                intval($movePlayer['line']),
                intval($movePlayer['row']),
            );
        }
        return $this->redirectToRoute('board_view', ['id' => $board->getId()]);
    }
}
