<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\Routing\Annotation\Route;

use App\Form\Type\NewBoardType;
use App\Form\Type\SignUpType;

class HomeController extends AbstractController
{

    #[Route('/', name: 'home')]
    public function index()
    {
        /** @var User $user */
        $user = $this->getUser();
        $boards = [];
        if ($user) {
            $boards = array_map(function ($game) {
                return $game->getBoard();
            }, $user->getGames()->toArray());
        }

        $newBoardForm = $this->createForm(NewBoardType::class, null, [
            'action' => $this->generateUrl('board_new'),
        ]);

        $signUpForm = $this->createForm(SignUpType::class, null, [
            'action' => $this->generateUrl('auth_sign_up_post'),
        ]);

        return $this->render('home/index.html.twig', [
            'newBoardForm' => $newBoardForm,
            'signUpForm' => $signUpForm,
            'boards' => $boards,
        ]);
    }
}
