<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Service\DomainServiceInterface;

class HomeController extends AbstractController
{
    const SESSION_BOARD_KEY = 'board';

    const TREASURE_EMOJIS = [
        '·' => ' ',
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
            $session->set(self::SESSION_BOARD_KEY, json_encode($newBoard));
        }

        $board = json_decode($session->get(self::SESSION_BOARD_KEY), true);
        return $this->render('home/index.html.twig', [
            'board' => $board,
            'emojis' => self::TREASURE_EMOJIS,
        ]);
    }
}