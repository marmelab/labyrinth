<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Service\DomainServiceInterface;

class HomeController extends AbstractController
{

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
        ]);
    }
}