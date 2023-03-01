<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Form\Type\RotateRemainingType;
use App\Service\Rotation;
use App\Service\DomainServiceInterface;

class BoardController extends AbstractController
{
    const SESSION_BOARD_KEY = 'board';

    public function __construct(
        private DomainServiceInterface $domainService
    )
    {

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
}