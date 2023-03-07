<?php

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\Player;
use App\Form\Type\SignInType;

class PlayerController extends AuthBaseController
{
    const SESSION_PLAYER_KEY = 'player';

    public function __construct(
        protected ManagerRegistry $doctrine,
    ) {
        parent::__construct($doctrine->getManager());
    }

    #[Route('/player/sign-in', name: 'player_sign_in', methods: 'POST')]
    public function postSignIn(Request $request, ManagerRegistry $doctrine): Response
    {
        $form = $this->createForm(SignInType::class, new Player());
        $form->handleRequest($request);

        if (!$form->isSubmitted() || !$form->isValid()) {
            return $this->render('player/sign_in.html.twig', [
                'form' => $form,
            ]);
        }

        $this->signInUser($request, $form->getData()->getName());
        return $this->redirectToRoute('home');
    }
    #[Route('/player/sign-out', name: 'player_sign_out', methods: 'POST')]
    public function postSignOut(Request $request): Response
    {
        $this->signOutUser($request);
        return $this->redirectToRoute('home');
    }
}
