<?php

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\Board;
use App\Form\Type\InsertTileType;
use App\Form\Type\RotateRemainingType;
use App\Service\Direction;
use App\Service\Rotation;
use App\Service\DomainServiceInterface;

class BoardController extends AbstractController
{
    const SESSION_BOARD_KEY = 'board';

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

    public function __construct(
        private DomainServiceInterface $domainService
    ) {
    }

    private function getBoardState(Board $board): array
    {
        return json_decode($board->getState(), true);
    }

    #[Route('/board/new', name: 'board_new')]
    public function getNew(ManagerRegistry $doctrine)
    {
        $entityManager = $doctrine->getManager();

        $boardState = $this->domainService->newBoard();

        $board = new Board();
        $board->setState(json_encode($boardState));

        $entityManager->persist($board);
        $entityManager->flush();

        return $this->redirectToRoute('board_view', [
            'id' => $board->getId(),
        ]);
    }

    #[Route('/board/{id}/view', name: 'board_view')]
    public function getView(Request $request, Board $board): Response
    {
        return $this->render('board/view.html.twig', [
            'board' => $board,
            'boardState' => $this->getBoardState($board),
            'emojis' => self::TREASURE_EMOJIS,
        ]);
    }

    private function rotateRemaining(Request $request, ManagerRegistry $doctrine, Board $board, Rotation $rotation): Response
    {
        $form = $this->createForm(RotateRemainingType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $updatedBoard = $this->domainService->rotateRemainingTile($this->getBoardState($board), $rotation);
            $board->setState(json_encode($updatedBoard));

            $doctrine->getManager()->flush();
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
        $form = $this->createForm(InsertTileType::class);
        $form->handleRequest($request);

        if ($form->isSubmitted() && $form->isValid()) {
            $insertTile = $form->getData();

            $updatedBoard = $this->domainService->insertTile(
                $this->getBoardState($board),
                $direction,
                intval($insertTile['index'])
            );

            $board->setState(json_encode($updatedBoard));

            $doctrine->getManager()->flush();
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
}
