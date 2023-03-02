<?php

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\Player;
use App\Form\Type\SignInType;

class PlayerController extends AbstractController
{

    #[Route('/player/sign-in', name: 'player_sign_in', methods: 'POST')]
    public function postSignIn(Request $request, ManagerRegistry $doctrine): Response
    {
        $form = $this->createForm(SignInType::class, new Player());
        $form->handleRequest($request);

        if ($form->isSubmitted() && !$form->isValid()) {
            return $this->render('player/sign_in.html.twig', [
                'form' => $form,
            ]);
        }

        $name = $form->getData()->getName();

        $entityManager = $doctrine->getManager();
        $playerRepository = $entityManager->getRepository(Player::class);

        $player = $playerRepository->findOneByName($name);
        if ($player == NULL) {
            $player = new Player();
            $player->setName($name);

            $entityManager->persist($player);
            $entityManager->flush();
        }

        $request->getSession()->set('player', $player);
        return $this->redirectToRoute('home');
    }
    #[Route('/player/sign-out', name: 'player_sign_out', methods: 'POST')]
    public function postSignOut(Request $request): Response
    {
        $request->getSession()->remove('player');
        return $this->redirectToRoute('home');
    }
}
