<?php

namespace App\Controller;

use Doctrine\Persistence\ObjectManager;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;

use App\Entity\Player;

const SESSION_PLAYER_KEY = 'player';

abstract class AuthBaseController extends AbstractController
{
    const SESSION_PLAYER_KEY = 'player';

    protected function __construct(
        protected ObjectManager $entityManager,
    ) {
    }

    protected function signInUser(Request $request, string $name)
    {
        $playerRepository = $this->entityManager->getRepository(Player::class);

        $player = $playerRepository->findOneByName($name);
        if ($player == NULL) {
            $player = new Player();
            $player->setName($name);

            $this->entityManager->persist($player);
            $this->entityManager->flush();
        }

        $request->getSession()->set(static::SESSION_PLAYER_KEY, $player);
    }

    protected function signOutUser(Request $request)
    {
        $request->getSession()->remove(static::SESSION_PLAYER_KEY);
    }
}
