<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Form\Type\NewBoardType;
use App\Form\Type\SignInType;

class HomeController extends AbstractController
{

    #[Route('/', name: 'home')]
    public function index(Request $request): Response
    {
        $signInForm = $this->createForm(SignInType::class, null, [
            'action' => $this->generateUrl('player_sign_in'),
        ]);

        $newBoardForm = $this->createForm(NewBoardType::class, null, [
            'action' => $this->generateUrl('board_new'),
        ]);

        return $this->render('home/index.html.twig', [
            'signInForm' => $signInForm,
            'newBoardForm' => $newBoardForm,
        ]);
    }
}
