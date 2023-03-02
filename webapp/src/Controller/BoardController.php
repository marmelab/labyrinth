<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Form\Type\InsertTileType;
use App\Form\Type\RotateRemainingType;
use App\Service\Direction;
use App\Service\Rotation;
use App\Service\DomainServiceInterface;

class BoardController extends AbstractController
{
    const SESSION_BOARD_KEY = 'board';

    public function __construct(
        private DomainServiceInterface $domainService
    ) {
    }

    private function rotateRemaining(Request $request, Rotation $rotation): Response
    {
        $session = $request->getSession();

        $form = $this->createForm(RotateRemainingType::class);
        $form->handleRequest($request);

        if ($session->has(self::SESSION_BOARD_KEY) && $form->isSubmitted() && $form->isValid()) {
            $savedBoard = $session->get(self::SESSION_BOARD_KEY);
            $updatedBoard = $this->domainService->rotateRemainingTile($savedBoard, $rotation);
            $session->set(self::SESSION_BOARD_KEY, $updatedBoard);
        }
        return $this->redirectToRoute('home');
    }

    #[Route('/board/rotate-remaining-clockwise', name: 'board_rotate_remaining_clockwise', methods: 'POST')]
    public function postRotateRemainingClockwise(Request $request): Response
    {
        return $this->rotateRemaining($request, Rotation::CLOCKWISE);
    }

    #[Route('/board/rotate-remaining-anticlockwise', name: 'board_rotate_remaining_anticlockwise', methods: 'POST')]
    public function postRotateRemainingAnticlockwise(Request $request): Response
    {
        return $this->rotateRemaining($request, Rotation::ANTICLOCKWISE);
    }

    private function insertTile(Request $request, Direction $direction): Response
    {
        $session = $request->getSession();

        $form = $this->createForm(InsertTileType::class);
        $form->handleRequest($request);

        if ($session->has(self::SESSION_BOARD_KEY) && $form->isSubmitted() && $form->isValid()) {
            $insertTile = $form->getData();

            $savedBoard = $session->get(self::SESSION_BOARD_KEY);
            $updatedBoard = $this->domainService->insertTile($savedBoard, $direction, intval($insertTile['index']));
            $session->set(self::SESSION_BOARD_KEY, $updatedBoard);
        }
        return $this->redirectToRoute('home');
    }

    #[Route('/board/insert-tile-top', name: 'board_insert_tile_top', methods: 'POST')]
    public function postInsertTileTop(Request $request): Response
    {
        return $this->insertTile($request, Direction::TOP);
    }

    #[Route('/board/insert-tile-right', name: 'board_insert_tile_right', methods: 'POST')]
    public function postInsertTileRight(Request $request): Response
    {
        return $this->insertTile($request, Direction::RIGHT);
    }

    #[Route('/board/insert-tile-bottom', name: 'board_insert_tile_bottom', methods: 'POST')]
    public function postInsertTileBottom(Request $request): Response
    {
        return $this->insertTile($request, Direction::BOTTOM);
    }

    #[Route('/board/insert-tile-left', name: 'board_insert_tile_left', methods: 'POST')]
    public function postInsertTileLeft(Request $request): Response
    {
        return $this->insertTile($request, Direction::LEFT);
    }
}
