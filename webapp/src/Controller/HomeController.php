<?php

namespace App\Controller;

use App\Form\Type\InsertTileType;
use App\Service\Rotation;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Form\Type\RotateRemainingType;
use App\Service\DomainServiceInterface;

class HomeController extends AbstractController
{
    const SESSION_BOARD_KEY = BoardController::SESSION_BOARD_KEY;

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
    )
    {

    }

    #[Route('/', name: 'home')]
    public function index(Request $request): Response
    {
        $session = $request->getSession();

        if (!$session->has(self::SESSION_BOARD_KEY)) {
            $newBoard = $this->domainService->newBoard();
            $session->set(self::SESSION_BOARD_KEY, $newBoard);
        }

        $rotateRemainingClockwise = $this->createForm(RotateRemainingType::class);
        $rotateRemainingAnticlockwise = $this->createForm(RotateRemainingType::class);

        $board = $session->get(self::SESSION_BOARD_KEY);
        return $this->render('home/index.html.twig', [
            'board' => $board,
            'emojis' => self::TREASURE_EMOJIS,
            'rotationForms' => [
                'clockwise' => $rotateRemainingClockwise->createView(),
                'anticlockwise' => $rotateRemainingAnticlockwise->createView(),
            ],
        ]);
    }
}