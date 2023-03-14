<?php

namespace App\Controller;

use Doctrine\Persistence\ManagerRegistry;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

use App\Entity\Player;

use App\Form\Type\NewBoardType;
use App\Form\Type\SignInType;

class HomeController extends AbstractController
{

    #[Route('/', name: 'home')]
    public function index()
    {
        /** @var User $user */
        $user = $this->getUser();
        $boards = [];
        if ($user) {
            $boards = $user->getBoards();
        }

        $newBoardForm = $this->createForm(NewBoardType::class, null, [
            'action' => $this->generateUrl('board_new'),
        ]);

        return $this->render('home/index.html.twig', [
            'newBoardForm' => $newBoardForm,
            'boards' => $boards,
        ]);
    }
}
