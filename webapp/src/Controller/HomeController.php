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

    #[Route('/', name: 'home')]
    public function index(Request $request): Response
    {
        return $this->render('home/index.html.twig');
    }
}