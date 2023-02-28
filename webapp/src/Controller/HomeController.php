<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Service\DomainServiceInterface;

class HomeController extends AbstractController
{
    const TREASURE_EMOJIS = [
        '·' => '',
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
    public function index(): Response
    {
        $board = $this->domainService->newBoard();

        return $this->render('home/index.html.twig', [
            'board' => $board,
            'emojis' => self::TREASURE_EMOJIS,
        ]);
    }
}